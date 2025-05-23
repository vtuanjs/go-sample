package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrInvalidToken     = 30001
	ErrCodeUserHasExist = 50002
	ErrCodeInvalidParam = 50003
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Invalid parameter",
	ErrInvalidToken:     "Invalid token",
	ErrCodeUserHasExist: "User has exist",
	ErrCodeInvalidParam: "Invalid parameter",
}
