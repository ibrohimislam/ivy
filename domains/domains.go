package domains

type DataRepository interface {
	Store(entity Entity) error
	FindById(id string) *Entity
}

type UserRepository interface {
	Store(user User) error
	FindById(id string) *User
}

type Data string

type MetaData struct {
	Label string `json:"label"`
	Type  string `json:"type"`
}

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"isAdmin"`
}

type Entity struct {
	Id       string
	DataSet  []Data     `json:"dataSet"`
	MetaData []MetaData `json:"metaData"`
	OwnerId  string     `json:"user"`
}
