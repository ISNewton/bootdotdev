package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	// ?

	emailStruct, ok := e.(email)

	if ok {
		return emailStruct.toAddress, emailStruct.cost()
	}

	smsStruct, ok := e.(sms)

	if ok {
		return smsStruct.toPhoneNumber, smsStruct.cost()
	}
	return "", 0

}

func main() {

	myEmail := email{toAddress: "hi@google.com", body: "HOw Are you", isSubscribed: true}

	fmt.Println(getExpenseReport(myEmail))

	mySms := sms{toPhoneNumber: "1212", body: "HOw Are you", isSubscribed: true}

	fmt.Println(getExpenseReport(mySms))
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
