package cmd

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "HTTP GET request",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			log.Info("requesting " + arg)
			res, err := http.Get(arg)
			if err != nil {
				log.Error(err)
				continue
			}
			log.Info(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
