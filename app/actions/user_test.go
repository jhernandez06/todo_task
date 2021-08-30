package actions_test

import (
	"TodoList/app/models"
	"fmt"
)

func (as *ActionSuite) Test_User_New() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)
	res := as.HTML("/user/new").Get()
	as.Equal(200, res.Code)
	body := res.Body.String()
	as.Contains(body, "New User")
}
func (as *ActionSuite) Test_Create_User() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)
	users := models.Users{
		{FirstName: "Javier",
			LastName: "Hernandez",
			Email:    "javier@gmail.com",
			Active:   true},
		{FirstName: "Eduardo",
			LastName: "Gomez",
			Email:    "eduardo@gmail.com",
			Active:   false},
	}
	for _, user := range users {
		resp := as.HTML("/user/create").Post(user)
		as.Equal(303, resp.Code)
	}
	count2, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(2, count2)
}

func (as *ActionSuite) Test_User_List() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)

	users := models.Users{
		{FirstName: "Javier",
			LastName: "Hernandez",
			Email:    "javier@gmail.com",
			Active:   true},
		{FirstName: "Eduardo",
			LastName: "Gomez",
			Email:    "eduardo@gmail.com",
			Active:   false},
	}
	for _, user := range users {
		resp := as.HTML("/user/create").Post(user)
		as.Equal(303, resp.Code)
	}
	res := as.HTML("/user/list").Get()
	as.Equal(200, res.Code)
	body := res.Body.String()
	for _, user := range users {
		as.Contains(body, fmt.Sprintf("%s", user.Email))
	}
	count1, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(2, count1)
}
func (as *ActionSuite) Test_Show_User() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)
	user := &models.User{
		FirstName: "Javier",
		LastName:  "Hernandez",
		Email:     "javier@gmail.com",
		Active:    true}
	verrs, err := as.DB.ValidateAndCreate(user)
	as.NoError(err)
	as.False(verrs.HasAny())
	resp404 := as.HTML("/user/show/javier").Get()
	as.Equal(404, resp404.Code)
	respShow := as.HTML("/user/show/{%s}", user.ID).Get()
	as.Equal(200, respShow.Code)
	body := respShow.Body.String()
	as.Contains(body, fmt.Sprintf("%s", user.Email))

	count1, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count1)
}
func (as *ActionSuite) Test_Delete_User() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)

	user := &models.User{
		FirstName: "Javier",
		LastName:  "Hernandez",
		Email:     "javier@gmail.com",
		Active:    true}
	verrs, err := as.DB.ValidateAndCreate(user)
	as.NoError(err)
	as.False(verrs.HasAny())
	count1, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count1)
	resp404 := as.HTML("/user/delete/javier").Get()
	as.Equal(404, resp404.Code)
	respDelete := as.HTML("/user/delete/{%s}", user.ID).Get()
	as.Equal(303, respDelete.Code)

	count2, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count2)
}
func (as *ActionSuite) Test_Edit_User() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)
	user := &models.User{
		FirstName: "Javier",
		LastName:  "Hernandez",
		Email:     "javier@gmail.com",
		Active:    true}
	verrs, err := as.DB.ValidateAndCreate(user)
	as.NoError(err)
	as.False(verrs.HasAny())
	resp404 := as.HTML("/user/edit/javier").Get()
	as.Equal(404, resp404.Code)
	respEdit := as.HTML("/user/edit/{%s}", user.ID).Get()
	as.Equal(200, respEdit.Code)
	body := respEdit.Body.String()
	as.Contains(body, fmt.Sprintf("%s", user.Email))
	as.Contains(body, "Edit User")
	count1, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count1)
}
func (as *ActionSuite) Test_Update_User() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)
	user := &models.User{
		FirstName: "Javier",
		LastName:  "Hernandez",
		Email:     "javier@gmail.com",
		Active:    false}
	verrs, err := as.DB.ValidateAndCreate(user)
	as.NoError(err)
	as.False(verrs.HasAny())
	respUpdate := as.HTML("/user/update/{%s}", user.ID).Put(&models.User{ID: user.ID, FirstName: "Test", LastName: "Update", Email: "javier@gmail.com", Active: false})
	as.Equal(303, respUpdate.Code)
	err = as.DB.Reload(user)
	as.NoError(err)
	as.Equal("Update", user.LastName)
	resp404 := as.HTML("/user/update/javier").Put(&models.User{ID: user.ID, FirstName: "Test", LastName: "Update", Email: "javier@gmail.com"})
	as.Equal(404, resp404.Code)
	respVacio := as.HTML("/user/update/{%s}", user.ID).Put(&models.User{ID: user.ID, FirstName: "", LastName: "Update", Email: "javier@gmail.com"})
	as.Equal(303, respVacio.Code)
	count1, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count1)
}
func (as *ActionSuite) Test_Active_User() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)
	user := &models.User{
		FirstName: "Javier",
		LastName:  "Hernandez",
		Email:     "javier@gmail.com",
		Active:    false}
	verrs, err := as.DB.ValidateAndCreate(user)
	as.NoError(err)
	as.False(verrs.HasAny())
	respActive := as.HTML("/user/active/{%s}", user.ID).Put(&models.User{ID: user.ID, FirstName: "Javier", LastName: "Hernandez", Email: "javier@gmail.com", Active: false})
	as.Equal(303, respActive.Code)
	err = as.DB.Reload(user)
	as.NoError(err)
	as.Equal(true, user.Active)
	resp404 := as.HTML("/user/update/javier").Put(&models.User{ID: user.ID, FirstName: "Test", LastName: "Update", Email: "javier@gmail.com"})
	as.Equal(404, resp404.Code)
	count1, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count1)
}
