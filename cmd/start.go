package cmd

import (
	"fmt"
	"github.com/sakiib/apiServer/api"
	"github.com/spf13/cobra"
	"os"
)

var username string
var password string

var port string
var authNeeded bool

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called! start the server from this point..")
		api.HandleRoutes(username, password, port, authNeeded)
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
	startCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "This flag will set the port, default 8080")
	startCmd.PersistentFlags().BoolVarP(&authNeeded, "auth", "a", true, "This flag will set the auth requirement")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
