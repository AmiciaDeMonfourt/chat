package factory

import "pawpawchat/pkg/profile/repository"

type RepositoryFactory interface {
	OpenProfile(dsn string, repotype string) repository.Profile
}
