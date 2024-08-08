package profiledb

import "pawpawchat/pkg/profile/repository"

type Database interface {
	Profile() repository.Profile
}

type Factory interface {
	OpenProfileDB(dsn string) (Database, error)
}
