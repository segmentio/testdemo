package partyparrot

import (
	"fmt"
	"io"
	"time"
)

var now func() time.Time = time.Now

func PrintRFC3339(w io.Writer) {
	fmt.Fprint(w, now().Format(time.RFC3339))
}
