package password

import (
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func Hash(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(applyPepper(raw)), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Verify(hashed string, raw string) error {
	peppered := applyPepper(raw)
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(peppered)); err == nil {
		return nil
	}

	// Compatibilidad retroactiva: hashes antiguos sin pepper.
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw))
}

func applyPepper(raw string) string {
	pepper := strings.TrimSpace(os.Getenv("PASSWORD_PEPPER"))
	if pepper == "" {
		pepper = strings.TrimSpace(os.Getenv("JWT_SECRET"))
	}
	if pepper == "" {
		return raw
	}
	return raw + "::" + pepper
}
