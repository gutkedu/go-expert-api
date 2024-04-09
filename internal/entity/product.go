package entity

import (
	"errors"
	"time"

	"github.com/gutkedu/go-expert-api/pkg/entity"
)

var (
	ErrRequiredId    = errors.New("id is required")
	ErrNameRequired  = errors.New("name is required")
	ErrInvalidId     = errors.New("id is invalid")
	ErrPriceRequired = errors.New("price is required")
	ErrInvalidPrice  = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	p := &Product{
		ID:        entity.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if err := p.Validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrRequiredId
	}

	if _, err := entity.ParseId(p.ID.String()); err != nil {
		return ErrInvalidId
	}

	if p.Name == "" {
		return ErrNameRequired
	}

	if p.Price == 0 {
		return ErrPriceRequired
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
