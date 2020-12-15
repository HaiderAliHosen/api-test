package apitest

//Getter --interface
type Getter interface {
	GetAll() []Item
}

//Adder -- interface
type Adder interface {
	Add(item Item)
}

// Item -- struct
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

// Repo -- struct
type Repo struct {
	Items []Item
}

// New -- repo
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

// Add -- list
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

//GetAll -- getall
func (r *Repo) GetAll() []Item {
	return r.Items
}
