package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Subham-Das-98/go-rest-api/internal/dto"
	"github.com/Subham-Das-98/go-rest-api/internal/service"
	"github.com/Subham-Das-98/go-rest-api/internal/utils/response"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	err := h.service.CreateUser(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		response.WriteError(w,
			http.StatusInternalServerError,
			fmt.Sprintf("failed to create new user [%s]", err.Error()),
		)
		return
	}

	response.WriteJSON(
		w,
		http.StatusCreated,
		struct {
			Message string `json:"message"`
		}{
			Message: "user created successfully",
		})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err := h.service.UpdateUser(r.Context(), id, req.Name, req.Email)
	if err != nil {
		response.WriteError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("failed to update user [%s]", err.Error()),
		)
		return
	}

	response.WriteJSON(
		w,
		http.StatusOK,
		struct {
			Message string `json:"message"`
		}{
			Message: "user updated successfully",
		})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.service.DeleteUser(r.Context(), id)
	if err != nil {
		response.WriteError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("failed to delete user [%s]", err.Error()),
		)
		return
	}

	response.WriteJSON(
		w,
		http.StatusOK,
		struct {
			Message string `json:"message"`
		}{
			Message: "user deleted successfully",
		},
	)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		response.WriteError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("failed to fetch user [%s]", err.Error()),
		)
		return
	}

	userData := dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	response.WriteJSON(
		w,
		http.StatusOK,
		struct {
			Message string `json:"message"`
			Data    any    `json:"data"`
		}{
			Message: "user fetched successfully",
			Data:    userData,
		})
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers(r.Context())
	if err != nil {
		response.WriteError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("failed to fetch users [%s]", err.Error()),
		)
	}

	data := make([]dto.UserResponse, 0, len(users))

	for _, user := range users {
		data = append(data, dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	response.WriteJSON(
		w,
		http.StatusOK,
		struct {
			Message string `json:"message"`
			Data    any    `json:"data"`
		}{
			Message: "users fetched successfully",
			Data:    data,
		})
}
