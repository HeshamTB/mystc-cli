package models

type LoginPhoneListRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginPhoneListResponse struct {
    PhoneNumbers []PhoneNumber `json:"phoneNumbers"`
    LoginToken string `json:"loginToken"`
}

type PhoneNumberType string
var (
    POSTPAID PhoneNumberType = "1"
    PREPAID PhoneNumberType = "2"
)

func (p PhoneNumberType) String() string {
    switch p {
    case "1":
        return "POSTPAID"
    case "2":
        return "PREPAID"
    default:
        return "UNKOWN"
    }
}

type PhoneNumber struct {
    Number string `json:"number"`
    Type PhoneNumberType `json:"type"`
    Contact string `json:"contact"`
}

type LoginRequest struct {
    LoginToken string `json:"loginToken"`
    PhoneNumber string `json:"phoneNumber"`
}

type LoginResponse struct {
    OtpToken string `json:"otpToken"`
    ExpiresIn string `json:"expiresIn"`
    TokenType string `json:"tokenType"`
}

type LoginVerificationRequest struct {
    Username string `json:"username"`
    OtpToken string `json:"otpToken"`
    Otp string `json:"otp"`
}

type LoginVerificationResponse struct {
    AccessToken string `json:"accessToken"`
    TokenType string `json:"tokenType"`
    ExpiresIn string `json:"expiresIn"`
    RefreshToken string `json:"refreshToken"`
}
