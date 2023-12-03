package checklist

import (
	"server/internal/storage"

	embedded "server/internal/service/checklist/embedded"
	micro "server/internal/service/checklist/microservice"

	"google.golang.org/grpc"
)

func NewEmbeddedChecklistService(checklistStorage storage.IChecklistStorage) *embedded.ChecklistService {
	return embedded.NewChecklistService(checklistStorage)
}

// TODO: Checklist microservice
func NewMicroChecklistService(checklistStorage storage.IChecklistStorage, connection *grpc.ClientConn) *micro.ChecklistService {
	return micro.NewChecklistService(checklistStorage)
}
