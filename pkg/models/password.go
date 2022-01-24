package models

import "errors"

type PasswordList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type PasswordItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	WebSite     string `json:"web_site" db:"web_site"`
	Login       string `json:"login" db:"login"`
	Phone       string `json:"phone" db:"phone"`
	Email       string `json:"email" db:"email"`
	Pass        string `json:"pass" db:"pass"`
	SecretWord  string `json:"secret_word" db:"secret_word"`
	BackupCodes string `json:"backup_codes" db:"backup_codes"`
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

type UpdatePasswordInput struct {
	Title       *string `json:"title"`
	WebSite     *string `json:"web_site"`
	Login       *string `json:"login"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
	Pass        *string `json:"pass"`
	SecretWord  *string `json:"secret_word"`
	BackupCodes *string `json:"backup_codes"`
}

func (i UpdatePasswordInput) Validate() error {
	if i.Title == nil && i.WebSite == nil && i.Login == nil && i.Phone == nil &&
		i.Email == nil && i.Pass == nil && i.SecretWord == nil && i.BackupCodes == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
