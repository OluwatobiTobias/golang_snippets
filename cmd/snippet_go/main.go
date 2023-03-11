package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("could not load .env by godotenv")
	}
	db_Conn, ok := os.LookupEnv("db_login")
	if !ok {
		fmt.Println("could not load .env by os package")
	}

	conn, err := pgx.Connect(context.Background(), db_Conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect\n %v\n", err)
	}

	defer conn.Close(context.Background())

	if row, err := conn.Query(context.Background(), "select content from public.posts"); err != nil {
		fmt.Println("query-level-error", err)

	} else {

		var man string

		for row.Next() {

			row.Scan(&man)
			fmt.Println(man)

		}
		if row.Err() != nil {
			fmt.Println("error is", err)
		}
	}
}
