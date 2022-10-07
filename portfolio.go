package alice_blue_api

import (
	"fmt"
	"strings"
	"errors"
)


// function to get position book
func (obj *Auth) PositionBook() (map[string]interface{},error){
	url := "https://a3.aliceblueonline.com/rest/AliceBlueAPIService/api/positionAndHoldings/positionBook"
	
	ret := "DAY"

	payload := fmt.Sprintf(`{ "ret":"%s"}`,ret)

	fmt.Println(payload)
	data := strings.NewReader(payload)

	response , err := PostRequest(url, obj, data)

	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	
	if(response["stat"]=="Ok"){
		return response,nil
	}

	return nil,errors.New("Something Happen Bad in position book id function may be their is nothing in it")
}

func (obj *Auth) GetHoldings() (map[string]interface{},error) {
	url := "https://a3.aliceblueonline.com/rest/AliceBlueAPIService/api/positionAndHoldings/holdings"

	response , err := GetRequest(url, obj)

	if err!=nil{
		fmt.Println("Error in Holdings",err)
		return nil,err
	}

	if(response["stat"]=="Ok"){
		return response,nil
	}
	return nil,errors.New("Something Happen Bad in Get holding function") 
}