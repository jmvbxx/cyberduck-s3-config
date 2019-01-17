package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "To delete an existing config file",
	Long: `Run the subcommand with the -y flag in order
to delete your existing configuration file

CAUTION:
Once your file is deleted there is no way to restore
it. You will have to reconstruct it.`,
	Run: delete,
}

// TODO: force the use of -y flag to delete otherwise exit

func delete(cmd *cobra.Command, args []string) {
	err := os.Remove("test.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
