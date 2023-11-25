package microservices

import (
	"fmt"
	"log"
	"net"
	"server/internal/config"
	"server/internal/services/microservice"

	"github.com/asaskevich/govalidator"
	"google.golang.org/grpc"
)

const configPath string = "config/config.yml"
const envPath string = "config/.env"

func main() {
	config, err := config.LoadConfig(envPath, configPath)
	govalidator.SetFieldsRequiredByDefault(true)
	if err != nil {
		log.Fatal(err.Error())
	}
	logger := config.Logging.Logger
	logger.Info("Config loaded")

	lstn, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Server.MicroservicePort))
	if err != nil {
		logger.Fatal("Can't listen to port, " + err.Error())
	}

	server := grpc.NewServer()

	// microservice.RegisterCSATSAnswerServiceServer(server, NewCSATAnswerService())
	microservice.RegisterCSATSAnswerServiceServer(server)

	server.Serve(lstn)
}
