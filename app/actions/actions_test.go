package actions_test

import (
	"TodoList/app"
	"testing"

	"github.com/gobuffalo/suite/v3"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	bapp := app.New()
	/*if err != nil {
		t.Error(err)
		t.FailNow()
	}*/

	as := &ActionSuite{suite.NewAction(bapp)}
	suite.Run(t, as)
}
