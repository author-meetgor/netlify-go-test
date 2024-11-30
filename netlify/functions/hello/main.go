package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tursodatabase/libsql-client-go/libsql"
)

type Blog struct {
	ID      int
	Title   string
	Content string
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	dbName := os.Getenv("TURSO_DATABASE_NAME")
	dbAuthToken := os.Getenv("TURSO_DATABASE_TOKEN")
	url := fmt.Sprintf("%s?authToken=%s", dbName, dbAuthToken)

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM blog")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var blogs []Blog

	for rows.Next() {
		var blog Blog

		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content); err != nil {
			fmt.Println("Error scanning row:", err)
		}

		blogs = append(blogs, blog)
		fmt.Println(blog.ID, blog.Title, blog.Content)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       blogs,
	}, nil
}

func main() {
	lambda.Start(handler)
}
