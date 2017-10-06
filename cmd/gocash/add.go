package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/alyyousuf7/gocash/transaction"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a transaction",
		RunE:  Add,
	}
)

func init() {
	addCmd.Flags().StringP("time", "t", "", "Time (3:02pm)")
	addCmd.Flags().IntP("amount", "a", 0, "Amount")
	addCmd.Flags().StringP("note", "n", "", "Note")
}

func Add(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("%s command does not take any arguments", os.Args[0])
	}

	timeStr, err := cmd.Flags().GetString("time")
	if err != nil {
		return err
	}

	parsedTime, err := parsePartialTime(timeStr)
	if err != nil {
		return err
	}

	amount, err := cmd.Flags().GetInt("amount")
	if err != nil {
		return err
	}

	if amount == 0 {
		return fmt.Errorf("Amount should be greater than 0")
	}

	note, err := cmd.Flags().GetString("note")
	if err != nil {
		return err
	}

	if err := app.AddTransaction(parsedTime, amount, note); err != nil {
		return err
	}
	fmt.Println("Transaction added.")

	summary, err := app.GetSummary()
	if err != nil {
		return err
	}

	fmt.Printf("Your new balance: %s %d\n", transaction.CURRENCY, summary.Amount())
	return nil
}

func parsePartialTime(value string) (time.Time, error) {
	if value == "" {
		return time.Now(), nil
	}

	type partialFormat struct {
		formats []string
		fix     func(time.Time) time.Time
	}

	timeFormats := []partialFormat{
		{
			formats: []string{
				"2006-01-02 15:04",
				"2006-01-02 03:04pm",
				"2006-01-02 3:04pm",
				"2006-01-02 03pm",
				"2006-01-02 3pm",
			},
			fix: func(t time.Time) time.Time {
				return t
			},
		}, {
			formats: []string{
				"01-02 15:04",
				"01-02 03:04pm",
				"01-02 3:04pm",
				"01-02 03pm",
				"01-02 3pm",
			},
			fix: func(t time.Time) time.Time {
				return t.AddDate(time.Now().Year(), 0, 0)
			},
		}, {
			formats: []string{
				"15:04",
				"03:04pm",
				"3:04pm",
				"03pm",
				"3pm",
			},
			fix: func(t time.Time) time.Time {
				return t.AddDate(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
			},
		},
	}

	var err error
	for _, pFormat := range timeFormats {
		for _, format := range pFormat.formats {
			parsedTime, terr := time.Parse(format, value)

			if terr == nil {
				return pFormat.fix(parsedTime), nil
			}

			err = terr
		}
	}

	return time.Time{}, err
}