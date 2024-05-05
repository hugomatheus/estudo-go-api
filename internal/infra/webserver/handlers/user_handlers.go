package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/hugomatheus/go-api/internal/dto"
	"github.com/hugomatheus/go-api/internal/entity"
	"github.com/hugomatheus/go-api/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

// GetJWT godoc
// @Summary Get a user JWT
// @Description Get a user JWT
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.GetJwtInput true "user credentials"
// @Success 200 {object} dto.GetJwtOutput
// @Failure 404 {object} dto.ErrorOutput
// @Failure 500 {object} dto.ErrorOutput
// @Router /users/login [post]
func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExperesIn := r.Context().Value("jwtExperesIn").(int)

	var user dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := dto.ErrorOutput{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, tokenString, err := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExperesIn)).Unix(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	accessToken := dto.GetJwtOutput{AccessToken: tokenString}
	w.Header().Set("Contenxt-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary Create user
// @Description Create user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserInput true "user request"
// @Success 201
// @Failure 500 {object} dto.ErrorOutput
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.ErrorOutput{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.ErrorOutput{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
