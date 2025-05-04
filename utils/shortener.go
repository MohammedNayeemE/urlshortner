package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateUrl() string  {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	ret := make([]byte ,6);

	for i:=range ret {
		ret[i] = charset[r.Intn(len(charset))];
	}
	return string(ret);
}
