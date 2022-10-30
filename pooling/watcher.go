package pooling

import (
	"fmt"
	"time"
)

func RunWatcher() {
	for range time.Tick(time.Second * 5) {
		fmt.Println("Foo")
	}
}
