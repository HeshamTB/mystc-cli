package endpoints

import (
	"encoding/json"
	"fmt"
	"io"

	"mystcapi/internal/urls"
	"mystcapi/pkg/models"
	"net/http"
)

func GetBalance(
    login models.LoginVerificationResponse, 
    phoneNumber string) (models.PrePaidBalanceResponse, error) {

    prepaidBalanceReq := models.PrePaidBalanceResponse{}

    req, err := makeGetPrepaidBalanceReq(phoneNumber) 
    if err != nil {
        return prepaidBalanceReq, err
    }
    req.Header.Set("Authorization", fmt.Sprintf("%s %s", login.TokenType, login.AccessToken))


    resp, err := client.Do(req)
    if err != nil {
        return prepaidBalanceReq, err
    }
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return prepaidBalanceReq, err
    }

    err = json.Unmarshal(body, &prepaidBalanceReq)
    return prepaidBalanceReq, err
}

func makeGetPrepaidBalanceReq(phoneNumber string) (*http.Request, error) {
    return  http.NewRequest(
        http.MethodGet,
        fmt.Sprintf(urls.PATH_PREPAID_BALANCE_F, phoneNumber),
        nil,
    )
}

