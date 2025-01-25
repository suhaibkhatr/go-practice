package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/suhaibkhatr/go-practice/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &modeUserStore{}
	handler := NewHandler(userStore)
	t.Run("shoud fail the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUser{
			FirstName: "Suhaib",
			LastName:  "Khater",
			Email:     "suhaib@gmail.com",
			Password:  "123456",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type modeUserStore struct {
}

func (mode *modeUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}
func (mode *modeUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (mode *modeUserStore) CreateUser(types.User) error {
	return nil
}
