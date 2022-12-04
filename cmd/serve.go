package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/oxodao/go-quickstart/config"
	"github.com/oxodao/go-quickstart/utils"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting " + utils.APP_NAME + " on addr " + config.GET.Web.ListeningAddr)

		r := mux.NewRouter()

		err := http.ListenAndServe(config.GET.Web.ListeningAddr, r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
