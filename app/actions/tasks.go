package actions

import (
	"TodoList/app/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
)

// Show List
func TasksList(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	status := c.Param("check_complet")

	q := tx.Q()

	if status == "true" || status == "false" {
		q.Where("check_complet = ?", status)
	}

	tasks := models.Tasks{}

	if err := q.Order("limit_data asc").All(&tasks); err != nil {
		return err
	}

	c.Set("tasks", tasks)
	return c.Render(http.StatusOK, r.HTML("task/index.plush.html"))
}

// New task
func NewTask(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	users := models.Users{}
	user := models.User{}
	q := tx.Q()
	q.Where("active = true")

	if err := q.Order("first_name asc").All(&users); err != nil {
		return err
	}

	UsersList := []map[string]interface{}{}

	for _, user := range users {
		oneUser := map[string]interface{}{
			user.FirstName + " " + user.LastName: user.ID,
		}
		UsersList = append(UsersList, oneUser)

	}
	c.Set("usersList", UsersList)
	c.Set("user", user)
	c.Set("users", users)
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
		return c.Redirect(http.StatusNotFound, "/")
	}

	c.Set("task", task)
	return c.Render(http.StatusOK, r.HTML("task/show.plush.html"))
}

func EditTask(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	task := models.Task{}
	users := models.Users{}
	user := models.User{}

	taskID := c.Param("task_id")

	if err := tx.Find(&task, taskID); err != nil {
		c.Flash().Add("danger", "a task with that ID was not found")
		return c.Redirect(http.StatusNotFound, "/")
	}
	if err := tx.All(&users); err != nil {
		return err
	}

	UsersList := []map[string]interface{}{}

	for _, user := range users {
		oneUser := map[string]interface{}{
			user.FirstName + " " + user.LastName: user.ID,
		}
		UsersList = append(UsersList, oneUser)

	}
	c.Set("usersList", UsersList)
	c.Set("user", user)
	c.Set("users", users)

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
		return c.Redirect(404, "/")
	}
	if err := c.Bind(&task); err != nil {
		return err
	}
	users := models.Users{}
	user := models.User{}
	if err := tx.All(&users); err != nil {
		return err
	}
	UsersList := []map[string]interface{}{}
	for _, user := range users {
		oneUser := map[string]interface{}{
			user.FirstName + " " + user.LastName: user.ID,
		}
		UsersList = append(UsersList, oneUser)
	}
	c.Set("usersList", UsersList)
	c.Set("user", user)
	c.Set("users", users)
	verrs := task.Validate()
	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("task", task)
		return c.Render(http.StatusSeeOther, r.HTML("task/edit.plush.html"))
	}
	if err := tx.Update(&task); err != nil {
		return err
	}
	c.Flash().Add("success", "task updated success")
	return c.Redirect(http.StatusSeeOther, "/")

}

//delete
func DestroyTask(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	task := models.Task{}
	taskID, _ := uuid.FromString(c.Param("task_id"))
	if err := tx.Find(&task, taskID); err != nil {
		c.Flash().Add("danger", "no task found with that ID")
		return c.Redirect(404, "/")
	}
	//taskdelete := models.Task{ID: taskID}
	if err := tx.Destroy(&task); err != nil {
		return err
	}
	c.Flash().Add("success", "task destroyed success")
	return c.Redirect(http.StatusSeeOther, "/")
}

func Delete(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	taskID, _ := uuid.FromString(c.Param("task_id"))

	if c.Param("task_id") == "" {
		c.Flash().Add("danger", "no task found with that ID")
		return c.Redirect(404, "/")
	}
	taskdelete := &models.Task{ID: taskID}
	if err := tx.Destroy(taskdelete); err != nil {
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
