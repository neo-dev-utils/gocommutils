package idUtil

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// GetToken 获取token
func GetToken() string {
	token := uuid.NewV4()
	t := fmt.Sprintf("%v", token)
	return t
}
