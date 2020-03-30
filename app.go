package go_mux_tutorial

import (
	"database/sql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB *sql.DB
}
