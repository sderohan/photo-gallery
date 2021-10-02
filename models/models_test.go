package models

import (
	"fmt"
	"testing"
	"time"
)

func testUserService() (UserService, error) {
	const (
		host     = "localhost"
		port     = 49154
		user     = "root"
		password = "root"
		dbname   = "photo_gallery_test"
	)
	mysqlInfo := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", user, password, host, port, dbname)

	us, err := NewUserService(mysqlInfo)
	if err != nil {
		panic(err)
	}
	// Clear the users table between tests
	us.DestructiveReset()
	return us, nil
}

func TestCreateUser(t *testing.T) {
	us, err := testUserService()
	if err != nil {
		t.Fatal(err)
	}
	user := User{
		Name:  "jack",
		Email: "jack@gmail.com",
	}
	err = us.Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	if user.ID == 0 {
		t.Errorf("Expected ID > 0 %d", user.ID)
	}
	if time.Since(user.CreatedAt) > time.Duration(5*time.Second) {
		t.Errorf("Expected CreatedAt to be recent. Received %s", user.CreatedAt)
	}
	if time.Since(user.UpdatedAt) > time.Duration(5*time.Second) {
		t.Errorf("Expected UpdatedAt to be recent. Received %s", user.UpdatedAt)
	}
}
