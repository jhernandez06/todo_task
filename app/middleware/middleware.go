// middleware package is intended to host the middlewares used
// across the app.
package middleware

import (
	"TodoList/app/models"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	tx "github.com/gobuffalo/buffalo-pop/v2/pop/popmw"
	csrf "github.com/gobuffalo/mw-csrf"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
)

var (
	// Transaction middleware wraps the request with a pop
	// transaction that is committed on success and rolled
	// back when errors happen.
	Transaction = tx.Transaction(models.DB())

	// ParameterLogger logs out parameters that the app received
	// taking care of sensitive data.
	ParameterLogger = paramlogger.ParameterLogger

	// CSRF middleware protects from CSRF attacks.
	CSRF = csrf.New
)

func NTasksIncomplet(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {

		tx := models.DB()
		q := tx.Q()
		tasks := models.Tasks{}
		q.Where("check_complet = false")
		if err := q.All(&tasks); err != nil {
			return err
		}
		c.Set("ntasks", len(tasks))
		return next(c)
	}
}
func TimeMW(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		t := time.Now()
		c.Set("Date", t.Format("Monday 02, Jan 2006"))
		return next(c)
	}
}

func EditTaskMW(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {

		tx := models.DB()

		task := models.Task{}
		taskID := c.Param("task_id")

		tx.Find(&task, taskID)
		if task.CheckComplet {
			c.Flash().Add("danger", "cannot edit a completed task")
			c.Redirect(http.StatusSeeOther, "/")
		}

		return next(c)
	}
}
