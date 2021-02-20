package Tools

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


//计算md5
func Md5(source string) string{
	md5 := md5.New()
	md5.Write([]byte(source))
	return hex.EncodeToString(md5.Sum(nil))
}