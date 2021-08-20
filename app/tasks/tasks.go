package tasks

import (
	"TodoList/app"

	"github.com/gobuffalo/buffalo"
)

// Init the tasks with some common tasks that come from
// grift
func init() {
	buffalo.Grifts(app.New())

	// type allTasks []models.Task

	// var Tasks_todo = allTasks{
	// 	{
	// 		Title:       "test create 1",
	// 		LimitData:   t,
	// 		Description: "Testing",
	// 	},
	// 	{
	// 		Title:       "test create 2",
	// 		LimitData:   t,
	// 		Description: "Testing",
	// 	},
	// 	{
	// 		Title:       "test create 3",
	// 		LimitData:   t,
	// 		Description: "Testing",
	// 	},
	// 	{
	// 		Title:       "test create 4",
	// 		LimitData:   t,
	// 		Description: "Testing",
	// 	},
	// }
	// //return models.DB.Create(Tasks_todo)
}
