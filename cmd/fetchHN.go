package cmd

import (
	"buzzGen/service"
	"github.com/spf13/cobra"
)

// fetchHNCmd represents the fetchHN command
var fetchHNCmd = &cobra.Command{
	Use:   "fetchHN",
	Short: "fetchHN",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		_ = service.FetchHnData()
	},
}

func init() {
	rootCmd.AddCommand(fetchHNCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchHNCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchHNCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
