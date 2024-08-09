package profiledb

import (
	"pawpawchat/pkg/profile/repository/orm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresFactory ...
type PostgresFactory struct{}

func NewPostgresFactory() Factory {
	return &PostgresFactory{}
}

// OpenProfileDB ...
func (f *PostgresFactory) OpenProfileDB(dsn string) (Database, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	profile := orm.NewPostgresProfileRepository(db)
	return newpostgres(profile), nil
}
