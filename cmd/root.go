package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// List of variables used as flags
var (
	AccountNumber string
	Email         string
	MFAAccount    string
	ProfileName   string
	Role          string
)

// Config struct used to build template
type Config struct {
	AccountNumber string
	AwsKey        string
	AwsSecret     string
	Email         string
	MFAAccount    string
	ProfileName   string
	Role          string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cyberduck-s3-config",
	Short: "CLI tool to manage config file for cyberduck",
	Long: `A command line tool to facilitate full management of the configuration
file required by Cyberduck in order to be to assume roles using AWS STS.`,
}

// Execute funtion part of core cobra code
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Flags required to produce config file
	rootCmd.PersistentFlags().StringVarP(&Email, "email", "e", "", "Your email")
	rootCmd.PersistentFlags().StringVarP(&Role, "role", "r", "", "Your role")
	rootCmd.PersistentFlags().StringVarP(&ProfileName, "name", "n", "", "The name of the profile you'd like to create")
	rootCmd.PersistentFlags().StringVarP(&AccountNumber, "account", "a", "", "The account number you wish to access")
	rootCmd.PersistentFlags().StringVarP(&MFAAccount, "mfa", "m", "", "The account number you'd like to MFA from")

}
