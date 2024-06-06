package slicestudy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTransitiveSliceAppend shows that manipulating slices of slices affects the shared backing array of all 'underlying' slices too.
func TestTransitiveSliceAppend(t *testing.T) {
	// arrange
	s1 := []int{1, 2, 3}
	s2 := s1[1:2]

	// act
	s3 := append(s2, 4)

	// assert
	assert.EqualValues(t, s3, []int{2, 4})
	assert.EqualValues(t, []int{1, 2, 3, 4}, s1)
}
