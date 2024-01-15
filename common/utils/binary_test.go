package utils

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFlipByGroups(t *testing.T) {
	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	flipped := FlipByGroups(original, 4)
	fmt.Println(flipped)
	require.EqualValues(t, flipped, []int{13, 14, 15, 16, 9, 10, 11, 12, 5, 6, 7, 8, 1, 2, 3, 4})
}
