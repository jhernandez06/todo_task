package actions

import (
	"TodoList/app/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func TasksList(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	status := c.Param("check_complet")

	q := tx.Q()

	if status != "" {
		q.Where("check_complet = ?", status)
	}

	tasks := models.Tasks{}

	if err := q.Order("limit_data asc").All(&tasks); err != nil {
		return err
	}

	c.Set("tasks", tasks)
	return c.Render(http.StatusOK, r.HTML("task/index.plush.html"))
}

func NewTask(c buffalo.Context) error {
	c.Set("task", models.Task{})
	return c.Render(http.StatusOK, r.HTML("task/new.plush.html"))
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

		return c.Render(http.StatusOK, r.HTML("task/new.plush.html"))
	}

	if err := tx.Create(&task); err != nil {
		return err
	}
	c.Flash().Add("success", "task created success")
	return c.Redirect(http.StatusSeeOther, "/")
}

func ShowTask(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	task := models.Task{}
	taskID := c.Param("task_id")

	if err := tx.Find(&task, taskID); err != nil {
		c.Flash().Add("danger", "a task with that ID was not found")
		return c.Redirect(302, "/")
	}

	c.Set("task", task)
	return c.Render(http.StatusOK, r.HTML("task/show.plush.html"))
}

func EditTask(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	task := models.Task{}
	taskID := c.Param("task_id")

	if err := tx.Find(&task, taskID); err != nil {
		c.Flash().Add("danger", "a task with that ID was not found")
		c.Redirect(http.StatusSeeOther, "/")
	}

	c.Set("task", task)
	return c.Render(http.StatusOK, r.HTML("task/edit.plush.html"))
}

//Update task
func UpdateTask(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	task := models.Task{}
	taskID := c.Param("task_id")

	if err := tx.Find(&task, taskID); err != nil {
		c.Flash().Add("danger", "a task with that ID was not found")
		return c.Redirect(302, "/")
	}

	if err := c.Bind(&task); err != nil {
		return err
	}
	verrs := task.Validate()
	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("task", task)

		return c.Render(http.StatusOK, r.HTML("task/edit.plush.html"))
	}
	if err := tx.Update(&task); err != nil {
		return err
	}
	if task.CheckComplet {
		c.Flash().Add("danger", "cannot update a completed task")
		return c.Redirect(404, "/")
	}
	c.Flash().Add("success", "task updated success")

	return c.Redirect(http.StatusSeeOther, "/")

}

//delete
func DestroyTask(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)

	task := models.Task{}
	taskID := c.Param("task_id")

	if err := tx.Find(&task, taskID); err != nil {
		c.Flash().Add("danger", "no task found with that ID")
		return c.Redirect(http.StatusSeeOther, "/")
	}
	if err := tx.Destroy(&task); err != nil {
		return err
	}
	c.Flash().Add("success", "task destroyed success")

	return c.Redirect(http.StatusSeeOther, "/")

}

// CheckComplet
func UpdateTaskCheck(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	task := models.Task{}
	taskID := c.Param("task_id")

	if err := tx.Find(&task, taskID); err != nil {
		return err
	}

	if err := c.Bind(&task); err != nil {
		return err
	}

	if !(task.CheckComplet) {
		task.CheckComplet = true
		c.Flash().Add("info", "task completed success, Congratulations")
	} else if task.CheckComplet {
		task.CheckComplet = false
		c.Flash().Add("info", "the task returned to incomplete tasks")
	}
	if err := tx.Update(&task); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
