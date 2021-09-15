package controllers

import (
	"fmt"
	"net/http"

	"github.com/sderohan/photo-gallery/models"
	"github.com/sderohan/photo-gallery/views"
)

func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		us:      us,
	}
}

type Users struct {
	NewView *views.View
	us      *models.UserService
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// This is used to render the form where a user can create a new user account.
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

// This is used to process the signup form when a user tries to create a new user account.
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm

	// fetch the data
	// destination for Decode function should be the pointer
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	// fmt.Fprintln(w, r.PostForm["email"])
	// fmt.Fprintln(w, r.PostForm["password"])
	user := &models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	err := u.us.Create(user)
	fmt.Fprintln(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
