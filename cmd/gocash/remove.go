package main

import (
	"fmt"
	"os"

	"github.com/alyyousuf7/gocash/transaction"
	"github.com/spf13/cobra"
)

var (
	removeCmd = &cobra.Command{
		Use:   "rm",
		Short: "Remove a transaction",
		RunE:  Remove,
	}
)

func init() {
	removeCmd.Flags().String("id", "", "Transaction ID")
}

func Remove(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("%s command does not take any arguments", os.Args[0])
	}

	id, err := cmd.Flags().GetString("id")
	if err != nil {
		return err
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
