package slice

import (
	"testing"

	"github.com/segmentio/testdemo/suite"
)

func TestSuite(t *testing.T) {
	loc := make(Local, 0)
	suite.Test(t, suite.Config{SP: &loc})
}
