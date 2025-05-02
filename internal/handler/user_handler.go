package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"foveo/internal/domain/dto"
	"foveo/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
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

	errs := req.IsValid()
	if len(errs) > 0 {
		for _, err := range errs {
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

func (h *UserHandler) GetList(w http.ResponseWriter, r *http.Request) {
	list, err := h.userService.GetList(r.Context())
	if err != nil {
		http.Error(w, "failed to get users", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	// обновление
	w.Write([]byte("update user"))
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// удаление
	w.Write([]byte("delete user"))
}
