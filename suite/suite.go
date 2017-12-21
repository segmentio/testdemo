package suite

import (
	"testing"

	"github.com/segmentio/testdemo"
)

type Config struct {
	SP settings.SettingsProvider
}

func Test(t *testing.T, conf Config) {
	tests := []struct {
		title string
		run   func(t *testing.T, c Config)
	}{
		{"should read and write settings", testBasic},
	}
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			test.run(t, conf)
		})
	}
}

func testBasic(t *testing.T, c Config) {
	err := c.SP.Set("foo", "bar")
	if err != nil {
		t.Errorf("unexpected err: %v", err)
		return
	}
	actual, err := c.SP.Get("foo")
	if err != nil {
		t.Errorf("unexpected err: %v", err)
		return
	}
	if actual != "bar" {
		t.Errorf("reading settings. expected: %v, actual: %v", "bar", actual)
	}
}
