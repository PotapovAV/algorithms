package calculations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHornerScheme(t *testing.T) {
	coefficients := []float64{2.0, 4.0, 7.0, 3.0}
	assert.Equal(t, float64(3742), HornerScheme(coefficients, 10), "invalid result of polynomial evaluation")
}
