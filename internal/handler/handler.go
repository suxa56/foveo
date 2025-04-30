package handler

type Handler struct {
	User *UserHandler
}

func NewHandler(user *UserHandler) *Handler {
	return &Handler{
		User: user,
	}
}
