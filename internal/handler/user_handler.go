package handler

import (
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
	// парсим запрос, валидируем
	// вызываем h.userService.CreateUser(...)
	// возвращаем ответ
	w.Write([]byte("create user"))
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
