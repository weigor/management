package server

import "management/pkg/server/register"

type Handler struct {
	*register.Service
}

func NewHandler() *Handler {
	return &Handler{Service: register.NewService()}
}
