// https://go.dev/play/p/pBGSnZ_Kjxo

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2022, time.November, 10, 13, 22, 23, 24, time.UTC)
	// 2009.11.10.12223
	fmt.Println(t.Format("2006.1.2.30405"))
}
