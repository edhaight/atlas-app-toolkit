package util

import (
	"github.com/infobloxopen/atlas-app-toolkit/util/cases"
)

// CamelToSnake is deprecated, use github.com/infobloxopen/util/cases.CamelToSnake.
func CamelToSnake(str string) string {
	return cases.CamelToSnake(str)
}
