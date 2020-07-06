package main

// To generate the models run: go run ./cmd/sqlbunny gen
//To generate the SQL code to crreate the tables run: go run ./cmd/sqlbunny migration gensql
import (
	. "github.com/sqlbunny/sqlbunny/gen/core"
	"github.com/sqlbunny/sqlbunny/gen/migration"
	"github.com/sqlbunny/sqlbunny/gen/stdtypes"
)

func main() {
	Run(
		&stdtypes.Plugin{},
		&migration.Plugin{},

		Model("book",
			Field("id", "string", PrimaryKey),
			Field("created_at", "time"),
			Field("name", "string"),
		),
	)
}