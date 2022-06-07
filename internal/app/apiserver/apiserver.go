package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/RomanDovgii/go-restapi/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost user=rdovgii password=123456 dbname=restapi_dev sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
