package reverse

import (
	"cheetah/internal/database"
	"cheetah/internal/reverse"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

type builder struct {
	*strings.Builder
}

func (b builder) WriteString(s string) {
	b.Builder.WriteString(s)
}

func (b builder) WriteRune(r rune) {
	b.Builder.WriteRune(r)
}

func (b builder) WriteByte(c byte) {
	b.Builder.WriteByte(c)
}

func (b builder) WriteQuoteTo(s string) {
	b.Builder.WriteByte('`')
	b.Builder.WriteString(s)
	b.Builder.WriteByte('`')
}

func (b builder) String() string {
	return b.Builder.String()
}

var ReverseCommand = cobra.Command{
	Use:   "reverse",
	Short: "reverse ddl statement from database",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("table") {
			host, _ := cmd.InheritedFlags().GetString("host")
			port, _ := cmd.InheritedFlags().GetInt("port")
			user, _ := cmd.InheritedFlags().GetString("user")
			password, _ := cmd.InheritedFlags().GetString("password")
			schema, _ := cmd.InheritedFlags().GetString("database")
			config := database.Config{Host: host, Port: port, User: user, Password: password, DBName: schema}
			db, err := database.Create(&config)
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()
			table, err := reverse.RevCreateTableStmt(db, schema, tableName)
			if err != nil {
				log.Fatal(err)
			}
			budler := builder{Builder: &strings.Builder{}}
			table.Build(&budler)
			log.Println(budler.String())
		}
	},
}

var (
	tableName string
)

func init() {
	setupFlags()
}

func setupFlags() {
	ReverseCommand.Flags().StringVarP(&tableName, "table", "t", "", "table name")
}
