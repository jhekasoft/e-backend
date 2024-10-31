package handler

import (
	internalHttp "e-backend/internal/http"
	internalModels "e-backend/internal/models"
	"e-backend/internal/modules/auth/models"
	"e-backend/internal/modules/auth/service"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	Handler struct {
		service *service.Service
	}

	CreateUserRequest struct {
		Username string `validate:"required"`
		Email    string `validate:"required,email"`
		Name     string `validate:"required"`
		Password string `validate:"required,gte=6"`
	}

	GetUserResponse struct {
		Data models.User
	}

	CreateUserResponse GetUserResponse

	SignInRequest struct {
		Credential string `validate:"required"`
		Password   string `validate:"required,gte=6"`
	}

	SignInResponse struct {
		Token string
		Data  models.User
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

func (h *Handler) SignIn(c echo.Context) error {
	req := new(SignInRequest)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	token, user, err := h.service.SignIn(req.Credential, req.Password)
	if errors.Is(err, service.ErrAuthUserNotFound) {
		return internalHttp.NewCustomValidationFieldError(
			"User is not found or password is incorrect",
			"Credential",
		)
	}
	if err != nil {
		return err
	}

	resp := SignInResponse{Token: token, Data: *user}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) CurrentUser(c echo.Context) error {
	authToken := c.Get("user").(*jwt.Token)
	fmt.Println(authToken)

	claims := authToken.Claims.(*internalModels.AuthClaims)
	subject := claims.Subject
	id, err := strconv.ParseUint(subject, 10, 32)
	if err != nil {
		return err
	}

	item, err := h.service.Get(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	if err != nil {
		return err
	}

	resp := GetUserResponse{Data: *item}

	return c.JSON(http.StatusOK, resp)
}
