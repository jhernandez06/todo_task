package render_test

import (
	"TodoList/app/render"
	"testing"
)

// func TestStatus(t *testing.T) {
// 	statusTrue := render.Status(true)
// 	if statusTrue != "Was completed on " {
// 		t.Error("Error")
// 		t.Fail()
// 	}
// 	statusFalse := render.Status(false)
// 	if statusFalse != "needs to be completed on " {
// 		t.Error("Error")
// 		t.Fail()
// 	}
// }

func TestIcon(t *testing.T) {
	var icon string
	if icon = render.Icon("info"); icon != "#info-fill" {
		t.Error("Error")
		t.Fail()
	}
	if icon = render.Icon("danger"); icon != "#exclamation-triangle-fill" {
		t.Error("Error")
		t.Fail()
	}
	if icon = render.Icon("success"); icon != "#check-circle-fill" {
		t.Error("Error")
		t.Fail()
	}

}
func TestCheck(t *testing.T) {
	if check := render.CheckStatus("status", "status"); check != "font-weight-lighter" {
		t.Error("Error")
		t.Fail()
	}
}
func TestAddTask(t *testing.T) {
	if add := render.AddTask("true"); add != "d-none" {
		t.Error("Error")
		t.Fail()
	}
}
