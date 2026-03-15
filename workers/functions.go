package workers

import (
	"context"
	"fmt"
	"msg_queue/config"
)

func StartWorkers(ctx context.Context, id int) {
	for {
		result, err := config.Rdb.BRPop(ctx, 0, "messageQueue").Result()
		if err != nil {

			if ctx.Err() != nil {
				fmt.Println("Worker", id, "stopping")
				return
			}

			fmt.Println("worker error:", err)
			continue
		}
		msg := result[1]

		fmt.Println("Worker", id, "processing:", msg)

		processMessage(msg)
	}
}

func processMessage(msg string) {
	fmt.Println("Processing:", msg)
}
