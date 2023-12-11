package workspace

import (
	"server/internal/storage"

	micro "server/internal/service/workspace/microservice"

	"google.golang.org/grpc"
)

// TODO: User microservice
func NewMicroWorkspaceService(workspaceStorage storage.IWorkspaceStorage, connection *grpc.ClientConn) *micro.WorkspaceService {
	return micro.NewWorkspaceService(workspaceStorage)
}
