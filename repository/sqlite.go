package repository

import (
  "database/sql"
	_ "github.com/mattn/go-sqlite3"
  "github.com/kakoprog/sandbox_go_api/config"
)

func ConnectSqlite() *sql.DB {
  var dbpath string = config.Read("sqlite", "file_path").(string)

	connection, err := sql.Open("sqlite3", dbpath)
	if err != nil {
    panic(err)
  }

  return connection
}
