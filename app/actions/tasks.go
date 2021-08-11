package actions

import (
	"TodoList/app/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func TasksList(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	tasks := models.Tasks{}

	if err := tx.All(&tasks); err != nil {
		return err
	}
	c.Set("ntasks", len(tasks))
	c.Set("tasks", tasks)
	return c.Render(http.StatusOK, r.HTML("todo-tasks/index.plush.html"))
}

func NewTask(c buffalo.Context) error {
	c.Set("task", models.Task{})
	return c.Render(http.StatusOK, r.HTML("todo-tasks/new-task.plush.html"))
}

func CreateTask(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	task := models.Task{}

	if err := c.Bind(&task); err != nil {
		return err
	}

	verrs := task.Validate()
	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("task", task)

		return c.Render(http.StatusOK, r.HTML("todo-tasks/new-task.plush.html"))
	}

	if err := tx.Create(&task); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func ShowTask(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	task := models.Task{}
	taskID := c.Param("task_id")

	if err := tx.Find(&task, taskID); err != nil {
		return err
	}

	c.Set("task", task)
	return c.Render(http.StatusOK, r.HTML("todo-tasks/show-task.plush.html"))
}

func EditTask(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	task := models.Task{}
	taskID := c.Param("task_id")

	if err := tx.Find(&task, taskID); err != nil {
		return err
	}

	c.Set("task", task)
	return c.Render(http.StatusOK, r.HTML("todo-tasks/edit-task.plush.html"))
}

//PENDIENTE
func UpdateTask(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	task := models.Task{}
	taskID := c.Param("task_id")

	if err := tx.Find(&task, taskID); err != nil {
		return err
	}

	if err := c.Bind(&task); err != nil {
		return err
	}
	verrs := task.Validate()
	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("task", task)

		return c.Render(http.StatusOK, r.HTML("todo-tasks/edit-task.plush.html"))
	}
	if err := tx.Update(&task); err != nil {
		return err
	}
	//return c.Redirect(http.StatusSeeOther, "taskUpdatePath()", render.Data{"task_id": task.ID})
	return c.Redirect(http.StatusSeeOther, "/")
}

//delete
func DestroyTask(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)

	task := models.Task{}
	taskID := c.Param("task_id")

	if err := tx.Find(&task, taskID); err != nil {
		return err
	}
	if err := tx.Destroy(&task); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

// tasks completed
func TasksCompleted(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("todo-tasks-complet/index.plush.html"))
}
