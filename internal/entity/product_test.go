package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {

	p, err := NewProduct("Product 1", 1000.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.CreatedAt)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 1000.0, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	_, err := NewProduct("", 1000)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	_, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	_, err := NewProduct("Product 1", -1)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
}
