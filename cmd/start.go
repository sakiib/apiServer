package cmd

import (
	"apiServer/api"
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var username string
var password string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called! start the server from this point..")
		api.HandleRoutes(username, password)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	username = os.Getenv("username")
	password = os.Getenv("password")
	fmt.Println("username: ", username, " & ", "password: ", password)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
