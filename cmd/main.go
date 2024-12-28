package main

import (
	"cheetah/cmd/reverse"
	"github.com/spf13/cobra"
	"os"
)

var root = cobra.Command{
	Use: "cheetah",
}

var (
	host     string
	port     int
	user     string
	password string
	dbName   string
)

func setupFlags() {
	root.PersistentFlags().StringVarP(&host, "host", "H", "localhost", "host to connect to")
	root.PersistentFlags().IntVarP(&port, "port", "P", 3306, "port to connect to")
	root.PersistentFlags().StringVarP(&user, "user", "u", "root", "user to connect to")
	root.PersistentFlags().StringVarP(&password, "password", "p", "root", "password to connect to")
	root.PersistentFlags().StringVarP(&dbName, "database", "D", "cheetah", "database to connect to")
}

func setupSubcommand() {
	root.AddCommand(&reverse.ReverseCommand)
}

func init() {
	setupFlags()
	setupSubcommand()
}

func main() {
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
