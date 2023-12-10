package microservice

/*
func TestNewWorkspaceService(t *testing.T) {
	type args struct {
		storage storage.IWorkspaceStorage
	}
	tests := []struct {
		name string
		args args
		want *WorkspaceService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWorkspaceService(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWorkspaceService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkspaceService_Create(t *testing.T) {
	type fields struct {
		storage storage.IWorkspaceStorage
	}
	type args struct {
		ctx  context.Context
		info dto.NewWorkspaceInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Workspace
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := WorkspaceService{
				storage: tt.fields.storage,
			}
			got, err := ws.Create(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkspaceService_Delete(t *testing.T) {
	type fields struct {
		storage storage.IWorkspaceStorage
	}
	type args struct {
		ctx context.Context
		id  dto.WorkspaceID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := WorkspaceService{
				storage: tt.fields.storage,
			}
			if err := ws.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWorkspaceService_GetUserWorkspaces(t *testing.T) {
	type fields struct {
		storage storage.IWorkspaceStorage
	}
	type args struct {
		ctx    context.Context
		userID dto.UserID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.AllWorkspaces
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := WorkspaceService{
				storage: tt.fields.storage,
			}
			got, err := ws.GetUserWorkspaces(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserWorkspaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserWorkspaces() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkspaceService_GetWorkspace(t *testing.T) {
	type fields struct {
		storage storage.IWorkspaceStorage
	}
	type args struct {
		ctx context.Context
		id  dto.WorkspaceID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Workspace
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := WorkspaceService{
				storage: tt.fields.storage,
			}
			got, err := ws.GetWorkspace(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWorkspace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWorkspace() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkspaceService_UpdateData(t *testing.T) {
	type fields struct {
		storage storage.IWorkspaceStorage
	}
	type args struct {
		ctx  context.Context
		info dto.UpdatedWorkspaceInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := WorkspaceService{
				storage: tt.fields.storage,
			}
			if err := ws.UpdateData(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdateData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWorkspaceService_UpdateUsers(t *testing.T) {
	type fields struct {
		storage storage.IWorkspaceStorage
	}
	type args struct {
		ctx  context.Context
		info dto.ChangeWorkspaceGuestsInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := WorkspaceService{
				storage: tt.fields.storage,
			}
			if err := ws.UpdateUsers(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
*/
