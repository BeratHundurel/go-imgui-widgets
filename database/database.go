package database

import (
	"imgui_try/types"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {
	Migrate()
}

type service struct {
	db *gorm.DB
}

var (
	dburl      = "imgui_widgets.db"
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	db, err := gorm.Open(sqlite.Open(dburl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func GetAllTodos() []types.TodoLists {
	db := dbInstance.db
	var todos []types.TodoLists
	db.Find(&todos)
	return todos
}

func GetTodoList(id int) types.TodoLists {
	db := dbInstance.db
	var todo types.TodoLists
	db.Preload("Items").First(&todo, id)
	return todo
}

func CreateTodoList(title string) {
	db := dbInstance.db
	db.Create(&types.TodoLists{Title: title})
}

func CreateTodoItem(todoListID int, text string) {
	db := dbInstance.db
	db.Create(&types.TodoItem{TodoListID: todoListID, Text: text})
}

func (s *service) Migrate() {
	s.db.AutoMigrate(&types.TodoLists{}, &types.TodoItem{})
}
