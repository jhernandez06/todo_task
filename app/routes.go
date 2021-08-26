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
	root.GET("/delete/{task_id}", actions.Delete)
	root.GET("/show/{task_id}", actions.ShowTask)
	root.GET("/edit/{task_id}", middleware.EditTaskMW(actions.EditTask))
	root.PUT("/update/{task_id}", actions.UpdateTask)
	root.PUT("/updateCheck/{task_id}", actions.UpdateTaskCheck)

	root.GET("/user/new", actions.NewUser)
	root.POST("/user/create", actions.CreateUser)
	root.GET("/user/list", actions.UsersList)
	root.GET("/user/show/{user_id}", actions.ShowUser)
	root.GET("/user/delete/{user_id}", actions.DestroyUser)
	root.GET("/user/edit/{user_id}", actions.EditUser)
	root.PUT("/user/update/{user_id}", actions.UpdateUser)
	root.PUT("/user/active/{user_id}", actions.UpdateUserActive)
	root.ServeFiles("/", base.Assets)
}
