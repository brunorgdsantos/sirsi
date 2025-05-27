package repositories

type UserRepository struct {
	*MongoRepositoryContext
}

func NewUserRepository(uri, dbname, collectionName string) (*UserRepository, error) {
	mongoRepo, err := NewMongoRepositoryContext(uri, dbname, collectionName)
	if err != nil {
		return nil, err
	}
	return &UserRepository{MongoRepositoryContext: mongoRepo}, nil
}
