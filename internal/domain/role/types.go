package role

/*
type RoleTitle string

const (
	User  RoleTitle = "users"
	Admin RoleTitle = "admin"
)
*/

type Role struct {
	Title string
	ID    int
}

type Input struct {
	RoleTitle string `json:"role" binding:"required"`
}
