package utils

import (
	"golang.org/x/crypto/argon2"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func GenerateToken(Snowflake int64) string {
	rand.Seed(time.Now().UnixNano())

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	salt := make([]byte, 12)
	for i := range salt {
		salt[i] = letters[rand.Intn(len(letters))]
	}

	snowflake := int64ToBase32(Snowflake) 
	timestamp := getTimestamp()

	token := snowflake + "." + timestamp + "." + string(salt)

	return token
}

func getTimestamp() string {
	currentTime := time.Now()

	year := currentTime.Year()
	month := int(currentTime.Month())
	day := currentTime.Day()
	hour := currentTime.Hour()
	minute := currentTime.Minute()
	second := currentTime.Second()

	data := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
	encoded := base32.StdEncoding.EncodeToString([]byte(data))

	return encoded
}

 func HashPassword(password, email string) string {
    salt := []byte(email)
    hashedPassword := argon2.IDKey([]byte(password), salt, 2, 64*1024, 4, 32)

    encodedHash := base64.URLEncoding.EncodeToString(hashedPassword)
    return encodedHash
}

func int64ToBase32(value int64) string {
	byteValue := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		byteValue[i] = byte(value & 0xFF)
		value >>= 8
	}

	encoding := base32.StdEncoding

	base32String := encoding.EncodeToString(byteValue)

	return base32String
}