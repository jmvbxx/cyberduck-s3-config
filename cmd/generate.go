package cmd

import (
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "To generate a configuration file",
	Run:   generate,
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generate(cmd *cobra.Command, args []string) {

	awsKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecret := os.Getenv("AWS_SECRET_ACCESS_KEY")

	conf := Config{
		AccountNumber,
		awsKey,
		awsSecret,
		Email,
		MFAAccount,
		ProfileName,
		Role,
	}

	// Empty line at the end added for babrams :-)
	t := template.Must(template.New("generateTemplate.txt").ParseFiles("generateTemplate.txt"))

	// I will correct the file and path once the rest of the repo is complete and
	// approved. For now it's `test.txt`. It will be `~/.aws/credentials`
	f, err := os.Create("test.txt")
	if err != nil {
		log.Panic(err)
	}

	err = t.Execute(f, conf)
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
}
