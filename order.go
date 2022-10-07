package alice_blue_api

// import (
// 	"fmt"
// )

// // https://v2api.aliceblueonline.com/orderManagement

// type Order struct{
// 	Discqty string `json:"discqty"`
// 	Trading_symbol string `json:"trading_symbol"`
// 	Exch string `json:"exch"`
// 	Transtype string `json:"transtype"`
// 	Ret string `json:"ret"`
// 	Prctyp string `json:"prctyp"`
// 	Qty string `json:"qty"`
// 	Symbol_id string `json:"symbol_id"`
// 	Price string `json:"price"`
// 	TrigPrice string `json:"trigPrice"`
// 	PCode string `json:"pCode"`
// 	Complexty string `json:"complexty"`
// 	OrderTag string `json:"orderTag"`
// }

// // order to string
// func orderToString(order Order) (string){
// 	q := "{"

// 	if order.Discqty!="" {
// 		q += fmt.Sprintf(`"discqty":"%s"`,order.Discqty)
// 	}

// 	if order.Trading_symbol!="" {
// 		q += fmt.Sprintf(`"trading_symbol":"%s"`,order.Trading_symbol)
// 	}

// 	if order.Exch!="" {
// 		q += fmt.Sprintf(`"exch":"%s"`,order.Exch)
// 	}
// 	if order.Transtype!="" {
// 		q += fmt.Sprintf(`"transtype":"%s"`,order.Transtype)
// 	}
// 	if order.Ret!="" {
// 		q += fmt.Sprintf(`"ret":"%s"`,order.Ret)
// 	}
// 	if order.Prctyp!="" {
// 		q += fmt.Sprintf(`"prctyp":"%s"`,order.Prctyp)
// 	}
// 	if order.Qty!="" {
// 		q += fmt.Sprintf(`"qty":"%s"`,order.Qty)
// 	}
// 	if order.Qty!="" {
// 		q += fmt.Sprintf(`"qty":"%s"`,order.Qty)
// 	}
// 	if order.Qty!="" {
// 		q += fmt.Sprintf(`"qty":"%s"`,order.Qty)
// 	}
// 	if order.Qty!="" {
// 		q += fmt.Sprintf(`"qty":"%s"`,order.Qty)
// 	}
// }

// // place order
// func PlaceOrder(data Order){
// 	url := BASE_URL+"placeOrder/executePlaceOrder"	


// }


// /*
// 	 [{ 
//           "discqty":"0",
//           "trading_symbol":"CRUDEOIL19APR3600CE",
//           "exch":"MCX",
//           "transtype":"BUY", 
//           "ret":" DAY ",  
//           "prctyp":"L",
//           "qty":"2",
//           "symbol_id":"5678943",
//           "price":"3550.00",
//           "trigPrice":"00.00",
//           "pCode":"MIS",
//           "complexty":"REGULAR",
//       "orderTag":"order1"
//       }] 
// */