package user

import (
	"encoding/json"
	"fmt"
	"go-server/cryptography"
	"go-server/database"
	"io/ioutil"
	"net/http"

	middleware "go-server/middleware"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("there is some error in sign up", err)
	}
	payload := make(map[string]string)
	error := json.Unmarshal(data, &payload)
	if error != nil {
		fmt.Println("error", error)
	}
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"username": payload["email"],
	// 	"password": payload["password"],
	// })
	// fmt.Println("token", token)
	// tokenString, err := token.SignedString([]byte("secret"))
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// fmt.Println("token string", tokenString)
	token := middleware.CreateToken(payload["email"], payload["password"])
	fmt.Println("token11111111111111", token)
	middleware.ValidateToken(token)
	var message []byte
	user := database.GetUser(map[string]string{"email": payload["email"]})
	key, keyerr := cryptography.Encrypt(payload["email"])
	if keyerr != nil {
		// panic(keyerr)
		fmt.Println("key err", keyerr)
	}
	secret, secreterr := cryptography.Encrypt(payload["password"])
	if secreterr != nil {
		// panic(secreterr)
		fmt.Println("secret err", secreterr)
	}
	fmt.Println("key", key)
	fmt.Println("secret", secret)
	payload["key"] = key
	payload["secret"] = secret
	hashedPassword := middleware.CreateHash([]byte(payload["password"]))
	// result := middleware.ComparePassword([]byte(hashedPassword), []byte(payload["password"]))
	// fmt.Println("result", result)
	fmt.Println("hashed password", hashedPassword)
	payload["hashedPassword"] = hashedPassword
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
	w.Write(message)
}

func GetDataTypeWise(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	fmt.Println("data", data)
	userId := data["userId"][0]
	typeDataStr := ""
	value, ok := data["type"]
	fmt.Println("value", value)
	if ok && data["type"][0] != "" {
		typeDataStr = data["type"][0]
	}
	// payload := make(map[string]string)
	// err := json.Unmarshal(data, &payload)
	// if err != nil {
	// 	fmt.Println("There is some problem in getting data", err)
	// }
	project := make(map[string]int)
	// fmt.Println("project",)
	if typeDataStr != "" {
		project[typeDataStr] = 1
	}
	query := make(map[string]string)
	value, ok = data["userId"]
	fmt.Println("value", value)
	if ok && data["userId"][0] != "" {
		query["userid"] = userId
	}
	typeData := database.GetData(query, project)
	fmt.Println("data", typeData)
	message, _ := json.Marshal(typeData)
	// message := "check"
	w.Write([]byte(message))
}

func GetData(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	fmt.Println("data", data)
	// typeData := data["type"][0]
	query := make(map[string]string)
	value, ok := data["type"]
	// fmt.Println("value1111", value)
	// typeData := data["type"][0]
	if ok && data["type"][0] != "" {
		query["datatype"] = data["type"][0]
	}
	value, ok = data["ip"]
	if ok && data["ip"][0] != "" {
		query["ip"] = data["ip"][0]
	}
	value, ok = data["reqId"]
	if ok && data["reqId"][0] != "" {
		query["reqid"] = data["reqId"][0]
	}
	fmt.Println("query", query)
	fmt.Println("value", value)
	typeDataResult := database.GetUserData(query)
	fmt.Println("type data result", typeDataResult)
	// w.Write([]byte(typeDataResult))
	message, _ := json.Marshal(typeDataResult)
	w.Write([]byte(message))
}
