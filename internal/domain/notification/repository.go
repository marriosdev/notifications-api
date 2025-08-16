package notification

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	mongoDb *mongo.Database
}

func NewRepository(mongo *mongo.Database) Repository {
	return Repository{
		mongoDb: mongo,
	}
}

func (r Repository) GetAll() {

}
