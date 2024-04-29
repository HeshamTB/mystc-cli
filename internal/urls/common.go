package urls

const (
    URL_MYSTC_API_BASE = "https://mystc.stc.com.sa/api"
    PATH_API_AUTH_BASE = URL_MYSTC_API_BASE + "/mystc-api-authentication"
    PATH_PHONES_LIST = PATH_API_AUTH_BASE + "/phones-list"
    PATH_LOGIN = PATH_API_AUTH_BASE + "/login"
    PATH_LOGIN_VERIFICATION =  PATH_API_AUTH_BASE + "/login-verification"

    PATH_API_NUM_MGMT_BASE = URL_MYSTC_API_BASE + "/mystc-api-user-number-management/api"
    PATH_PREPAID_BALANCE_F = PATH_API_NUM_MGMT_BASE + "/phone-number/%s/prepaid-balance"
)

