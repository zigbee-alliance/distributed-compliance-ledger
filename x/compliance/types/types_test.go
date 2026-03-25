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
			id:      "FAM123456",
			isValid: true,
		},
		{
			name:    "valid family id with letters only",
			id:      "FAMabc",
			isValid: true,
		}, {
			name:    "valid family id with numbers and letter",
			id:      "FAM123abc",
			isValid: true,
		},
		{
			name:    "invalid family id with special characters",
			id:      "FAM123abc!",
			isValid: false,
		},
		{
			name:    "invalid family id without 'FAM' prefix",
			id:      "123abc",
			isValid: false,
		},
		{
			name:    "invalid family id with hyphen",
			id:      "FAM-123",
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
