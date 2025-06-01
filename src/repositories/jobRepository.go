package repositories

type JobRepository struct {
	*MongoRepositoryContext
}

func NewJobRepository(uri, dbName, collectionName string) (*JobRepository, error) {
	mongoRepo, err := NewMongoRepositoryContext(uri, dbName, collectionName)
	if err != nil {
		return nil, err
	}
	return &JobRepository{MongoRepositoryContext: mongoRepo}, nil
}
