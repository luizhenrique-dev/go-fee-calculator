package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfIGetAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validade(), "ID is required")
}

func TestIfIGetAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validade(), "invalid price")
}

func TestIfIGetAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{ID: "123", Price: 10.00}
	assert.Error(t, order.Validade(), "invalid tax")
}

func TestWithAllValidParams(t *testing.T) {
	order := Order{ID: "123", Price: 10.00, Tax: 1.00}
	assert.NoError(t, order.Validade())
	assert.Equal(t, 10.00, order.Price)
	assert.Equal(t, 1.00, order.Tax)

	order.CalculateFinalPrice()
	assert.Equal(t, 11.00, order.FinalPrice)
}
