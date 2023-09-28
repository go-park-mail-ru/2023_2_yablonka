package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"

	jwt "github.com/dgrijalva/jwt-go"

	datatypes "server/internal/pkg/datatypes"
)

type Api struct {
	users      map[string]*datatypes.User
	sessions   map[string]uint64
	boards     map[uint64]*datatypes.Board
	jwt_secret []byte
	mu         *sync.Mutex
}

func TestApi() *Api {
	jwt_secret := []byte(os.Getenv("JWT_SECRET"))

	return &Api{
		users: map[string]*datatypes.User{
			"test@email.com": {
				ID:           1,
				Email:        "test@email.com",
				PasswordHash: "123456",
				FullName:     "Никита Архаров",
			},
			"email@example.com": {
				ID:           2,
				Email:        "email@example.com",
				PasswordHash: "135790",
				FullName:     "Даниил Капитанов",
			},
		},
		sessions: make(map[string]uint64, 10),
		boards: map[uint64]*datatypes.Board{
			1: {
				ID:        1,
				BoardName: "Тест доска 1",
				OwnerID:   1,
			},
			2: {
				ID:        2,
				BoardName: "Тест доска 2",
				OwnerID:   1,
			},
			3: {
				ID:        3,
				BoardName: "Бэкэнд Tabula",
				OwnerID:   2,
			},
		},
		jwt_secret: jwt_secret,
		mu:         &sync.Mutex{},
	}
}

func (api *Api) GenerateJWT(user *datatypes.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"ID":    user.ID,
	})

	str, err := token.SignedString(api.jwt_secret)
	if err != nil {
		return "", err
	}
	return str, nil
}

func (api *Api) GetSessions() map[string]uint64 {
	return api.sessions
}

func (api *Api) GetUsers() map[string]*datatypes.User {
	return api.users
}

func (api *Api) GetUserByEmail(email string) (*datatypes.User, bool) {
	user, ok := api.users[email]
	return user, ok
}

func (api *Api) GetHighestID() uint64 {
	if len(api.users) == 0 {
		return 0
	}
	var highest uint64 = 0
	for _, user := range api.users {
		if user.ID > highest {
			highest = user.ID
		}
	}
	return highest
}

func (api *Api) HandleSignupUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var signup datatypes.SignupInfo
	err := json.NewDecoder(r.Body).Decode(&signup)
	if err != nil {
		http.Error(w, `Signup error`, http.StatusBadRequest)
		return
	}

	api.mu.Lock()
	_, ok := api.users[signup.Email]
	api.mu.Unlock()
	if ok {
		http.Error(w, `User already exists`, http.StatusConflict)
		return
	}

	api.mu.Lock()
	var newID uint64 = api.getHighestID() + 1
	api.users[signup.Email] = &datatypes.User{
		ID:           newID,
		Email:        signup.Email,
		PasswordHash: signup.PasswordHash,
		FullName:     "",
	}
	api.mu.Unlock()

	w.Write([]byte(`{"body": {}}`))
}

func (api *Api) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var login datatypes.LoginInfo
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, `Login error`, http.StatusBadRequest)
		return
	}

	api.mu.Lock()
	user, ok := api.users[login.Email]
	api.mu.Unlock()
	if !ok {
		http.Error(w, `User not found`, http.StatusNotFound)
		return
	}

	if user.PasswordHash != login.PasswordHash {
		http.Error(w, `Wrong password!`, http.StatusUnauthorized)
		return
	}

	token, err := api.GenerateJWT(user)
	if err != nil {
		http.Error(w, `Session error`, http.StatusInternalServerError)
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}

	http.SetCookie(w, cookie)

	api.mu.Lock()
	api.sessions[token] = user.ID
	api.mu.Unlock()

	w.Write([]byte(`{"body": {}}`))
}

func (api *Api) HandleGetUserBoards(w http.ResponseWriter, r *http.Request) {
}

func (api *Api) HandleVerifyAuth(w http.ResponseWriter, r *http.Request) {
	// encoder := json.NewEncoder(w)
	// api.Mu.Lock()
	// err := encoder.Encode(api.Boards)
	// api.Mu.Unlock()
	// if err != nil {
	// 	log.Printf("error while marshalling JSON: %s", err)
	// 	w.Write([]byte("{}"))
	// 	return
	// }
}
