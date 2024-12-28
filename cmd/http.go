package cmd/home/amir/GitHub/without-redis/logger/home/amir/GitHub/without-redis/logger

import (/home/amir/GitHub/without-redis/logger
	"github.com/docker/docker/daemon/logger"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "launching the http rest listen server",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Logger.Info("http rest server is starting")
		http.Start()
	},
}
