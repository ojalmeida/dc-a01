package db

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-api/config"
	"go-api/entities"
)

var db *sqlx.DB

func init() {

	var err error

	db, err = sqlx.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",

		config.Config.Network.DB.User,
		config.Config.Network.DB.Password,
		config.Config.Network.DB.Host,
		config.Config.Network.DB.Port,
		config.Config.Network.DB.DatabaseName,
	),
	)

	if err != nil {
		panic(err.Error())
	}

	db.SetMaxOpenConns(config.Config.Performance.MaxNumberOfWorkers)
	db.SetMaxIdleConns(config.Config.Performance.MaxNumberOfWorkers)

}

/*
CreateItem insert an item in the database with the connection established returning the inserted item.
If ctx is canceled the insertion also is canceled.
*/
func CreateItem(ctx context.Context, i entities.Item) (item entities.Item, err error) {

	err = db.GetContext(ctx, &item, `INSERT INTO tb01 (col_texto, col_dt) VALUES ($1, $2) RETURNING *`, i.Text, i.Date)

	return

}
