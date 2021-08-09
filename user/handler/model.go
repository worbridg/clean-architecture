package handler

type Model struct {
	User User `json:"user"`
}

type User struct {
	Name string `json:"name"`
}
