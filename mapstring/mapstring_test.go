package mapstring

import (
	"testing"

	"github.com/segmentio/testdemo/suite"
)

func TestSuite(t *testing.T) {
	rem := make(Remote)
	suite.Test(t, suite.Config{SP: rem})
}
