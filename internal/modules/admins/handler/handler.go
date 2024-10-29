package handler

import (
	"e-backend/internal/modules/admins/models"
	"e-backend/internal/modules/admins/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	Handler struct {
		service *service.Service
	}

	AdminListFilter struct {
		Offset int               `query:"Offset"`
		Limit  int               `query:"Limit"`
		Role   *models.AdminRole `query:"Role"`
		Search string            `query:"Search"`
	}

	CreateAdminRequest struct {
		Username string `validate:"required"`
		Name     string
		Role     models.AdminRole `validate:"required"`
		Password string           `validate:"required"`
	}

	UpdateAdminRequest struct {
		Name string           `validate:"required"`
		Role models.AdminRole `validate:"required"`
	}

	GetAdminResponse struct {
		Data models.Admin
	}

	CreateAdminResponse GetAdminResponse

	UpdateAdminResponse GetAdminResponse

	AdminListResponse struct {
		Data   []models.Admin
		Offset int
		Limit  int
		Total  int64
	}
)

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) CreateItem(c echo.Context) error {
	req := new(CreateAdminRequest)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	item := models.Admin{
		Username: req.Username,
		Name:     req.Name,
		Role:     req.Role,
		Password: req.Password, // TODO: add hashing
	}
	createdItem, err := h.service.Create(item)
	if err != nil {
		return err
	}

	resp := CreateAdminResponse{Data: *createdItem}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetList(c echo.Context) error {
	req := new(AdminListFilter)
	err := c.Bind(req)
	if err != nil {
		return err
	}

	filter := models.AdminListFilter{
		Offset: req.Offset,
		Limit:  req.Limit,
		Role:   req.Role,
		Search: req.Search,
	}

	list, total, err := h.service.GetManyWithTotal(filter)
	if err != nil {
		return err
	}

	resp := AdminListResponse{
		Data:   list,
		Offset: req.Offset,
		Limit:  req.Limit,
		Total:  total,
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetItem(c echo.Context) error {
	// Get ID parameter
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
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

	resp := GetAdminResponse{Data: *item}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateItem(c echo.Context) error {
	// Get ID parameter
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return err
	}

	req := new(UpdateAdminRequest)
	err = c.Bind(req)
	if err != nil {
		return err
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	item := models.Admin{
		Name: req.Name,
		Role: req.Role,
	}

	updatedItem, err := h.service.Update(uint(id), item)
	if err != nil {
		return err
	}

	resp := UpdateAdminResponse{Data: *updatedItem}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteItem(c echo.Context) error {
	// Get ID parameter
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return err
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
