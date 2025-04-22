package handler

import (
	"encoding/json"
	"geoservise-jwt/internal/model"
	"geoservise-jwt/internal/usecase/responder"
	"geoservise-jwt/internal/usecase/service"
	"net/http"
)

type UserHandler struct {
	service   service.UserService
	responder responder.Responder
}

func NewUserHandler(s service.UserService, r responder.Responder) *UserHandler {
	return &UserHandler{
		service:   s,
		responder: r,
	}
}

// Create godoc
// @Summary Create user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.CreateUserRequest true "User name and password"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal error"
// @Router /register [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.responder.Error(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.Create(r.Context(), model.User{Username: req.Username, Password: req.Password})
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err)
		return
	}
	h.responder.JSON(w, http.StatusCreated, map[string]interface{}{"id": id, "message": "user created"})
}
