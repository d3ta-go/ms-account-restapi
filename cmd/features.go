package cmd

import (
	"github.com/d3ta-go/project-template/cmd/db"
	"github.com/d3ta-go/project-template/cmd/server"
)

func init() {
	RootCmd.AddCommand(db.DBCmd)
	RootCmd.AddCommand(server.ServerCmd)
	// Add your custom command
}
