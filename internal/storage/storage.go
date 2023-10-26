package storage

type Storage struct {
	auth  IAuthStorage
	user  IUserStorage
	board IBoardStorage
}
