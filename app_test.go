package main_test

import (
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"

	"encoding/json"
	"net/http"
	"testing"

	"rr/research-tool"
)

type ResBody struct {
	Message string `json:"message"`
}

func TestIndexRoute(t *testing.T) {
	app := main.App{}
	app.Initialize()

	r := gofight.New()

	r.GET("/").
		Run(app.Router, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, res.Code,
				"It should return 200 status code",
			)

			expected, _ := json.Marshal(&ResBody{"Welcome to go API"})
			actual := res.Body.String()

			assert.Equal(t, string(expected), actual,
				"It should return welcome message object",
			)
		})
}
