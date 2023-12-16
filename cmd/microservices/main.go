package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"server/internal/config"
	logging "server/internal/logging"
	"server/internal/storage"
	"server/internal/storage/postgresql"
	auth "server/microservices/auth"
	csat "server/microservices/csat"
	csrf "server/microservices/csrf"
	user "server/microservices/user"

	"github.com/asaskevich/govalidator"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const configPath string = "config/config.yml"
const envPath string = "config/.env"

func main() {
	config, err := config.LoadConfig(envPath, configPath)
	govalidator.SetFieldsRequiredByDefault(true)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Config loaded")

	logger, err := logging.NewLogrusLogger(config.Logging)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Logger configured")

	dbConnection, err := postgresql.GetDBConnection(*config.Database)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer dbConnection.Close()
	logger.Info("Database connection established")

	storages := storage.NewPostgresStorages(dbConnection)
	logger.Info("Storages configured")

	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	reg := prometheus.NewRegistry()
	reg.MustRegister(srvMetrics)
	exemplarFromContext := func(ctx context.Context) prometheus.Labels {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return prometheus.Labels{"traceID": span.TraceID().String()}
		}
		return nil
	}

	panicsTotal := promauto.With(reg).NewCounter(prometheus.CounterOpts{
		Name: "grpc_req_panics_recovered_total",
		Help: "Total number of gRPC requests recovered from internal panic.",
	})
	grpcPanicRecoveryHandler := func(p any) (err error) {
		panicsTotal.Inc()
		logger.Error("*************** PANIC ***************")
		logger.Error(fmt.Sprintf("Recovered from panic %v", err))
		logger.Error("*************** CONTINUING ***************")
		return status.Errorf(codes.Internal, "%s", p)
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			// Order matters e.g. tracing interceptor have to create span first for the later exemplars to work.
			// otelgrpc.UnaryServerInterceptor(),
			srvMetrics.UnaryServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			// logging.UnaryServerInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
			// selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(authFn), selector.MatchFunc(allButHealthZ)),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
		grpc.ChainStreamInterceptor(
			// otelgrpc.StreamServerInterceptor(),
			srvMetrics.StreamServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			// logging.StreamServerInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
			// selector.StreamServerInterceptor(auth.StreamServerInterceptor(authFn), selector.MatchFunc(allButHealthZ)),
			recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
	)

	lstn, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Server.MicroservicePort))
	if err != nil {
		logger.Fatal("Can't listen to port, " + err.Error())
	}

	logger.Info("Registering CSAT services")
	csat.RegisterServices(storages, server, &logger)
	logger.Info("Registering auth services")
	auth.RegisterServices(config, storages, server, &logger)
	logger.Info("Registering CSRF services")
	csrf.RegisterServices(config, storages, server, &logger)
	logger.Info("Registering user services")
	user.RegisterServices(storages, server, &logger)

	go func() {
		http.Handle("/metrics", promhttp.HandlerFor(
			reg,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			},
		))
		http.ListenAndServe(":8013", nil)
	}()

	err = server.Serve(lstn)
	if err != nil {
		logger.Fatal("Server failure " + err.Error())
	}
}
