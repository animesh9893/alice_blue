package alice_blue_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"io"
	"encoding/json"
)

// make post request return byte
func PostRequestReturnByte(url string,auth *Auth, data io.Reader) ([]byte,error){
	if auth.Bearear_token == "" {
		auth.GenerateBearearToken()
	}

	fmt.Println(url)

	client := http.Client{}

	req , err := http.NewRequest("POST", url, data)

	if err != nil {
		fmt.Println("Error in http/PostRequest")
		return nil,err
	}

	req.Header = http.Header{
	    "Content-Type": {"application/json"},
	    "Authorization": {"Bearer "+auth.Bearear_token},
	}

	resp , err := client.Do(req)
	if err != nil {
		fmt.Println("Error in http/PostRequest")
		return nil,err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in http/PostRequest")
		return nil,err
	}
	return body,nil
}

// post request return result in map
func PostRequest(url string,auth *Auth, data io.Reader) (map[string]interface{},error){
	if auth.Bearear_token == "" {
		auth.GenerateBearearToken()
	}

	fmt.Println(url)

	client := http.Client{}
	req , err := http.NewRequest("POST", url, data)
	if err != nil {
		fmt.Println("Error in http/PostRequest")
		return nil,err
	}

	req.Header = http.Header{
	    "Content-Type": {"application/json"},
	    "Authorization": {"Bearer "+auth.Bearear_token},
	}

	resp , err := client.Do(req)
	if err != nil {
		fmt.Println("Error in http/PostRequest")
		return nil,err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in http/PostRequest")
		return nil,err
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	return result,nil
}

// make get request return byte
func GetRequestReturnByte(url string,auth *Auth) ([]byte,error){
	if auth.Bearear_token == "" {
		auth.GenerateBearearToken()
	}
	fmt.Println(url)

	client := http.Client{}

	req , err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error in http/GetRequest")
		return nil,err
	}

	req.Header = http.Header{
	    "Content-Type": {"application/json"},
	    "Authorization": {"Bearer "+auth.Bearear_token},
	}

	resp , err := client.Do(req)
	if err != nil {
		fmt.Println("Error in http/GetRequest")
		return nil,err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in http/GetRequest")
		return nil,err
	}
	return body,nil
}

// Get requet return result in map
func GetRequest(url string,auth *Auth) (map[string]interface{},error){
	if auth.Bearear_token == "" {
		auth.GenerateBearearToken()
	}
	fmt.Println(url)

	client := http.Client{}

	req , err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error in http/GetRequest")
		return nil,err
	}

	req.Header = http.Header{
	    "Content-Type": {"application/json"},
	    "Authorization": {"Bearer "+auth.Bearear_token},
	}

	resp , err := client.Do(req)
	if err != nil {
		fmt.Println("Error in http/GetRequest")
		return nil,err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in http/GetRequest")
		return nil,err
	}	

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	return result,nil
}