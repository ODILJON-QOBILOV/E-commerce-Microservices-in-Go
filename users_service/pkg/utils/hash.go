package utils

import "golang.org/x/crypto/bcrypt"

func Password(pw string) string {
    hash, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
    return string(hash)
}

func CheckPassword(pw, hash string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw)) == nil
}
