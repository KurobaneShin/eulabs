package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/KurobaneShin/eulabs/types"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) CreateProduct(product *types.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockDB) GetProductById(id string) (types.Product, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(types.Product), args.Error(1)
	}
	return types.Product{}, args.Error(1)
}

func (m *MockDB) UpdateProduct(product *types.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockDB) DeleteProduct(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestHandleCreateProduct(t *testing.T) {
	e := echo.New()

	mockDB := new(MockDB)
	productHandler := &ProductHandler{db: mockDB}

	product := &types.Product{
		Title:       "Test Product",
		Description: "This is a test product",
		Price:       100,
	}

	// Setting up the expected calls and returns
	mockDB.On("CreateProduct", product).Return(nil)

	// Create JSON payload
	payload, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// Run the handler
	err := productHandler.HandleCreateProduct(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var createdProduct types.Product
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &createdProduct)) {
			assert.Equal(t, product.Title, createdProduct.Title)
			assert.Equal(t, product.Description, createdProduct.Description)
			assert.Equal(t, product.Price, createdProduct.Price)
		}
	}

	// Assert that the expectations were met
	mockDB.AssertExpectations(t)
}

func TestHandleGetProduct(t *testing.T) {
	e := echo.New()

	// Create an instance of MockDB
	mockDB := new(MockDB)
	productHandler := &ProductHandler{db: mockDB}

	// Test the case where the product is found
	product := types.Product{
		Id:          1,
		Title:       "Test Product",
		Description: "This is a test product",
		Price:       100,
	}
	mockDB.On("GetProductById", "1").Return(product, nil)

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := productHandler.HandleGetProduct(c)

	// Assertions for found case
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var foundProduct types.Product
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &foundProduct)) {
			assert.Equal(t, product.Title, foundProduct.Title)
			assert.Equal(t, product.Description, foundProduct.Description)
			assert.Equal(t, product.Price, foundProduct.Price)
		}
	}

	// Test the case where the product is not found
	mockDB.On("GetProductById", "2").Return(types.Product{}, errors.New("not found"))

	req = httptest.NewRequest(http.MethodGet, "/products/2", nil)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err = productHandler.HandleGetProduct(c)

	// Assertions for not found case
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Contains(t, rec.Body.String(), "Not Found")

	// Assert that the expectations were met
	mockDB.AssertExpectations(t)
}

func TestHandleUpdateProduct(t *testing.T) {
	e := echo.New()

	// Create an instance of MockDB
	mockDB := new(MockDB)
	productHandler := &ProductHandler{db: mockDB}

	// Test case where update is successful
	product := &types.Product{
		Id:          1,
		Title:       "Updated Product",
		Description: "This is an updated product",
		Price:       200,
	}
	mockDB.On("UpdateProduct", product).Return(nil)

	// Prepare request
	payload, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPut, "/products/1", bytes.NewBuffer(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Call handler
	err := productHandler.HandleUpdateProduct(c)

	// Assertions for success case
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var updatedProduct types.Product
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &updatedProduct)) {
			assert.Equal(t, product.Title, updatedProduct.Title)
			assert.Equal(t, product.Description, updatedProduct.Description)
			assert.Equal(t, product.Price, updatedProduct.Price)
		}
	}

	// Test case where ID parsing fails
	req = httptest.NewRequest(http.MethodPut, "/products/abc", bytes.NewBuffer(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("abc")

	err = productHandler.HandleUpdateProduct(c)

	// Assertions for bad request case (invalid ID)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Assert that the expectations were met
	mockDB.AssertExpectations(t)
}

func TestHandleDeleteProduct(t *testing.T) {
	e := echo.New()

	// Create an instance of MockDB
	mockDB := new(MockDB)
	productHandler := &ProductHandler{db: mockDB}

	// Test case where delete is successful
	mockDB.On("DeleteProduct", "1").Return(nil)

	// Prepare request
	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Call handler
	err := productHandler.HandleDeleteProduct(c)

	// Assertions for success case
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// Extract ID from JSON response
		var id string
		err := json.Unmarshal(rec.Body.Bytes(), &id)
		if assert.NoError(t, err) {
			assert.Equal(t, "1", id) // Check extracted ID
		}
	}

	// Assert that the expectations were met
	mockDB.AssertExpectations(t)
}
