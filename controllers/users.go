package controllers

import (
	"fmt"
	"net/http"

	"github.com/sderohan/photo-gallery/models"
	"github.com/sderohan/photo-gallery/views"
)

func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView:   views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:        us,
	}
}

type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        *models.UserService
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type LoginForm struct {
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

// Login is used to verify the provided email address and
// Password and then log the user in if they are correct.
//
// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	user, err := u.us.Authenticate(form.Email, form.Password)

	switch err {
	case models.ErrorNotFound:
		fmt.Fprintln(w, "Invalid email address.")
	case models.ErrInvalidPassword:
		fmt.Fprintln(w, "Invalid password provided.")
	case nil:
		fmt.Fprintln(w, user)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
