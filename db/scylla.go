package db

import (
	"fmt"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func Init() {
	cluster := gocql.NewCluster("127.0.0.1:9042") // Update with your ScyllaDB cluster IP
	cluster.Keyspace = "system"               // Using the system keyspace for creating the keyspace and table
	tempSession, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	// Create keyspace if not exists
	err = tempSession.Query(`
		CREATE KEYSPACE IF NOT EXISTS todos_keyspace 
		WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`).Exec()
	if err != nil {
		panic(err)
	}

	// Switch to the newly created keyspace
	cluster.Keyspace = "todos_keyspace"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	

	// Create table if not exists
	err = Session.Query(`
		CREATE TABLE IF NOT EXISTS todos (
			id TEXT,
			user_id INT,
			title TEXT,
			description TEXT,
			status TEXT,
			created TIMESTAMP,
			updated TIMESTAMP,
			PRIMARY KEY (user_id, id)
		)`).Exec()
	if err != nil {
		panic(err)
	}

	fmt.Println("Initialized ScyllaDB keyspace and table")
}
