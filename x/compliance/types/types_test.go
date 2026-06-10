package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidFamilyID(t *testing.T) {
	testCases := []struct {
		name    string
		id      string
		isValid bool
	}{
		{
			name:    "valid family id with numbers only",
			id:      "123456",
			isValid: true,
		},
		{
			name:    "valid family id with letters only",
			id:      "abc",
			isValid: true,
		}, {
			name:    "valid family id with numbers and letter",
			id:      "123abc",
			isValid: true,
		},
		{
			name:    "invalid family id with special characters",
			id:      "123abc!",
			isValid: false,
		},
		{
			name:    "invalid family id with hyphen",
			id:      "123-abc",
			isValid: false,
		},
		{
			name:    "empty family id",
			id:      "",
			isValid: false,
		},
		{
			name:    "invalid family id with only spaces",
			id:      "  ",
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			valid := IsValidFamilyID(tc.id)
			require.Equal(t, tc.isValid, valid)
		})
	}
}
