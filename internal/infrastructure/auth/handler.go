package auth

import (
	"encoding/json"
	"errors"
	"geoservise-jwt/internal/usecase/responder"
	"geoservise-jwt/internal/usecase/service"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type AuthHandler struct {
	TokenAuth   *jwtauth.JWTAuth
	UserService service.UserService
	Responder   responder.Responder
}

func NewAuthHandler(tokenAuth *jwtauth.JWTAuth, userService service.UserService, r responder.Responder) *AuthHandler {
	return &AuthHandler{
		TokenAuth:   tokenAuth,
		UserService: userService,
		Responder:   r,
	}
}

// LoginHandler godoc
// @Summary Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body Credentials true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "bad request"
// @Router /login [post]
func (a *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req Credentials
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.Responder.Error(w, http.StatusBadRequest, errors.New("неправильный формат запроса"))
		return
	}

	if req.Username == "" || req.Password == "" {
		a.Responder.Error(w, http.StatusBadRequest, errors.New("не хватает username или password"))
		return
	}

	user, err := a.UserService.GetByUsername(r.Context(), req.Username)
	if err != nil {
		a.Responder.Error(w, http.StatusUnauthorized, errors.New("пользователь не найден"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		a.Responder.Error(w, http.StatusUnauthorized, errors.New("неверный пароль"))
		return
	}

	if a.TokenAuth == nil {
		a.Responder.Error(w, http.StatusInternalServerError, errors.New("TokenAuth не инициализирован"))
		return
	}

	// Генерация токена
	_, token, _ := a.TokenAuth.Encode(map[string]interface{}{"sub": req.Username})

	// Установка токена в куку
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	a.Responder.JSON(w, http.StatusOK, map[string]string{"message": "logged in"})
}
