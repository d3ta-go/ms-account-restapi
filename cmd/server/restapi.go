package server

import (
	"fmt"

	"github.com/d3ta-go/project-template/interface/http-apps/restapi/echo"
	"github.com/spf13/cobra"
)

// restAPICmd represents the restapi server command
var restAPICmd = &cobra.Command{
	Use:   "restapi",
	Short: "Shows the restapi server command.",
	Long:  `Shows the restapi server command.`,
	Run: func(cmd *cobra.Command, args []string) {

		if err := echo.StartRestAPIServer(); err != nil {
			fmt.Println("Error while running `StartRestAPIServer()`: ")
			panic(err)
		}
	},
}

func init() {
	ServerCmd.AddCommand(restAPICmd)
}
