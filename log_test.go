package jsonlog

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestLevel(t *testing.T) {
	var buf bytes.Buffer
	l := New(&buf, LevelError)
	l.Info("info", nil)
	if buf.Len() != 0 {
		t.Errorf("Expected no output, got %q", buf.String())
		t.Failed()
	}
	err := errors.New("error lol")
	l.Err(err, nil)
	if strings.Contains(buf.String(), "error lol\n") {
		t.Errorf("Expected error output, got %q", buf.String())
		t.Failed()
	}
}
