package handlers

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"

	"github.com/KurobaneShin/eulabs/db"
	"github.com/KurobaneShin/eulabs/types"
)

type ProductHandler struct {
	db db.DB
}

func NewProductHandler(db db.DB) *ProductHandler {
	return &ProductHandler{db: db}
}

func (h *ProductHandler) HandleGetProduct(c echo.Context) error {
	id := c.Param("id")

	product, err := h.db.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) HandleCreateProduct(c echo.Context) error {
	p := new(types.Product)
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.db.CreateProduct(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, p)
}

func (h *ProductHandler) HandleUpdateProduct(c echo.Context) error {
	p := new(types.Product)
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	val, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	p.Id = val

	err = h.db.UpdateProduct(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}

func (h *ProductHandler) HandleDeleteProduct(c echo.Context) error {
	id := c.Param("id")

	err := h.db.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}
