package types

type PhoneBook struct {
	ID          string `json:"id"`
	PhoneNumber int64  `json:"phone_number"`
	UserName    string `json:"userName"`
}
