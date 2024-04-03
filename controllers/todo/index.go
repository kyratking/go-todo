package todo_controllers

import (
	"encoding/json"
	"example/todo/config"
	"example/todo/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	id          int
	title       string
	description string
	created_at  string
	updated_at  string
	database    = config.Database()
)

func Show(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	// Validate and convert limit parameter to integer
	limit := "10"
	if limitStr != "" {
		limitInt, err := strconv.Atoi(limitStr)
		if err != nil || limitInt <= 0 {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
		limit = limitStr
	}

	// Validate and convert offset parameter to integer
	offset := "0"
	if offsetStr != "" {
		offsetInt, err := strconv.Atoi(offsetStr)
		if err != nil || offsetInt < 0 {
			http.Error(w, "Invalid offset parameter", http.StatusBadRequest)
			return
		}
		offset = offsetStr
	}

	if limitStr == "" {
		limit = "10"
	}

	if offsetStr == "" {
		limit = "0"
	}

	query := fmt.Sprintf("SELECT * FROM todos LIMIT %s OFFSET %s", limit, offset)

	log.Println(query)

	statement, err := database.Query(query)

	if err != nil {
		fmt.Println(err)
	}

	var todos []models.Todo

	for statement.Next() {
		err = statement.Scan(&id, &title, &description, &created_at, &updated_at)

		if err != nil {
			fmt.Println(err)
		}

		todo := models.Todo{
			Id:          id,
			Title:       title,
			Description: description,
			CreatedAt:   created_at,
			UpdatedAt:   updated_at,
		}

		todos = append(todos, todo)
	}

	json.NewEncoder(w).Encode(todos)
}

func Add(w http.ResponseWriter, r *http.Request) {}

func Delete(w http.ResponseWriter, r *http.Request) {}

func Edit(w http.ResponseWriter, r *http.Request) {}
