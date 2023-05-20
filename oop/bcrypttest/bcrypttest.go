package bcrypttest

import "golang.org/x/crypto/bcrypt"

type BCrypt struct {
	pepper []byte
}

func (b BCrypt) Equal(hashedPw, rawPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPw),
		append([]byte(rawPw), b.pepper...))
	if err == nil {
		return true
	}

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}

	// If this happens it means our password hash isn't valid
	// or some other unrecoverable error happened.
	panic(err)

}

func (b BCrypt) hash(rawPw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		append([]byte(rawPw), b.pepper...),
		bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func NewBCrypt(pepper string) BCrypt {
	return BCrypt{pepper: []byte(pepper)}
}
