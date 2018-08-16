package db

import (
	"crossent/micro/studio/db/lock"
)

type RepositoryFactory interface {
	View() ViewRepository
	Compose() ComposeRepository

}

//var ErrDataIsEncrypted = errors.New("failed to decrypt data that is encrypted")
//var ErrDataIsNotEncrypted = errors.New("failed to decrypt data that is not encrypted")


type repositoryFactory struct {
	conn        Conn
	lockFactory lock.LockFactory
}

func NewRepositoryFactory(conn Conn, lockFactory lock.LockFactory) RepositoryFactory {
	return &repositoryFactory{
		conn:        conn,
		lockFactory: lockFactory,
	}
}

func (factory *repositoryFactory) View() ViewRepository {
	return newViewRepository(factory.conn, factory.lockFactory)
}

func (factory *repositoryFactory) Compose() ComposeRepository {
	return newComposeRepository(factory.conn, factory.lockFactory)
}


