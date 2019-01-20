package cmd

import (
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

// appendCmd represents the append command
var appendCmd = &cobra.Command{
	Use:   "append",
	Short: "To append a new account to an existing configuration file",
	Run:   append,
}

func append(cmd *cobra.Command, args []string) {

	type Append struct {
		AccountNumber string
		Email         string
		MFAAccount    string
		ProfileName   string
		Role          string
	}
	conf := Append{
		AccountNumber,
		Email,
		MFAAccount,
		ProfileName,
		Role,
	}

	t := template.Must(template.New("appendTemplate.txt").ParseFiles("appendTemplate.txt"))

	// I will correct the file and path once the rest of the repo is complete and
	// approved. For now it's `test.txt`. It will be `~/.aws/credentials`

	// https://stackoverflow.com/questions/33851692/golang-bad-file-descriptor
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(f, conf)
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
}

func init() {
	rootCmd.AddCommand(appendCmd)

}
