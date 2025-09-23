package tests

import (
	"testing"

	"github.com/zawhtetnaing10/FoodDeliveryApp-Backend/handlers"
)

func TestHashPassword(t *testing.T) {
	type testCase struct {
		// Expected inputs
		password string
	}

	t.Run("hash password test", func(t *testing.T) {
		testCases := []testCase{
			{password: "simple password"},
		}

		for _, testCase := range testCases {
			_, err := handlers.HashPassword(testCase.password)
			if err != nil {
				t.Errorf("Failed to hash password %v", err)
			}
		}
	})
}

func TestCheckPasswordHash(t *testing.T) {
	type testCase struct {
		// Input params
		hash     string
		password string
	}

	t.Run("check password hash test", func(t *testing.T) {
		// Hash passwords
		hashPasswordPositive, err := handlers.HashPassword("simplepassword")
		if err != nil {
			t.Errorf("Failed to hash positive password %v", err)
		}
		hashPasswordNegative, err := handlers.HashPassword("otherpassword")
		if err != nil {
			t.Errorf("Failed to hash negative password %v", err)
		}

		// Create test cases
		testCases := []testCase{
			{hash: hashPasswordPositive, password: "simplepassword"},
			{hash: hashPasswordNegative, password: "newpassword"},
		}

		for index, testCase := range testCases {
			err := handlers.CheckPasswordHash(testCase.hash, testCase.password)
			if index == 0 {
				// Positive
				if err != nil {
					t.Errorf("Password and hash are equal and tests must pass")
				}
			} else {
				// Negative
				if err == nil {
					t.Errorf("Password and hash are not equal and tests must fail")
				}
			}
		}
	})
}
