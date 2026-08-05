package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type LogEntry struct {
	Timestamp     string `json:"timestamp"`
	Level         string `json:"level"`
	CorrelationID string `json:"correlation_id"`
	Service       string `json:"service"`
	Message       string `json:"message"`
	Method        string `json:"method,omitempty"`
	Path          string `json:"path,omitempty"`
	StatusCode    int    `json:"status_code,omitempty"`
	LatencyMS     int64  `json:"latency_ms,omitempty"`
}

type ILogger interface {
	LogHTTPOriginRequest(r *http.Request)
	LogHTTPProxyRTequest(r *http.Request)
}

type Logger struct {
	serviceName string
}

// LogHTTPRequest implements ILogger.
func (l *Logger) LogHTTPOriginRequest(req *http.Request) {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%v %v %v\n\n", req.Method, req.URL.Path, req.Proto)

	for key, val := range req.Header {
		fmt.Fprintf(&sb, "%v: %v\n", key, strings.Join(val, ","))
	}

	if req.Body != nil {
		defer req.Body.Close()
		buf, _ := io.ReadAll(req.Body)
		req_rc := io.NopCloser(bytes.NewBuffer(buf))
		req.Body = req_rc

		sb.WriteString("\n")
		sb.Write(buf)
	}

	log.Printf("\n%v", sb.String())
}

// LogHTTPResponse implements ILogger.
func (l *Logger) LogHTTPProxyRTequest(req *http.Request) {
	var sb strings.Builder
	resp, _ := http.DefaultTransport.RoundTrip(req)

	fmt.Fprintf(&sb, "%v\n", resp.StatusCode)

	for key, val := range resp.Header {
		fmt.Fprintf(&sb, "%v: %v\n", key, strings.Join(val, ","))
	}

	if resp.Body != nil {
		defer resp.Body.Close()
		buf, _ := io.ReadAll(resp.Body)
		req_rc := io.NopCloser(bytes.NewBuffer(buf))
		resp.Body = req_rc

		sb.WriteString("\n")
		sb.Write(buf)
	}

}

func NewLogger(serviceName string) ILogger {
	return &Logger{
		serviceName: serviceName,
	}
}
