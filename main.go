package main

import (
    "github.com/spf13/cobra"
    "github.com/MMazoni/most-used-features/internal/commands"
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "muf",
        Short: "A tool for analyzing access logs",
        Long:  "A tool for analyzing access logs and generating CSV files that contain the most accessed features.",
    }
    accessCmd := &cobra.Command{
        Use:   "access",
        Short: "Analyze access logs and generate CSV file",
        Long:  `Analyze access logs and generate a CSV file that contains the most accessed features.`,
        Run:   commands.AccessCommand,
    }
    csrfCmd := &cobra.Command{
        Use:   "csrf",
        Short: "Analyze csrf logs and generate CSV file",
        Long:  `Analyze csrf logs and generate a CSV file that contains the most url errors.`,
        Run:   commands.CsrfCommand,
    }
    rootCmd.AddCommand(csrfCmd)
    rootCmd.AddCommand(accessCmd)
    rootCmd.Execute()
}
