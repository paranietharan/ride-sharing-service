package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

// Connect establishes a connection to the PostgreSQL database
func Connect() (*pgx.Conn, error) {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "root"
	dbname := "ride_service"

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)

	// Establish the connection
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	log.Println("Connected to PostgreSQL......")
	return conn, nil
}

func Disconnect(conn *pgx.Conn) {
	err := conn.Close(context.Background())
	if err != nil {
		log.Printf("Error closing connection: %v", err)
	} else {
		log.Println("Disconnected from PostgreSQL")
	}
}
