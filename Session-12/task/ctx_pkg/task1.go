package ctx_pkg

import (
	"context"
	"fmt"
	"time"
)

func PrintNumbers(ctx context.Context) {

	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("(cancellation)")
		default:
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}

	}
}
