package utility

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	conf, err := LoadConfig("./banana")

	if conf != nil {
		t.Error("cconf - should not exist")
	}

	if err == nil {
		t.Error("err  - should exist")
	}
}
