package internal

import (
	"encoding/json"
	"testing"

	"github.com/ingvaar/indeks-api/internal/helper"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

type logLine struct {
	Level    string  `json:"level"`
	ReqID    string  `json:"req-id"`
	Method   string  `json:"method"`
	URL      string  `json:"url"`
	Status   int     `json:"status"`
	Size     int     `json:"size"`
	Duration float64 `json:"duration"`
}

func Test_ServerLog(t *testing.T) {
	ls := helper.LogSink{}
	logHook := helper.LogHook{}
	logger := zerolog.New(&ls)
	log.Logger = logger.Hook(&logHook)

	router := gin.New()

	server := NewServer(router, nil, Config{
		DevMode: true,
	})
	server.initMiddleware()

	router.GET("/test", func(ctx *gin.Context) {
		ctx.Status(200)
	})

	for i, c := range []struct {
		Method string
		Path   string
		Params map[string]string
	}{
		{
			Method: "GET",
			Path:   "/test",
			Params: nil,
		},
		{
			Method: "POST",
			Path:   "/test",
			Params: nil,
		},
	} {
		resp := helper.PerformRequest(router, c.Method, c.Path, c.Params)

		amountOfLogEvents := len(logHook.LogEvents)

		assert.Equal(t, i+1, amountOfLogEvents)

		ll := logLine{}
		err := json.Unmarshal([]byte(ls.Logs[i]), &ll)

		assert.NoError(t, err)
		assert.Equal(t, ll.Level, "info")
		assert.Equal(t, ll.Method, c.Method)
		assert.Equal(t, ll.URL, c.Path)
		assert.Equal(t, ll.Status, resp.Code)
		assert.NotEmpty(t, ll.ReqID)
	}
}
