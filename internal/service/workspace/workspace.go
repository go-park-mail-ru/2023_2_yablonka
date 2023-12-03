package workspace

import (
	"server/internal/storage"

	embedded "server/internal/service/workspace/embedded"
	micro "server/internal/service/workspace/microservice"

	"google.golang.org/grpc"
)

func NewEmbeddedWorkspaceService(workspaceStorage storage.IWorkspaceStorage) *embedded.WorkspaceService {
	return embedded.NewWorkspaceService(workspaceStorage)
}

// TODO: User microservice
func NewMicroWorkspaceService(workspaceStorage storage.IWorkspaceStorage, connection *grpc.ClientConn) *micro.WorkspaceService {
	return micro.NewWorkspaceService(workspaceStorage)
}
