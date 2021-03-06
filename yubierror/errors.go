package yubierror

type YubiKeyError uint32

const (
	_                               = iota
	ErrorChkWrong      YubiKeyError = iota
	ErrorWrongPassword YubiKeyError = iota
	ErrorUserCancled   YubiKeyError = iota
)

func (e YubiKeyError) Error() string {
	switch e {
	case ErrorChkWrong:
		return "CHK wrong (this is a bug)"
	case ErrorWrongPassword:
		return "Wrong YubiKey password"
	case ErrorUserCancled:
		return "User canceled"
	}
	return "unknown error"
}
