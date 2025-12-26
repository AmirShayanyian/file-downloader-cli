package cli

import (
	"fmt"

	"github.com/AmirShayanyian/file-downloader-cli/internal/queue"
	"github.com/AmirShayanyian/file-downloader-cli/internal/storage"
	"github.com/spf13/cobra"
)

func init() {
	queueCmd := &cobra.Command{
		Use:   "queue",
		Short: "Manage download queues",
	}
	queueCmd.AddCommand(createQueueCmd())

	rootCmd.AddCommand(queueCmd)
}

func createQueueCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create [name]",
		Short: "Create a new queue",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			store := storage.NewJSONStore("data/queues.json")

			q := queue.Queue{
				Name: args[0],
			}

			if err := store.Save(q); err != nil {
				return err
			}

			fmt.Println("Queue created:", q.Name)
			return nil
		},
	}
}

func listQueueCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List queues",
		RunE: func(cmd *cobra.Command, args []string) error {
			store := storage.NewJSONStore("data/queues.json")

			queues, err := store.List()
			if err != nil {
				return err
			}

			if len(queues) == 0 {
				fmt.Println("No queues found.")
				return nil
			}

			for _, q := range queues {
				fmt.Println("-", q.Name)
			}

			return nil
		},
	}
}
