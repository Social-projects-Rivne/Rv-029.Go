package model

import (
	"crypto/rand"
	"io"

	"github.com/gocql/gocql"
	"golang.org/x/crypto/scrypt"
)

//User type
type User struct {
	UUID      gocql.UUID
	Email     string
	FirstName string
	LastName  string
	Password  []byte
	Salt      []byte
	Role      string
	CreatedAt 
	UpdatedAt 
}

const (
	PW_SALT_BYTES = 32
	PW_HASH_BYTES = 64
)

//generateSalt is generating random salt
func (u *User) generateSalt() error {
	u.Salt = make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, u.Salt)
	if err != nil {
		return err
	}
	return nil
}

//EncryptPassword is encrypt your password by salt
func (u *User) EncryptPassword() error {

	u.generateSalt()

	pas, err := scrypt.Key(u.Password, u.Salt, 1<<14, 8, 1, PW_HASH_BYTES)
	u.Password = pas
	if err != nil {
		return err
	}

	return nil

}


