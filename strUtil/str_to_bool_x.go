package strUtil

// StrToBool string转bool
func StrToBool(v string) bool {
	i := StrToUint64(v)
	return i == 1
}
