package services

import (
	"Accounting/config"
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

//LoginService for initial all function
type LoginService interface {
	HandleAuthenticator(string) bool
	GenerateAccessToken() (string, error)
}

type loginService struct {
	accessToken string
}

//NewLoginService construct to access function interface
func NewLoginService() *loginService {
	return &loginService{
		accessToken: config.Getenv("ACCESS_TOKEN"),
	}
}

func (s *loginService) HandleAuthenticator(accessToken string) bool {
	if accessToken == "" {
		return false
	}
	const authType = "Bearer "
	accessToken = accessToken[len(authType):]
	return CheckMatchHashing(s.accessToken, accessToken)
}

func (s *loginService) GenerateAccessToken() (string, error) {
	return GenerateHashing(s.accessToken, config.TOKENCOST)
}

//GenerateHashing be used for hanshing Password
func GenerateHashing(str string, n int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(EncryptionWithMD5(str)), n)
	return string(bytes), err
}

//EncryptionWithMD5 be used for encryption document with AES Chipper
func EncryptionWithMD5(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

//CheckMatchHashing be used to validate password hashing
func CheckMatchHashing(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(EncryptionWithMD5(password)))
	return err == nil
}
