package helper

import (
	"net/http"
	"net/http/httptest"

	"github.com/rs/zerolog"
)

// PerformRequest for testing gin router.
func PerformRequest(r http.Handler, method, path string, headers map[string]string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// LogSink provides an output to attach to Zerolog to catch logs.
type LogSink struct {
	Logs []string
}

func (l *LogSink) Write(p []byte) (n int, err error) {
	l.Logs = append(l.Logs, string(p))
	return len(p), nil
}

// LogHook is used as a Zerolog Hook to catch log events.
type LogHook struct {
	LogEvents []zerolog.Event
}

// Run catches log events.
func (logHook *LogHook) Run(logEvent *zerolog.Event, level zerolog.Level, message string) {
	logHook.LogEvents = append(logHook.LogEvents, *logEvent)
}
