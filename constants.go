
package alice_blue_api

const BASE_URL = `https://a3.aliceblueonline.com/rest/AliceBlueAPIService/api/`

// urls
const Encryption_key = "customer/getAPIEncpkey"
const Getsessiondata = "customer/getUserSID"

const Marketwatch_scrips = "marketWatch/fetchMWScrips"
const Addscrips= "marketWatch/addScripToMW"
const Getmarketwatch_list= "marketWatch/fetchMWList"
const Scripdetails= "ScripDetails/getScripQuoteDetails"
const Getdelete_scrips= "marketWatch/deleteMWScrip"

const Squareoffposition= "positionAndHoldings/sqrOofPosition"
const Position_conversion= "positionAndHoldings/positionConvertion"
const Placeorder= "placeOrder/executePlaceOrder"
const Modifyorder= "placeOrder/modifyOrder"
const Marketorder= "placeOrder/executePlaceOrder"
const Exitboorder= "placeOrder/exitBracketOrder"
const Bracketorder= "placeOrder/executePlaceOrder"
const Positiondata= "positionAndHoldings/positionBook"
const Orderbook= "placeOrder/fetchOrderBook"
const Tradebook= "placeOrder/fetchTradeBook"
const Holding= "positionAndHoldings/holdings"
const Orderhistory= "placeOrder/orderHistory"
const Cancelorder= "placeOrder/cancelOrder"
const Profile= "customer/accountDetails"

const Fundsrecord= "limits/getRmsLimits"

const Base_url_socket ="wss://ws1.aliceblueonline.com/NorenWS/"



// refer -- 
// https://github.com/jerokpradeep/pya3/blob/main/pya3/alicebluepy.py


// encyp key
// Y49T77B7MTOX5CBHGPZNQP5AAQ38T5HW
// sesson id
// uF2em9FPgjgLO5h6pbPuo4DxEyuYtC7YUECdT3TFQSAiZ6mhpB758i5iu1BZ5KmjaWj8ziBaZ9cxRVWicZOT8NkkGdLd7Q10zd1uTgmVagvNgmhD5ejY3YKay9XuR8e67nSerIXmOf6dO8gOLBusqPG89tb2nrfDeUa49ejGpt4WCjZlMcBZftBhv16SDLdAqAoSUW0NVrJQs4DPAUoeu0tu8DBBE6cbletCkmHJhEmegzHeEaDSbqE1DdVl2Hqp


/*


{
    "k":"NSE|3923","t":"t"
}
*/