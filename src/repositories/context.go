package repositories

import (
	"Sirsi/src/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepositoryContext struct {
	*mongo.Collection
	client *mongo.Client
}

func NewMongoRepositoryContext(uri, dbName, collectionName string) (*MongoRepositoryContext, error) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, utils.InternalServerError(fmt.Sprintf("Erro ao Conectar com o MongoDB: %v", err))
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, utils.InternalServerError(fmt.Sprintf("Erro ao Pingar o MongoDB: %v", err))
	}

	collection := client.Database(dbName).Collection(collectionName)

	return &MongoRepositoryContext{
		Collection: collection,
		client:     client,
	}, nil

}

func (r *MongoRepositoryContext) Create(contextServer context.Context, document interface{}) error {
	_, err := r.Collection.InsertOne(contextServer, document)
	if err != nil {
		return utils.BadRequestError(fmt.Sprintf("Erro ao inserir documento: %v", err))
	}
	return nil
}

func (r *MongoRepositoryContext) GetAll(ctx context.Context) ([]interface{}, error) {
	cursor, err := r.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, utils.InternalServerError(fmt.Sprintf("Erro ao buscar dados: %v", err))
	}
	defer cursor.Close(ctx)

	var results []interface{}
	for cursor.Next(ctx) {
		var doc map[string]interface{}
		if err := cursor.Decode(&doc); err != nil {
			return nil, utils.InternalServerError(fmt.Sprintf("Erro ao decodificar documento: %v", err))
		}
		results = append(results, doc)
	}
	if err := cursor.Err(); err != nil {
		return nil, utils.InternalServerError(fmt.Sprintf("Erro no cursor: %v", err))
	}
	return results, nil
}

func (r *MongoRepositoryContext) CreateJobsDb(contextServer context.Context, document interface{}) error {
	_, err := r.Collection.InsertOne(contextServer, document)
	if err != nil {
		return utils.BadRequestError(fmt.Sprintf("Erro ao inserir dados: %v", err))
	}
	return nil
}
