package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

var my_secert = "xiaofanyi"

type H map[string]interface{}

func EncodeJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

// use md5 to Encrypt Password
func EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(my_secert))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

// validator error
func ValidatorError(err error) string {
	return "输入内容不合法,请重新输入"
}

//	func GetIdentity() int64 {
//		var st time.Time
//		st, err := time.Parse("2006-01-02", "2023-02-24")
//		if err != nil {
//			log.Fatal(err)
//		}
//		snowflake.Epoch = st.UnixNano() / 10000
//		node, err := snowflake.NewNode(1)
//		if err != nil {
//			log.Fatal(err)
//		}
//		return node.Generate().Int64()
//	}

func GenUserNumber() int64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Int63n(1000000000)
}
