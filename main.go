package main

import (
	"bufio"
	"flag"
	"fmt"
	"gitea.hbanafa.com/hesham/mystc-api/pkg/endpoints"
	"gitea.hbanafa.com/hesham/mystc-api/pkg/models"
	"os"
	"strings"
)

const URL_MYSTC_API_BASE = "https://mystc.stc.com.sa"
const PATH_PHONES_LIST = "/api/mystc-api-authentication/phones-list"

func main() {

    id := flag.String("id", "", "National ID or Iqama")
    pw := flag.String("password", "", "MySTC Password")
    flag.Parse()

    phones, err := endpoints.GetPhonesList(*id, *pw)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    reader := bufio.NewReader(os.Stdin)
    var selectedNumber models.PhoneNumber

    mainloop:
    for {
        for idx, n := range phones.PhoneNumbers {
            fmt.Printf("%d: %s %s\n", idx, n.Number, n.Type.String())
        }
        fmt.Printf("Select a number: ")
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
        }
        input = strings.Trim(input, "\n")
        for idx, n := range phones.PhoneNumbers {
            if input == fmt.Sprint(idx) {
                selectedNumber = n
                break mainloop
            }
        }
        
    }

    fmt.Printf("selectedNumber: %v\n", selectedNumber)

    loginOTP, err := endpoints.RequestLoginOTP(phones.LoginToken, selectedNumber.Number)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Println("OTP Sent!")
    fmt.Print("Enter OTP: ")
    otp, err := reader.ReadString('\n')
    otp = strings.Trim(otp, "\n")

    login, err := endpoints.VerifyLoginOTP(*id, loginOTP.OtpToken, otp)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Println("Logged in!")
    balance, err := endpoints.GetBalance(login, selectedNumber.Number)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Printf("Balance %f\n", balance.Balance)
    
    


}
