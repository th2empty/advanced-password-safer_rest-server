package models

type PasswordList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type PasswordItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	WebSite     string `json:"web_site"`
	Login       string `json:"login"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Pass        string `json:"pass"`
	SecretWord  string `json:"secret_word"`
	BackupCodes string `json:"backup_codes"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type ListItems struct {
	Id     int
	ListId int
	ItemId int
}
