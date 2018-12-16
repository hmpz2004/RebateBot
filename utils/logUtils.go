package utils

import (
	"fmt"
)

func LogDebug(str ...string) {
	content := ""
	if len(str) == 1 {
		content = str[0]
	} else if  len(str) == 2 {	
		// content = str[0] + ":" + str[1]
		content = fmt.Sprintf("%-25s:%s", str[0], str[1])
	}
	fmt.Println(content)
}

func logRelease(str string) {
	fmt.Println(str)
}
