package alice_blue_api

import (
	"fmt"
	_"io/ioutil"
	_"net/http"
	"strings"
	"errors"
	"crypto/sha256"
)

// struct to store auth data
type Auth struct {
	User_id string `json:"user_id"`
	Api_key string `json:"api_key"`
	Encryption_key string `json:"encryption_key"`
	Disable_ssl bool `json:"disable_ssl"`
	Session_id string `json:"session_id"`
	Base string `json:"base"`
	Api_name string `json:"api_name"`
	Version string `json:"version"`
	Bearear_token string `json:"bearear_token"`
}

// generate bearear token
func (obj *Auth) GenerateBearearToken(){
	obj.Bearear_token = obj.User_id + " " + obj.Session_id
}

// string to sha256 string
func toSHA256(s string) (string){
	sum := sha256.Sum256([]byte(s))
	res := fmt.Sprintf("%x",sum)
	return string(res)
}

// get encryption key
func (obj *Auth) GetEncryptionKey() (error) {
	url := `https://a3.aliceblueonline.com/rest/AliceBlueAPIService/api/customer/getAPIEncpkey`

	payload := fmt.Sprintf(`{ "userId":"%s"}`,obj.User_id)

	fmt.Println(payload)
	data := strings.NewReader(payload)
	
	// make request
	response , err := PostRequest(url, obj, data)

	if err != nil {
		fmt.Println(err)
		return err
	}
	if(response["stat"]=="Ok"){
		obj.Encryption_key = response["encKey"].(string)
		return nil
	}
	
	return errors.New("Something Happen Bad")
}


func (obj *Auth) GetSessionId() (error) {
	url := `https://a3.aliceblueonline.com/rest/AliceBlueAPIService/api/customer/getUserSID`

	key := toSHA256(fmt.Sprintf("%s%s%s",obj.User_id,obj.Api_key,obj.Encryption_key))

	payload := fmt.Sprintf(`{ "userId":"%s", "userData":"%s"}`,obj.User_id,key)

	fmt.Println(payload)
	data := strings.NewReader(payload)
	
	// make request
	response , err := PostRequest(url, obj, data)

	if err != nil {
		fmt.Println(err)
		return err
	}
	if(response["stat"]=="Ok"){
		obj.Encryption_key = response["sessionID"].(string)
		return nil
	}
	obj.Bearear_token=""
	
	return errors.New("Something Happen Bad in session id function")
}












func (obj *Auth) GetUserAgent() string {
	return obj.Api_name+" "+obj.Version
	// return "Animesh Shrivatri"+"1.0.25"
}

func (obj *Auth ) GetUserAuthorization() string {
	if obj.Session_id!="" {
		return "Bearer "+obj.User_id + " "+obj.Session_id
	}else{
		return ""
	}
}