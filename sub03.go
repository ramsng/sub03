package sub03

import (
	"strings"

	"github.com/ramsng/sub02"
)

func Praise() string {
	return "\n Nice spin - Keep it up" + strings.ToUpper(sub02.Desc())
}
