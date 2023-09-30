package util

import (
	"crypto/md5"
	"encoding/base64"
	"time"
)

var ProcessID string

func NewID() string {
	hashInput := time.Now().String()
	hash := md5.Sum([]byte(hashInput))
	id := base64.StdEncoding.EncodeToString(hash[:])
	return id
}
