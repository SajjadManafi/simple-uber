package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	testCases := []struct {
		min, max int64
	}{
		{0, 0},
		{0, 1},
		{0, 10},
		{-20, -10},
		{-20, 20},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d-%d", tc.min, tc.max), func(t *testing.T) {
			r := RandomInt(tc.min, tc.max)
			result := tc.min <= r && r <= tc.max
			require.True(t, result)
		})
	}
}

func TestRandomString(t *testing.T) {
	testCases := []struct {
		length int
	}{
		{0},
		{67},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.length), func(t *testing.T) {
			r := RandomString(tc.length)
			require.Equal(t, tc.length, len(r))
		})
	}
}

func TestRandomUsername(t *testing.T) {
	r := RandomUsername()
	require.Equal(t, 6, len(r))
}

func TestRandomEmail(t *testing.T) {
	r := RandomEmail()
	b := r[6:] == "@email.com"
	require.Equal(t, 16, len(r))
	require.True(t, b)
}
