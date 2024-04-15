package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/tsukinoko-kun/http/internal"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "HTTP GET request",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			fmt.Println("requesting " + arg)
			start := time.Now()
			res, err := http.Get(arg)
			dur := time.Since(start)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			if res != nil {
				internal.PrintResponse(res, dur)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
