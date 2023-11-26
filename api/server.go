package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"server/api/config"
	"server/api/handlers"
	sqlm "server/api/models/sql"

	_ "github.com/go-sql-driver/mysql"
)

type handlersMap map[string]map[string]http.HandlerFunc

func (hMap handlersMap) addHandlerToMap(pattern, method string, handler func(http.ResponseWriter, *http.Request)) {
	if _, ok := hMap[pattern]; !ok {
		hMap[pattern] = make(map[string]http.HandlerFunc)
	}
	hMap[pattern][method] = handler
}

func (hMap handlersMap) initRoutes(app *config.Application) {
	hMap.addHandlerToMap("/users", "GET", handlers.UsersList(app))
	hMap.addHandlerToMap("/users", "POST", handlers.CreateUser(app))
	hMap.addHandlerToMap("/user", "GET", handlers.User(app))
}

func (hMap handlersMap) mux(app *config.Application) *http.ServeMux {
	hMap.initRoutes(app)

	mux := http.NewServeMux()
	for url := range hMap {
		func(pattern string) {
			mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
				if handler, ok := hMap[pattern][r.Method]; ok {
					handler(w, r)
					return
				}
				http.NotFound(w, r)
			})
		}(url)
	}
	return mux
}

func RunServer(address string) error {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/backend?parseTime=true")
	if err != nil {
		return err
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		return err
	}
	app := &config.Application{
		InfoLog: log.New(os.Stdout, "INFO", log.Ldate|log.Ltime),
		ErrLog:  log.New(os.Stderr, "ERROR", log.Ldate|log.Ltime|log.Lshortfile),
		Users:   &sqlm.UserModel{Db: db},
	}
	server := &http.Server{
		Addr:     address,
		ErrorLog: app.ErrLog,
		Handler:  handlersMap{}.mux(app),
	}
	return server.ListenAndServe()
}
