package main

import (
	"fmt"
	"os"

	"github.com/alyyousuf7/gocash/transaction"
	"github.com/spf13/cobra"
)

var (
	removeCmd = &cobra.Command{
		Use:   "rm <transaction id>",
		Short: "Remove a transaction",
		RunE:  Remove,
	}
)

func Remove(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("%s only needs trasaction ID to remove", os.Args[0])
	}

	id := args[0]
	if len(id) == 0 {
		return fmt.Errorf("Empty ID was detected")
	}

	if err := app.RemoveTransaction(id); err != nil {
		return err
	}
	fmt.Println("Transaction removed.")

	summary, err := app.GetSummary()
	if err != nil {
		return err
	}

	fmt.Printf("Your new balance: %s %d\n", transaction.CURRENCY, summary.Amount())
	return nil
}
