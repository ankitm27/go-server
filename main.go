package main

import (
	"crypto/sha1"
	"fmt"
	database "go-server/database"
	redisclient "go-server/redis"
	socketserver "go-server/socket"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/fvbock/endless"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	// message := r.URL.Path
	// fmt.Println("message", message)
	// fmt.Println("r", r)
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("there is some error in sign up", err)
	}
	fmt.Println("data", string(data))
	message := "check"
	// uuidData, err := uuid.NewV4()
	// if err != nil {
	// fmt.Println("error in creating uuid", err)
	// }
	// fmt.Println("uuid", uuidData)
	// fmt.Printf("UUIDv4: %s\n", uuidData)
	// fmt.Println("type of", reflect.TypeOf(uuidData))
	// var buf [36]byte
	// encodeHex(buf[:], uuid)
	// uuidStr := buf[:]
	// uuisStr := uuid.String(uuidData)
	h := sha1.New()
	h.Write([]byte(message))
	bs := h.Sum(nil)
	fmt.Println("bs %x", string(bs))
	w.Write([]byte(message))
}

func authenticateData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

func main() {
	redisclient.RedisClient("localhost")
	database.DatabaseConnect()
	redisclient.Schedule(1 * time.Second)
	go socketserver.CreateServer(3333)
	// fmt.Println("")
	signUpFunctionCall := http.HandlerFunc(signUp)
	mux := mux.NewRouter()
	mux.Handle("/signup", authenticateData(signUpFunctionCall))

	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	fmt.Println("error in http server", err)
	// }
	err := endless.ListenAndServe("localhost:8080", mux)
	if err != nil {
		fmt.Println("there are some error in starting the golang server", err)
	}
}

// func encodeHex(dst []byte, uuid UUID) {
// 	hex.Encode(dst, uuid[:4])
// 	dst[8] = '-'
// 	hex.Encode(dst[9:13], uuid[4:6])
// 	dst[13] = '-'
// 	hex.Encode(dst[14:18], uuid[6:8])
// 	dst[18] = '-'
// 	hex.Encode(dst[19:23], uuid[8:10])
// 	dst[23] = '-'
// 	hex.Encode(dst[24:], uuid[10:])
// }
