package service_models

type PhoneNumber struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Contact struct {
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
	Name         string        `json:"name"`
}
