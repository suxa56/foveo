package handler

import (
	"encoding/json"
	"fmt"
	"foveo/internal/domain/dto"
	"foveo/internal/service"
	"net/http"
)

type UserHandler struct {
	userService service.IUserService
}

func InitUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	errors := req.IsValid()
	if len(errors) > 0 {
		for _, err := range errors {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// Передаем в сервис
	err = h.userService.Register(r.Context(), req)
	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User registered successfully")
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// получение по id
	w.Write([]byte("get user by id"))
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	// обновление
	w.Write([]byte("update user"))
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// удаление
	w.Write([]byte("delete user"))
}
