package factory

import (
	"log"
	"log/slog"
	"pawpawchat/pkg/profile/repository"
	"pawpawchat/pkg/profile/repository/orm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepositoryFactory struct{}

func NewPostgresRepositoryFactory() RepositoryFactory {
	return &PostgresRepositoryFactory{}
}

func (f *PostgresRepositoryFactory) OpenProfile(dsn string, repotype string) repository.Profile {
	switch repotype {
	case "orm":
		db, err := gorm.Open(postgres.Open(dsn))
		if err != nil {
			log.Fatal(err)
		}
		slog.Debug("connection with ppc_profile", "dsn", dsn)
		return orm.NewPostgresProfileRepository(db)

	default:
		log.Fatal("unsupported repository type")
		return nil
	}
}
