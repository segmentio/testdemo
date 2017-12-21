package partyparrot_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/segmentio/testdemo/partyparrot"
)

func TestParty(t *testing.T) {
	d := time.Date(1906, time.December, 9, 0, 0, 0, 0, time.UTC)
	*partyparrot.Now = func() time.Time {
		return d
	}
	t.Run("times are equal", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		partyparrot.PrintRFC3339(b)
		if d.Format(time.RFC3339) != b.String() {
			t.Fatalf("no match. expected: %s, got: %s", d.Format(time.RFC3339), b.String())
		}
	})
}
