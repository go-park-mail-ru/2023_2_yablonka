package checklist_item

import (
	"server/internal/storage"

	embedded "server/internal/service/checklist_item/embedded"
	micro "server/internal/service/checklist_item/microservice"

	"google.golang.org/grpc"
)

func NewEmbeddedChecklistItemService(checklistItemStorage storage.IChecklistItemStorage) *embedded.ChecklistItemService {
	return embedded.NewChecklistItemService(checklistItemStorage)
}

// TODO: Checklist item microservice
func NewMicroChecklistItemService(checklistItemStorage storage.IChecklistItemStorage, connection *grpc.ClientConn) *micro.ChecklistItemService {
	return micro.NewChecklistItemService(checklistItemStorage)
}
