package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("x = ", x)
	/*	nilai x tidak akan mencapai 100.000 dan berubah-ubah setiap kali dijalankan.
		Hal ini karena setiap goroutine mengakses variable x secara bersamaan.
	 */
}
