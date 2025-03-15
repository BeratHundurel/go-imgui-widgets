package types

type TodoLists struct {
	Id    int `gorm:"primary_key"`
	Items []TodoItem `gorm:"foreignkey:TodoListID"`
	Title string
}

type TodoItem struct {
	Id        int `gorm:"primary_key"`
	Text      string
	Completed bool
	TodoListID int `gorm:"column:todo_list_id"`
}
