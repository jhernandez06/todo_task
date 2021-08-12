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

	//root.GET("/", home.Index)
	root.GET("/tasks", actions.TasksList)

	root.GET("/tasks/new", actions.NewTask)
	root.POST("/tasks/create", actions.CreateTask)
	root.GET("/tasks/show/{task_id}", actions.ShowTask)
	root.GET("/tasks/edit/{task_id}", actions.EditTask)
	root.PUT("/tasks/update/{task_id}", actions.UpdateTask)
	root.GET("/tasks/delete/{task_id}", actions.DestroyTask)
	root.PUT("/tasks/updateCheck/{task_id}", actions.UpdateTaskCheck)
	root.ServeFiles("/", base.Assets)
}
