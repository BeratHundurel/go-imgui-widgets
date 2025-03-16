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
	GetAllTodos() []types.TodoLists
	GetTodoList(id int) types.TodoLists
	CreateTodoList(title string)
	CreateTodoItem(todoListID int, text string)
	DeleteTodoItem(todoItemID int)
	CompleteTodoItem(todoItemID int)
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

func  (s *service )GetAllTodos() []types.TodoLists {
	var todos []types.TodoLists
	s.db.Preload("Items").Find(&todos)
	return todos
}

func (s *service) GetTodoList(id int) types.TodoLists {
	var todo types.TodoLists
	s.db.Preload("Items").First(&todo, id)
	return todo
}

func (s *service) CreateTodoList(title string) {
	s.db.Create(&types.TodoLists{Title: title})
}

func (s *service) CreateTodoItem(todoListID int, text string) {
	s.db.Create(&types.TodoItem{TodoListID: todoListID, Text: text, Completed: false})
}

func (s *service) DeleteTodoItem(todoItemID int) {
	s.db.Delete(&types.TodoItem{}, todoItemID)
}

func (s *service) CompleteTodoItem(todoItemID int) {
	s.db.Model(&types.TodoItem{}).Where("id = ?", todoItemID).Update("completed", true)
}

func (s *service) Migrate() {
	s.db.AutoMigrate(&types.TodoLists{}, &types.TodoItem{})
}
