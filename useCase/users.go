package user

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"go-server/database"
	"io/ioutil"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	// message := r.URL.Path
	// fmt.Println("message", message)
	// fmt.Println("r", r)
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("there is some error in sign up", err)
	}
	payload := make(map[string]string)
	error := json.Unmarshal(data, &payload)
	if error != nil {
		// panic(error)
		fmt.Println("error", error)
	}
	var message []byte
	user := database.GetUser(map[string]string{"email": payload["email"]})
	// fmt.Println("user1111", user)
	// fmt.Println("type user", reflect.TypeOf(user))
	if user.ID == "" {
		user, err := database.CreateUser(payload)
		if err != nil {
			fmt.Println("error", err)
			return
		}
		fmt.Println("user1212", user)
		user1 := database.GetUser(map[string]string{"email": payload["email"]})
		fmt.Println("user 1111", user1)
		message, _ = json.Marshal(user1)
	} else if user.Password == payload["Password"] {
		fmt.Println("Invalid password")
	} else {
		message, _ = json.Marshal(user)
	}
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
	h.Write(message)
	bs := h.Sum(nil)
	fmt.Println("bs %x", string(bs))
	w.Write(message)
}
