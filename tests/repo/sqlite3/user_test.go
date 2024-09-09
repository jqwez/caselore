package sqlite_test

import (
	"os"
	"testing"

	"github.com/jqwez/caselore/model"
	"github.com/jqwez/caselore/repo"
)

var userRepo, db = repo.NewUserRepository("mem_sqlite")

func TestMain(m *testing.M) {
	err := userRepo.CreateTable()
	if err != nil {
		panic(err)
	}

	code := m.Run()

	db.Close()
	os.Exit(code)
}

func TestUserCreate(t *testing.T) {
	t.Run("Empty user should not save", func(t *testing.T) {
		u := &model.User{}
		err := userRepo.Create(u)
		if err == nil {
			t.Fail()
		}
	})
	t.Run("User should be unique username", func(t *testing.T) {
		u := &model.User{
			Username: "testUniqueUser",
			Password: "random",
			Role:     "user",
		}
		if err := userRepo.Create(u); err != nil {
			t.Fatal("Failed to save username")
		}
		if err := userRepo.Create(u); err == nil {
			t.Fatal("Username saved twice")
		}
	})

}

func TestUserGetByID(t *testing.T) {
	u := &model.User{
		Username: "testUser",
		Password: "testPassword",
		Role:     "user",
	}
	err := userRepo.Create(u)
	if err != nil {
		t.Fatal(err)
	}
	fetchedUser, err := userRepo.GetByID(u.ID)
	if fetchedUser.ID != u.ID && fetchedUser.Username != u.Username {
		t.Fatal("Fetched User not the same")
	}
}

func TestUserUpdate(t *testing.T) {
	u := &model.User{
		Username: "testUpdateUser",
		Password: "testPassword",
		Role:     "user",
	}
	newUsername := "testUsernameUpdate"
	userRepo.Create(u)
	t.Run("Username update doesnt throw", func(*testing.T) {
		if err := userRepo.UpdateUsername(u, newUsername); err != nil {
			t.Fail()
		}
	})
	t.Run("Username is updated", func(*testing.T) {
		_, err := userRepo.GetByID(u.ID)
		if err != nil {
			t.Fatal(err)
		}
		user, _ := userRepo.GetByID(u.ID)
		if user.Username != newUsername {
			t.Log(user.Username)
			t.Log(newUsername)
			t.Fatal("Usernames do not match")
		}
	})
}
