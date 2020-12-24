package models

import (
	"foodapi/db"
	"os"
)

// Mongo server ip -> localhost -> 127.0.0.1 -> 0.0.0.0
var server = os.Getenv("DATABASE")
// Database name
var databaseName = os.Getenv("DATABASE_NAME")
// Create a connection
var dbConnect = db.NewConnection(server, databaseName)