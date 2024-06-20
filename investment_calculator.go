package main // To know Golang where begin to start

import (
	"fmt"
	"math"
)

func main() {
	// var investmentAmount float64 = 1000
	// ဒီလို  ကြော်ညာတာက ပုံမှန်ဘဲ။ ကိုလိုချင်တဲ့  Data Type  တစ်ခါထဲ သတ်မှတ်ထားတာ

	// var expectedReturnRate = 5.5
	// ဒါကတော့ ကိုထည့်လိုက်တဲ့ value ပေါ်မူတည်ပြီး Data Type သတ်မှတ်တာ

	//expectedReturnRate := 5.5
	// ဒါကတော့ var ထည့်စရာ မလိုတဲ့ ပုံစံ, Data Type လဲ  value ပေါ်မူတည်ပြီး Data Type သတ်မှတ်တာ, recommand လဲ  ပေးထားပါတယ်။

	// var investmentAmount, years = 1000, "10"
	// ဒါကတော့ variable 2 ခုကို တစ်ခါထဲ  ကြော်ညာတာ

	var investmentAmount float64 = 1000
	var years float64 = 10
	expectedReturnRate := 5.5

	var amount float64 = 1.2
	fmt.Scan(&amount)
	//fmt.Scan သည် User ဆီကနေ Terminal ကနေပြီး Data Input လှမ်းတောင်းချင်သော အခါတွင် သုံးသည်။ variable ရှေ့တွင် & ထည့်ပေးရမည်။

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}

// 20
