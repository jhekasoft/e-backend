package handler

import (
	"e-backend/internal/crud"
	"e-backend/modules/sum/models"
	"e-backend/modules/sum/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

type ArticleListRequest struct {
	Offset int    `query:"Offset"`
	Limit  int    `query:"Limit"`
	Search string `query:"Search"`
}

func (req ArticleListRequest) ToFilter() models.ArticleListFilter {
	return models.ArticleListFilter{
		ListFilter: crud.ListFilter{
			Offset: req.Offset,
			Limit:  req.Limit,
		},
		Search: req.Search,
	}
}

type ArticleListResponse struct {
	Data   []models.Article
	Offset int
	Limit  int
	Total  int64
}

type ArticleWordResponse struct {
	Data         *models.Article
	Alternatives []string
}

func (h *Handler) GetList(c echo.Context) error {
	var req ArticleListRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	filter := req.ToFilter()

	list, total, err := h.service.GetManyWithTotal(filter)
	if err != nil {
		return err
	}

	resp := ArticleListResponse{
		Data:   list,
		Offset: filter.GetOffset(),
		Limit:  filter.GetLimit(),
		Total:  total,
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetWord(c echo.Context) error {
	// Get Word parameter
	wordParam := c.Param("word")
	if len(wordParam) < 1 {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}

	item, alternatives, err := h.service.GetWordOrAlternatives(wordParam)
	if err != nil {
		return err
	}

	if item == nil && len(alternatives) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}

	resp := ArticleWordResponse{
		Data:         item, // TODO: make mapping to the API type
		Alternatives: alternatives,
	}
	return c.JSON(http.StatusOK, resp)
}
