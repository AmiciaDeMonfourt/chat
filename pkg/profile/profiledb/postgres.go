package profiledb

import (
	"pawpawchat/pkg/profile/repository"
	"pawpawchat/pkg/profile/repository/orm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postrgres ...
type Postgres struct {
	profile repository.Profile
}

func NewPostgresDB(profile repository.Profile) Database {
	return &Postgres{profile}
}

func (p *Postgres) Profile() repository.Profile {
	return p.profile
}

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
	return NewPostgresDB(profile), nil
}
