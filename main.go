package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/libsql/go-libsql"

	"github.com/MaximilianFullStack/htmx-blog/components"
	"github.com/MaximilianFullStack/htmx-blog/structs"
)

func queryArticles(db *sql.DB) []structs.Article {
	rows, err := db.Query("SELECT * FROM articles")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var articles []structs.Article

	for rows.Next() {
		var article structs.Article

		if err := rows.Scan(&article.ID, &article.Title, &article.Description, &article.Tags, &article.Content); err != nil {
			fmt.Println("Error scanning row:", err)
			log.Panic("arg")
		}

		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	return articles
}

// Creates a Turso Embedded DB from a remote DB
func accessEmbeddedDb() (dir string, conn *libsql.Connector, db *sql.DB) {
	dbName := "./local.db"
	primaryUrl := os.Getenv("DB_URL")
	if primaryUrl == "" {
		fmt.Println("DB URL Empty")
		os.Exit(1)
	}

	authToken := os.Getenv("DB_ACCESS_TOKEN")

	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		fmt.Println("Error creating temporary directory:", err)
		os.Exit(1)
	}

	dbPath := filepath.Join(dir, dbName)
	// duration := time.Minute

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl, authToken /** duration */)
	if err != nil {
		fmt.Println("Error creating connector:", err)
		os.Exit(1)
	}

	db = sql.OpenDB(connector)
	return dir, connector, db
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dir, conn, db := accessEmbeddedDb()
	defer os.RemoveAll(dir)
	defer conn.Close()
	defer db.Close()

	articles := queryArticles(db)

	component := components.Home(articles)

	http.Handle("/", templ.Handler(component))

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Server running on [::]:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
