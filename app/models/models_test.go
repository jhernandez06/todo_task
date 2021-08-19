package models_test

import (
	"testing"

	"github.com/gobuffalo/suite/v3"
)

type ModelSuite struct {
	*suite.Model
}

func Test_ModelSuite(t *testing.T) {
	suite.Run(t, &ModelSuite{
		Model: suite.NewModel(),
	})
}
