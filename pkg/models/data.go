package models

type PrePaidBalanceResponse struct {
    PhoneNumber string `json:"phoneNumber"`
    Balance float64 `json:"balance"`
    CurrencyCode string `json:"currencyCode"`
    AirTimeExpiryDate string `json:"airtimeExpiryDate"`
    AccountExpiryDate string `json:"accountExpiryDate"`
    CustomerId string `json:"customerId"`
    ValidationType string `json:"validationType"`
    TotalOtherBalances float64 `json:"totalOtherBalances"`
    TotalFrozenBalances float64 `json:"totalFrozenBalances"`
}

