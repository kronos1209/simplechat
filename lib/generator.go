package lib

import "math/rand"

const runes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int) string {
	var res string
	for i := 0; i < length; i++ {
		r := runes[rand.Intn(len(runes)-1)]
		res += string(r)
	}
	return res
}
