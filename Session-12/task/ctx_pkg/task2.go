package ctx_pkg

import (
	"context"
	"fmt"
	"time"
)

func SimulatesTask(ctx context.Context) {

	select {
	case <-time.After(time.Second * 3):
		fmt.Println("Task Started...")
	case <-time.After(time.Second * 5):
		fmt.Println("Task ended after 5 seconds")
	case <-ctx.Done():
		fmt.Println("Task cancelled due to timeout")
	}

}
