package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetRegisteredDriver(t *testing.T) {
	assert.Equal(t, []string{"mysql"}, sql.Drivers())
}

func TestPingMySql(t *testing.T) {
	EnvLoad()
	databasePassword := os.Getenv("MYSQL_ROOT_PASSWORD")

	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@(localhost:3306)/sample_db", databasePassword))

	assert.Nil(t, err)
	assert.NotNil(t, db)

	defer db.Close()

	err = db.Ping()
	assert.Nil(t, err)
}
