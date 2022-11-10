package storage

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Comment struct {
	ID          int32    `json:"id"`
	Author  	string `json:"author"`
	ProfileImg 	string `json:"profile_img"`
	Date 		string `json:"date"`
	Comment 	string `json:"comment"`
	CommentImg 	string `json:"comment_img"`
}

type User struct {
	ID                int32    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
	ProfileName       string `json:"profile_name"`
	ProfileImg        string `json:"profile_img"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}

		return nil
	}

}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc

	}

	return nil
}
