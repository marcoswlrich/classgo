package interactions

import (
	"fmt"
)

func Interar() {
	for i := 10; i > 1; i-- {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
