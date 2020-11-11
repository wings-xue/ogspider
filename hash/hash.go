package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func Hash(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])

}
