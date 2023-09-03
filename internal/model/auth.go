package model

type RegClientDTO struct {
	Login      string `json:"login"`
	UniqueName string `json:"unique_name"`
	Nickname   string `json:"nickname"`
	Hash       string `json:"hash"`
}

type AuthClientDTO struct {
	Login string `json:"login"`
	Hash  string `json:"hash"`
}
