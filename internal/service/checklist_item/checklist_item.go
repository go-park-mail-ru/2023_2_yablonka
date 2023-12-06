package checklist_item

import (
	"server/internal/storage"

	micro "server/internal/service/checklist_item/microservice"

	"google.golang.org/grpc"
)

// TODO: Checklist item microservice
func NewMicroChecklistItemService(checklistItemStorage storage.IChecklistItemStorage, connection *grpc.ClientConn) *micro.ChecklistItemService {
	return micro.NewChecklistItemService(checklistItemStorage)
}
