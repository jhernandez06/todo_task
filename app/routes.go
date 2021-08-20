package app

import (
	base "TodoList"
	"TodoList/app/actions"
	"TodoList/app/middleware"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.Transaction)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)
	root.Use(middleware.NTasksIncomplet)
	root.Use(middleware.TimeMW)

	root.GET("/", actions.TasksList)
	root.GET("/new", actions.NewTask)
	root.POST("/create", actions.CreateTask)
	root.GET("/show/{task_id}", actions.ShowTask)
	root.GET("/edit/{task_id}", middleware.EditTaskMW(actions.EditTask))
	root.PUT("/update/{task_id}", actions.UpdateTask)
	root.GET("/delete/{task_id}/", actions.DestroyTask)
	root.PUT("/updateCheck/{task_id}", actions.UpdateTaskCheck)
	root.ServeFiles("/", base.Assets)
}
