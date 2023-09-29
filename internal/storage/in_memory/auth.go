package in_memory

import (
	"server/internal/pkg/datatypes"
	"sync"
)

type LocalAuthStorage struct {
	authData map[string]datatypes.LoginInfo
	mu       *sync.Mutex
}

func NewLocalTestStorage() *LocalAuthStorage {
	return &LocalAuthStorage{
		authData: map[string]datatypes.LoginInfo{
			"test@email.com": {
				Email:        "test@email.com",
				PasswordHash: "$2a$08$YkQXrizJ.TDF.dYo58hNFuHwATMIdZHbWwgfI.vuSQEEurB6zpgvy",
			},
			"email@example.com": {
				Email:        "email@example.com",
				PasswordHash: "$2a$08$5vGskE/R50Ju92.4AbbZyeQiBT26Hiiq.4RqoRf5yGOrExfKDCW52",
			},
		},
		mu: &sync.Mutex{},
	}
}

func (a *LocalAuthStorage) VerifyLogin() error {
	return nil
}
