package profiledb

import (
	"pawpawchat/pkg/profile/repository"
)

type postgresdb struct {
	profile repository.Profile
}

func newpostgres(profile repository.Profile) Database {
	return &postgresdb{profile}
}

func (p *postgresdb) Profile() repository.Profile {
	return p.profile
}
