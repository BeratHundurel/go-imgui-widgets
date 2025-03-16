package types

type TodoLists struct {
	Id    int        `gorm:"primary_key"`
	Title string     `gorm:"not null"`
	Items []TodoItem `gorm:"foreignkey:TodoListID"`
}

type TodoItem struct {
	Id         int    `gorm:"primary_key"`
	Text       string `gorm:"not null"`
	Completed  bool
	TodoListID int `gorm:"column:todo_list_id"`
}
