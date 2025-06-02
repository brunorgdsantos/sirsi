package repositories

type CurriculoRepository struct {
	*MongoRepositoryContext
}

func NewCurriculoRepository(uri, dbName, collectionName string) (*CurriculoRepository, error) {
	mongoRepo, err := NewMongoRepositoryContext(uri, dbName, collectionName)
	if err != nil {
		return nil, err
	}
	return &CurriculoRepository{MongoRepositoryContext: mongoRepo}, nil
}
