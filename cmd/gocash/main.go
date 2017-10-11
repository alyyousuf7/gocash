package main

import (
	"fmt"
	"os"
	"path"

	"github.com/alyyousuf7/gocash"
	"github.com/alyyousuf7/gocash/transaction"
	"github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func main() {
	mainCmd.Execute()
}

var (
	app     *gocash.App
	mainCmd = &cobra.Command{
		Use:   os.Args[0],
		Short: "Petty cash summary",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				return err
			}

			config, err := loadConfiguration(configPath)
			if err != nil {
				return err
			}

			app, err = gocash.NewApp(config)
			if err != nil {
				return err
			}

			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return app.Close()
		},
		RunE: Summary,
	}
)

func init() {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	configPath := path.Join(homeDir, ".gocash.yml")

	mainCmd.PersistentFlags().StringP("config", "c", configPath, "Config file path")
	mainCmd.AddCommand(
		debitCmd,
		creditCmd,
		removeCmd,
	)
}

func Summary(cmd *cobra.Command, args []string) error {
	summary, err := app.GetSummary()
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)

	table.SetHeader([]string{
		"ID",
		"Date",
		"Note",
		"Amount",
	})

	for _, tx := range summary {
		table.Append([]string{
			tx.ID(),
			tx.Time().Format("Jan 02"),
			tx.Note(),
			formatAmount(tx.Amount()),
		})
	}

	table.SetFooter([]string{
		"",
		"",
		"Total",
		formatAmount(summary.Amount()),
	})
	table.Render()
	return nil
}

func formatAmount(amount int) string {
	if amount < 0 {
		return fmt.Sprintf("- %s %d", transaction.CURRENCY, -amount)
	}
	return fmt.Sprintf("+ %s %d", transaction.CURRENCY, amount)
}
