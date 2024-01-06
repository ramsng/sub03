package sub03

import (
	"fmt"
	"strings"

	"github.com/ramsng/sub02"
)

func Praise() string {
	fmt.Printf("comments in sub03", sub02.Desc())
	return "\n Nice spin - Keep it up" + strings.ToUpper(sub02.Desc())
}
