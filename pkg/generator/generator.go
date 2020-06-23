package generator

import (
	"context"
	"golang-grpc-password-generator/pkg/api"
	"math/rand"
	"strings"
	"time"
)

type Generator struct {
}

func (*Generator) Generator(ctx context.Context, req *api.GeneratorRequest) (*api.GeneratorResponse, error) {
	length := req.Length
	num := req.Num
	up := req.Up

	password := generatePassword(length, num, up)

	response := &api.GeneratorResponse{Password: password}

	return response, nil
}

func generatePassword(length int64, num bool, up bool) string {
	char := "qwertyuiopasdfghjklzxcvbnm"
	var password string

	if up {
		char += strings.ToUpper(char)
	}

	if num {
		char += "0123456789"
	}

	charLength := int64(len(char))

	var i int64 = 0
	for ; i < length; i++ {
		rand.Seed(time.Now().UTC().UnixNano())

		random := rand.Int63n(charLength)
		randomChar := string(char[random])

		password += randomChar
	}

	return password
}
