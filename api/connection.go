package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"time"
)

func (app *application) setupDB() *sql.DB {
	dsn := app.config.dsn
	conn, err := sql.Open("sqlite", dsn)
	conn.SetMaxOpenConns(5)
	conn.SetMaxIdleConns(5)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = conn.PingContext(ctx)
	if err != nil {
		fmt.Println("Unable to connect to DB", err)
		panic(err)
	}
	return conn
}

func setupUsersTable(conn *sql.DB) {
	_, err := conn.Exec("CREATE TABLE IF NOT EXISTS users (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT," +
		"first_name TEXT NOT NULL," +
		"last_name TEXT NOT NULL," +
		"email TEXT NOT NULL UNIQUE," +
		"password TEXT NOT NULL" +
		");")
	if err != nil {
		return
	}
}
