package log

import (
	"fmt"
	"time"
)

func ErrPrint(err error) {
	fmt.Printf("[ SCV ] - [%s] \"Error: %s\"\n", time.Now(), err)
}

func LogPrint(log string) {
	fmt.Printf("[ SCV ] - [%s] \" %s\"\n", time.Now(), log)
}