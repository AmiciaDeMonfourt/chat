package profiledb

import "pawpawchat/pkg/profile/repository"

// Database ...
type Database interface {
	Profile() repository.Profile
}

// Factory ...
type Factory interface {
	OpenProfileDB(dsn string) (Database, error)
}
