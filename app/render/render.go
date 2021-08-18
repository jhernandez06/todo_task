package render

import (
	base "TodoList"

	"github.com/gobuffalo/buffalo/render"
)

// Engine for rendering across the app, it provides
// the base for rendering HTML, JSON, XML and other formats
// while also defining thing like the base layout.
var Engine = render.New(render.Options{
	HTMLLayout:   "application.plush.html",
	TemplatesBox: base.Templates,
	AssetsBox:    base.Assets,
	Helpers:      Helpers,
})

// Helpers available for the plush templates, there are
// some helpers that are injected by Buffalo but this is
// the list of custom Helpers.
var Helpers = map[string]interface{}{
	// partialFeeder is the helper used by the render engine
	// to find the partials that will be used, this is important
	"partialFeeder": base.Templates.FindString,
	"FormatDate":    FormatDate,
	"Icon":          Icon,
	"status":        CheckStatus,
	//"incomplet":     Incomplet,
	"addTask": AddTask,
}

func FormatDate(t bool) string {
	var status string
	if t {
		status = "Was completed on "
	} else {
		status = "needs to be completed on "
	}
	return status
}
func Icon(k string) string {
	var icon string
	if k == "info" {
		icon = "#info-fill"
	} else if k == "danger" {
		icon = "#exclamation-triangle-fill"
	} else if k == "success" {
		icon = "#check-circle-fill"
	}
	return icon
}
func CheckStatus(x string, y string) string {

	if x == y {
		y = "font-weight-lighter"
	}
	return y
}
func AddTask(x string) string {
	var y string
	if x == "true" {
		y = "d-none"
	}
	return y
}
