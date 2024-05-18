package main

import (
	// "encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGetLeaves(t *testing.T) {
	app := setup()

	test := struct {
		description  string
		expectStatus int
	}{
		description:  "Valid Get",
		expectStatus: fiber.StatusOK,
	}

	t.Run(test.description, func(t *testing.T) {
		// reqBody, _ := json.Marshal(test.reqBody)
		req := httptest.NewRequest("GET", "/leaves", nil)
		res, _ := app.Test(req)

		assert.Equal(t, test.expectStatus, res.StatusCode)
	})
}

func TestGetLeave(t *testing.T) {
	app := setup()

	tests := []struct {
		description  string
		id int
		expectStatus int
	}{
		{
			description:  "Valid Get",
			id: 1,
			expectStatus: fiber.StatusOK,
		},
	}

	for _, tc := range tests{
		t.Run(tc.description, func(t *testing.T) {
			// reqBody, _ := json.Marshal(test.reqBody)
			req := httptest.NewRequest("GET", "/leaves", nil)
			res, _ := app.Test(req)
	
			assert.Equal(t, tc.expectStatus, res.StatusCode)
		})
	}
}
