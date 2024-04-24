package db

import (
	"aiexplorer/configs"
	"database/sql"
	"fmt"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)
	conn, err := sql.Open("postgres", sc)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
