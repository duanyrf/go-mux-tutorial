package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func (app *App) Initialize(user, password, dbname string) {}

func (app *App) Run(addr string) {}

type App struct {
	Router *mux.Router
	DB *sql.DB
}
