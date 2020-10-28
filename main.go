package main

import (
	"log"
	"os"

	"github.com/muhammadisa/go-rest-boilerplate/api"
	"github.com/muhammadisa/gosqlexec"
	"github.com/urfave/cli"
)

func main() {
	qe := gosqlexec.GoSQLExec{
		AlterQuery:  "db/alter/alter_tables.sql",
		DropQuery:   "db/drop/drop_tables.sql",
		CustomQuery: "db/query/custom_query.sql",
		Schemas: []string{
			"db/schemas/foobars.sql",
			"db/schemas/users.sql",
		},
	}

	app := cli.NewApp()
	app.Name = "Restful Boilerplate Service"
	app.Usage = "Restful Boilerplate service CLI tools"

	app.Commands = []cli.Command{
		gosqlexec.MigrateCommand(qe),
		gosqlexec.DropTablesCommand(qe),
		gosqlexec.AlterTablesCommand(qe),
		gosqlexec.CustomQueryExecCommand(qe),
		{
			Name:  "run-server",
			Usage: "Start Server",
			Action: func(c *cli.Context) error {
				api.Run()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
