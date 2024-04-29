package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mystcapi/internal/constants"
	"mystcapi/internal/urls"
	"mystcapi/pkg/models"
	"net/http"
)


func GetPhonesList(id, password string) (models.LoginPhoneListResponse, error) {

    phoneList := models.LoginPhoneListResponse{}

    reqJson, err := json.Marshal(models.LoginPhoneListRequest{
        Username: id,
        Password: password,
    })
    if err != nil {
        return phoneList, err
    }

    resp, err := client.Post(
        urls.PATH_PHONES_LIST,
        constants.CONTENT_TYPE_JSON,  
        bytes.NewReader(reqJson),
    )
    if err != nil {
        return phoneList, err
    }

    if resp.StatusCode != http.StatusOK {
        return phoneList, fmt.Errorf("http: Login got %d", resp.StatusCode)
    }

    defer resp.Body.Close()


    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return phoneList, err
    }

    err = json.Unmarshal(body, &phoneList)
    return phoneList, err
}

func RequestLoginOTP(loginToken, phoneNumber string) (models.LoginResponse, error) {
    loginResp := models.LoginResponse{}

    reqJson, err := json.Marshal(
        models.LoginRequest{
            LoginToken: loginToken,
            PhoneNumber: phoneNumber,
        },
    )
    if err != nil {
        return loginResp, err
    }

    resp, err := client.Post(
        urls.PATH_LOGIN,
        constants.CONTENT_TYPE_JSON,
        bytes.NewReader(reqJson),
    )
    if err != nil {
        return loginResp, err
    }

    if resp.StatusCode != http.StatusOK {
        return loginResp, fmt.Errorf("http: RequestLoginOTP got %d", resp.StatusCode)
    }
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return loginResp, err
    }

    err = json.Unmarshal(body, &loginResp)
    return loginResp, err
}

func VerifyLoginOTP(
    username, otpToken, 
    otp string) (models.LoginVerificationResponse, error) {
    loginVerResp := models.LoginVerificationResponse{}

    reqJson, err := json.Marshal(
        models.LoginVerificationRequest{
            Username: username,
            OtpToken: otpToken,
            Otp: otp,
        },
    )
    if err != nil {
        return loginVerResp, err
    }

    resp, err := client.Post(
        urls.PATH_LOGIN_VERIFICATION,
        constants.CONTENT_TYPE_JSON,
        bytes.NewReader(reqJson),
    )
    if err != nil {
        return loginVerResp, err
    }
    if resp.StatusCode != http.StatusOK {
        return loginVerResp, fmt.Errorf("http: RequestLoginOTP got %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return loginVerResp, err
    }

    err = json.Unmarshal(body, &loginVerResp)
    return loginVerResp, err
}


