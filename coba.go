package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	urlExample := "postgres://cahyono:csteam512@@localhost:5432/aksi"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
  fmt.Println("coonnnnect")
	defer conn.Close(context.Background())

}
