package uuid

import (
	"testing"
)

func TestNewV4(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	u := NewV4()

	if &u == nil {
		t.Error("NewV4() returned nil")
	}
}
