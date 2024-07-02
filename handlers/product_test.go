package handlers

import (
	"bytes"
	"encoding/json"
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
	// args := m.Called(id)
	return types.Product{}, nil
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
