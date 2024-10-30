package handler

import (
	"e-backend/internal/modules/auth/models"
	"e-backend/internal/modules/auth/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Handler struct {
		service *service.Service
	}

	CreateUserRequest struct {
		Username string `validate:"required"`
		Email    string `validate:"required"`
		Name     string `validate:"required"`
		Password string `validate:"required"`
	}

	CreateUserResponse struct {
		Data models.User
	}
)

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) CreateItem(c echo.Context) error {
	req := new(CreateUserRequest)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	item := models.User{
		Username: req.Username,
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password, // TODO: add hashing
	}
	createdItem, err := h.service.Create(item)
	if err != nil {
		return err
	}

	resp := CreateUserResponse{Data: *createdItem}
	return c.JSON(http.StatusOK, resp)
}
