package actions

import (
	"TodoList/app/models"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
)

func NewUser(c buffalo.Context) error {
	c.Set("user", models.User{})
	return c.Render(http.StatusOK, r.HTML("user/new.plush.html"))
}

func CreateUser(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	verrs, err := user.Validate(tx)
	if err != nil {
		return err
	}
	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("user", user)
		return c.Render(http.StatusOK, r.HTML("user/new.plush.html"))
	}
	user.Email = strings.ToLower(user.Email)
	if err := tx.Create(&user); err != nil {
		return err
	}
	c.Flash().Add("success", "user created successfully")
	return c.Redirect(http.StatusSeeOther, "/user/list")
}
func UsersList(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	users := models.Users{}

	if err := tx.Order("first_name asc").All(&users); err != nil {
		return err
	}
	c.Set("users", users)
	return c.Render(http.StatusOK, r.HTML("user/list.plush.html"))
}
func ShowUser(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := models.User{}
	userID := c.Param("user_id")
	if err := tx.Find(&user, userID); err != nil {
		c.Flash().Add("danger", "a task with that ID was not found")
		return c.Redirect(http.StatusNotFound, "/user/list")
	}
	c.Set("user", user)
	return c.Render(http.StatusOK, r.HTML("user/show.plush.html"))
}
func EditUser(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := models.User{}
	userID := c.Param("user_id")
	if err := tx.Find(&user, userID); err != nil {
		c.Flash().Add("danger", "a user with that ID was not found")
		return c.Redirect(http.StatusNotFound, "/user/list")
	}
	c.Set("user", user)
	return c.Render(http.StatusOK, r.HTML("user/edit.plush.html"))
}
func UpdateUser(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := models.User{}
	userID := c.Param("user_id")
	if err := tx.Find(&user, userID); err != nil {
		c.Flash().Add("danger", "a user with that ID was not found")
		return c.Redirect(404, "/user/list")
	}
	if err := c.Bind(&user); err != nil {
		return err
	}
	verrs, err := user.Validate(tx)
	if err != nil {
		return err
	}
	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("user", user)
		return c.Render(http.StatusSeeOther, r.HTML("user/edit.plush.html"))
	}
	if err := tx.Update(&user); err != nil {
		return err
	}
	c.Flash().Add("success", "user updated successfully")
	return c.Redirect(http.StatusSeeOther, "/user/list")

}
func DestroyUser(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := models.User{}
	userID, _ := uuid.FromString(c.Param("user_id"))
	if err := tx.Find(&user, userID); err != nil {
		c.Flash().Add("danger", "no user found with that ID")
		return c.Redirect(404, "/user/list")
	}
	if err := tx.Destroy(&user); err != nil {
		return err
	}
	c.Flash().Add("success", "user destroyed successfully")
	return c.Redirect(http.StatusSeeOther, "/user/list")
}

func UpdateUserActive(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := models.User{}
	userID := c.Param("user_id")
	if err := tx.Find(&user, userID); err != nil {
		return err
	}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if !(user.Active) {
		user.Active = true
		c.Flash().Add("info", "User Activated successfully")
	} else if user.Active {
		user.Active = false
		c.Flash().Add("info", "User disable")
	}
	if err := tx.Update(&user); err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/user/list")
}
