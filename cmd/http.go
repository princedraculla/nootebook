package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nootebook.com/internal/gateway/http"
)

func init() {
	rootCmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "running http server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http server called")
		http.ServerInit()
	},
}
