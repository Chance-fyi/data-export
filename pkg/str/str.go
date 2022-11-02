package str

import (
	"github.com/iancoleman/strcase"
	"math/rand"
)

// Snake => chance_fyi
func Snake(s string) string {
	return strcase.ToSnake(s)
}

// Camel => ChanceFyi
func Camel(s string) string {
	return strcase.ToCamel(s)
}

// LowerCamel => chanceFyi
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
