package pwd

import (
	"golang.org/x/crypto/bcrypt"
)

// Bcrypt contain cost
type Bcrypt struct {
	cost int
}

// NewBcrypt is a constructor for init
func NewBcrypt(cost int) *Bcrypt {
	return &Bcrypt{cost: cost}
}

// Hash is method for Bcypt to hash password return to string
func (b *Bcrypt) Hash(password string) (string, error) {
	// bcrypt easy steady but slow
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// Compare is method for Bcypt to compare between password and hash
func (b *Bcrypt) Compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
