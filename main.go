package main

import (
	"fmt"
	database "go-server/database"
	"net/http"

	"github.com/gorilla/mux"

	user "go-server/useCase"

	middleware "go-server/middleware"

	"github.com/fvbock/endless"
)

// func signUp(w http.ResponseWriter, r *http.Request) {
// 	// message := r.URL.Path
// 	// fmt.Println("message", message)
// 	// fmt.Println("r", r)
// 	data, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Println("there is some error in sign up", err)
// 	}
// 	payload := make(map[string]string)
// 	error := json.Unmarshal(data, &payload)
// 	if error != nil {
// 		// panic(error)
// 		fmt.Println("error", error)
// 	}
// 	var message []byte
// 	user := database.GetUser(map[string]string{"email": payload["email"]})
// 	// fmt.Println("user1111", user)
// 	// fmt.Println("type user", reflect.TypeOf(user))
// 	if user.ID == "" {
// 		user, err := database.CreateUser(payload)
// 		if err != nil {
// 			fmt.Println("error", err)
// 			return
// 		}
// 		fmt.Println("user1212", user)
// 		user1 := database.GetUser(map[string]string{"email": payload["email"]})
// 		fmt.Println("user 1111", user1)
// 		message, _ = json.Marshal(user1)
// 	} else if user.Password == payload["Password"] {
// 		fmt.Println("Invalid password")
// 	} else {
// 		message, _ = json.Marshal(user)
// 	}
// 	// uuidData, err := uuid.NewV4()
// 	// if err != nil {
// 	// fmt.Println("error in creating uuid", err)
// 	// }
// 	// fmt.Println("uuid", uuidData)
// 	// fmt.Printf("UUIDv4: %s\n", uuidData)
// 	// fmt.Println("type of", reflect.TypeOf(uuidData))
// 	// var buf [36]byte
// 	// encodeHex(buf[:], uuid)
// 	// uuidStr := buf[:]
// 	// uuisStr := uuid.String(uuidData)
// 	h := sha1.New()
// 	h.Write(message)
// 	bs := h.Sum(nil)
// 	fmt.Println("bs %x", string(bs))
// 	w.Write(message)
// }

// func authenticateData(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("Executing middlewareOne")
// 		next.ServeHTTP(w, r)
// 		log.Println("Executing middlewareOne again")
// 	})
// }

func main() {
	// redisclient.RedisClient("localhost")
	database.DatabaseConnect()
	// redisclient.Schedule(1 * time.Second)
	// go socketserver.CreateServer(3333)
	// fmt.Println("")
	// signUpFunctionCall := http.HandlerFunc(user.SignUp)
	mux := mux.NewRouter()
	// mux.Handle("/signup", authenticateData(signUpFunctionCall))
	mux.Handle("/signup", middleware.AuthenticateData(http.HandlerFunc(user.SignUp)))
	mux.Handle("/getdatatypewise", middleware.AuthenticateData(http.HandlerFunc(user.GetDataTypeWise)))
	mux.Handle("/getdata", middleware.AuthenticateData(http.HandlerFunc(user.GetData)))
	// http.Handle("/", authenticateData(mux))
	// delivery.Check()
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	fmt.Println("error in http server", err)
	// }
	// a := delivery.UserDelivery()
	// fmt.Println("a", a)
	// mux.Handle("/", delivery.UserDelivery())
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
