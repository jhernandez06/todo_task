package actions_test

import (
	"TodoList/app/models"
	"fmt"
	"time"
)

func (as *ActionSuite) Test_Task_New() {
	res := as.HTML("/").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_Users_Create() {
	count, err := as.DB.Count("tasks")
	as.NoError(err)
	as.Equal(0, count)

	str := "2014-11-12T11:45:26.371Z"
	t, err := time.Parse(time.RFC3339, str)

	if err != nil {
		fmt.Println(err)
	}

	u := &models.Task{
		Title:       "test create",
		LimitData:   t,
		Description: "Testing",
	}

	res := as.HTML("/create").Post(u)
	as.Equal(303, res.Code)
	//as.Equal("Learn Go", u.Title)
	count, err = as.DB.Count("tasks")
	as.NoError(err)
	as.Equal(1, count)
}

func (as *ActionSuite) Test_Task_Update() {
	str := "2014-11-12T11:45:26.371Z"
	t, err := time.Parse(time.RFC3339, str)

	if err != nil {
		fmt.Println(err)
	}

	u := &models.Task{
		Title:       "test create",
		LimitData:   t,
		Description: "Testing",
	}
	verrs, err := as.DB.ValidateAndCreate(u)
	as.NoError(err)
	as.False(verrs.HasAny())

	res := as.HTML("/edit/%s", u.ID).Put(&models.Task{Title: "Learn Go"})
	as.Equal(200, res.Code)

	err = as.DB.Reload(u)
	as.NoError(err)
	as.Equal("Learn Go", u.Title)

}
