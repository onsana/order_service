package dto

type UserDto struct {
	ID          int      `json:"id"`
	Name        string   `json:"user_name"`
	PhoneNumber string   `json:"phone_number"`
	Roles       []string `json:"roles"`
	Blocked     bool     `json:"is_blocked"`
}
