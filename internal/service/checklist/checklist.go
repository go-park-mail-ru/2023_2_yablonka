package checklist

import (
	"server/internal/storage"

	micro "server/internal/service/checklist/microservice"

	"google.golang.org/grpc"
)

// TODO: Checklist microservice
func NewMicroChecklistService(checklistStorage storage.IChecklistStorage, connection *grpc.ClientConn) *micro.ChecklistService {
	return micro.NewChecklistService(checklistStorage)
}
