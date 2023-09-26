package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	datatypes "server/internal/pkg"
)

type Handlers struct {
	Users  []datatypes.User
	Boards []datatypes.Board
	Mu     *sync.Mutex
}

func (h *Handlers) HandleSignupUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	newUserInput := new(datatypes.SignupInfo)
	err := decoder.Decode(newUserInput)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.Write([]byte("{}"))
		return
	}

	fmt.Println(newUserInput)
	h.Mu.Lock()

	var id uint64 = 0
	// Сначала проверить почту на существование
	if len(h.Users) > 0 {
		id = h.Users[len(h.Users)-1].ID + 1
	}

	h.Users = append(h.Users, datatypes.User{
		ID:           id,
		Email:        newUserInput.Email,
		PasswordHash: newUserInput.PasswordHash,
		FullName:     "",
	})
	h.Mu.Unlock()
}

func (h *Handlers) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	newUserInput := new(datatypes.SignupInfo)
	err := decoder.Decode(newUserInput)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.Write([]byte("{}"))
		return
	}

	fmt.Println(newUserInput)
	h.Mu.Lock()
	// do stuff
	h.Mu.Unlock()
}

func (h *Handlers) HandleGetBoards(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	h.Mu.Lock()
	err := encoder.Encode(h.Boards)
	h.Mu.Unlock()
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.Write([]byte("{}"))
		return
	}
}

func (h *Handlers) HandleVerifyAuth(w http.ResponseWriter, r *http.Request) {
	// encoder := json.NewEncoder(w)
	// h.Mu.Lock()
	// err := encoder.Encode(h.Boards)
	// h.Mu.Unlock()
	// if err != nil {
	// 	log.Printf("error while marshalling JSON: %s", err)
	// 	w.Write([]byte("{}"))
	// 	return
	// }
}
