package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/api/config"
	"server/api/models"
	"strconv"
)

type Parent struct {
	Wife string
	Name string
	Age  int
}

func UsersList(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := app.Users.Latest()
		if err != nil {
			app.ErrLog.Print(err)
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		result, err := json.Marshal(users)
		if err != nil {
			app.ErrLog.Print(err)
			http.NotFound(w, r)
			return
		}
		w.Write(result)
	}
}

func CreateUser(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			app.ErrLog.Print(err)
			http.NotFound(w, r)
			return
		}
		r.Body.Close()

		var user models.User
		if err = json.Unmarshal(body, &user); err != nil {
			app.ErrLog.Print(err)
			http.NotFound(w, r)
			return
		}

		id, err := app.Users.Insert(user.Name, user.Surname, user.Age)
		if err != nil {
			app.ErrLog.Print(err)
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/user?id=%d", id), http.StatusFound)
	}
}

func User(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			app.ErrLog.Print(err)
			http.NotFound(w, r)
			return
		}

		user, err := app.Users.Get(id)
		if err != nil {
			app.ErrLog.Print(err)
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		result, err := json.Marshal(user)
		if err != nil {
			app.ErrLog.Print(err)
			http.NotFound(w, r)
			return
		}
		w.Write(result)
	}
}
