package app

type Item struct {
	ID    int
	Title string
	Done  bool
}

type ToDoMap map[int]Item

func (m ToDoMap) AddItem(title string) {
	id := len(m) + 1
	m[id] = Item{ID: id, Title: title, Done: false}
}

func (m ToDoMap) GetAllItems() []Item {
	items := make([]Item, 0, len(m))
	for _, item := range m {
		items = append(items, item)
	}
	return items
}
