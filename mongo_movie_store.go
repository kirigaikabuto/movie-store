package movie_store

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	collection *mongo.Collection
)

type movieStore struct {
	db *mongo.Database
}

func NewMongoStore(config MongoConfig) (MovieStore, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + config.Host + ":" + config.Port)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(config.Database)
	collection = db.Collection("movies")
	return &movieStore{db: db}, nil
}

func (ms *movieStore) List(count int64) ([]Movie, error) {
	findOptions := options.Find()
	if count != 0 {
		findOptions = findOptions.SetLimit(count)
	}
	var movies []Movie
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var movie Movie
		err := cur.Decode(&movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(context.TODO())
	return movies, nil
}

func (ms *movieStore) Create(movie *Movie) (*Movie, error) {
	movies, err := ms.List(0)
	n := len(movies)
	if n != 0 {
		last_material := movies[n-1]
		movie.Id = last_material.Id + 1
	} else {
		movie.Id = 1
	}
	_, err = collection.InsertOne(context.TODO(), movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
func (ms *movieStore) GetById(id int64) (*Movie, error) {
	filter := bson.D{{"id", id}}
	movie := &Movie{}
	err := collection.FindOne(context.TODO(), filter).Decode(&movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
func (ms *movieStore) Update(movie *MovieUpdate) (*Movie, error) {
	return nil, nil
}
func (ms *movieStore) Delete(id int64) error {
	return nil
}
func (ms *movieStore) GetByName(name string) (*Movie, error) {
	filter := bson.D{{"name", name}}
	movie := &Movie{}
	err := collection.FindOne(context.TODO(), filter).Decode(&movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
