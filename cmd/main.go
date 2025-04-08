package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-app/infrastructure/mysql"
	"todo-app/interface/handler"
	"todo-app/usecase"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Print("Failed to connect to database")
	}

	if err := db.Ping(); err != nil {
		log.Print("Database ping failed")
	}

	userRepo := mysql.NewUserRepository(db)
	todoRepo := mysql.NewTodoRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)

	userHandler := handler.NewUserHandler(userUsecase)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			userHandler.GetUser(w, r)
		case "POST":
			userHandler.CreateUser(w, r)
		case "PUT":
			userHandler.UpdateUser(w, r)
		case "DELETE":
			userHandler.DeleteUser(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			todoHandler.GetTodo(w, r)
		case "POST":
			todoHandler.CreateTodo(w, r)
		case "PUT":
			todoHandler.UpdateTodo(w, r)
		case "DELETE":
			todoHandler.DeleteTodo(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", nil)
}
