package store

import (
	"errors"
	"context"
	"ccmainproject/domain"
	"go.mongodb.org/mongo-driver/bson"
	
	"go.mongodb.org/mongo-driver/mongo"
	
)

const (
	DATABASE   = "mainlibrary"
	COLLECTION = "users"
)

type UsersMongoDBStore struct {
	users *mongo.Collection
}

func NewUsersMongoDBStore(client *mongo.Client) *UsersMongoDBStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UsersMongoDBStore{
		users: users,
	}
}


func (store *UsersMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UsersMongoDBStore) GetByJmbg(jmbg string) (*domain.User, error) {
	filter := bson.M{"jmbg": jmbg}
	return store.filterOne(filter)
}

func (store *UsersMongoDBStore) GetByUserId(userId int) (*domain.User, error) {
	filter := bson.M{"userId": userId}
	return store.filterOne(filter)
}

func (store *UsersMongoDBStore) Insert(User *domain.User) error {
	_, err := store.users.InsertOne(context.TODO(), User)
	if err != nil {
		return err
	}
	return nil
}

func (store *UsersMongoDBStore) UpdateBooksNum(user *domain.User) error {
	result, err := store.users.UpdateOne(
		context.TODO(),
		bson.M{"jmbg": user.Jmbg},
		bson.D{
			{"$set", bson.D{{"books", user.BooksNum}}},
		},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("one document should've been updated")
	}
	return nil
}



func (store *UsersMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UsersMongoDBStore) filterOne(filter interface{}) (User *domain.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
	return
}


func decode(cursor *mongo.Cursor) (users []*domain.User, err error) {
	for cursor.Next(context.TODO()) {
		var User domain.User
		err = cursor.Decode(&User)
		if err != nil {
			return
		}
		users = append(users, &User)
	}
	err = cursor.Err()
	return
}

