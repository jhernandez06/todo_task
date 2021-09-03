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

	//root.Use(middleware.Authorize)
	//root.Use(middleware.SetCurrentUser)

	root.GET("/", actions.Index)
	root.POST("/signin", actions.AuthCreate)
	root.GET("/signout", actions.AuthDestroy)

	root.GET("/tasks/", actions.TasksList)
	root.GET("/tasks/new", middleware.Active(actions.NewTask))
	root.POST("/tasks/create", middleware.Active(actions.CreateTask))
	root.GET("/tasks/delete/{task_id}", middleware.Active(actions.DestroyTask))
	root.GET("/tasks/show/{task_id}", actions.ShowTask)
	root.GET("/tasks/edit/{task_id}", middleware.EditTaskMW(actions.EditTask))
	root.PUT("/tasks/update/{task_id}", middleware.Active(actions.UpdateTask))
	root.PUT("/tasks/updateCheck/{task_id}", middleware.Active(actions.UpdateTaskCheck))

	root.GET("/user/new", actions.NewUser)
	root.GET("/user/newByAdmin", middleware.Admin(actions.NewUserByAdmin))
	root.POST("/user/create", actions.CreateUser)
	root.POST("/user/createByAdmin", middleware.Admin(actions.CreateUserByAdmin))
	root.GET("/user/list", middleware.Admin(actions.UsersList))
	root.GET("/user/show/{user_id}", actions.ShowUser)
	root.GET("/user/delete/{user_id}", actions.DestroyUser)
	root.GET("/user/edit/{user_id}", actions.EditUser)
	root.PUT("/user/update/{user_id}", actions.UpdateUser)
	root.PUT("/user/active/{user_id}", middleware.Admin(actions.UpdateUserActive))

	root.Middleware.Skip(middleware.SetCurrentUser, actions.Index, actions.AuthCreate, actions.AuthDestroy, actions.NewUser, actions.CreateUser)
	root.Middleware.Skip(middleware.Authorize, actions.Index, actions.AuthCreate, actions.AuthDestroy, actions.NewUser, actions.CreateUser)

	root.ServeFiles("/", base.Assets)
}
