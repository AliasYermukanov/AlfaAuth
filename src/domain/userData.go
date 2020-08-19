package domain

type User struct {
	Uid               string   `json:"uid"`
	MobilePhone       string   `json:"mobile_phone"`
	Password          string   `json:"password,omitempty"`
	EncryptedPassword []byte   `json:"encrypted_password,omitempty"`
	Scope             map[string]string `json:"scope"`
}

type UsersScope struct {
	Scope             map[string]string `json:"scope"`
}