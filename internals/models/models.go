package models

type User struct {
	GID string `json:"gid"`
	Email string `json:"email"`
	Name string `json:"name"`

}

type UserResponse struct {
	Data User `json:"data"`
}

type UsersResponse struct {
	Data []User `json:"data"`
}