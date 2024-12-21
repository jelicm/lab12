package store

import (
	"context"
	"ccproject/domain"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"fmt"
)

const (
	DATABASE   = "citylibrary"
	COLLECTION = "books"
)

type BooksMongoDBStore struct {
	books *mongo.Collection
}

func NewBooksMongoDBStore(client *mongo.Client) *BooksMongoDBStore {
	books := client.Database(DATABASE).Collection(COLLECTION)
	return &BooksMongoDBStore{
		books: books,
	}
}


func (store *BooksMongoDBStore) GetAll() ([]*domain.Book, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *BooksMongoDBStore) Insert(Book *domain.Book) error {
	
	_, err := store.books.InsertOne(context.TODO(), Book)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}


func (store *BooksMongoDBStore) GetByUserIdIsbn(userId int, isbn string) (*domain.Book, error) {
	filter := bson.M{"userID": userId, "isbn": isbn}
	return store.filterOne(filter)
}

func (store *BooksMongoDBStore) DeleteOne(userId int, isbn string) error {
	filter := bson.M{"userID": userId, "isbn": isbn}
	_, err := store.books.DeleteOne(context.TODO(), filter)
	return err
}

func (store *BooksMongoDBStore) filterOne(filter interface{}) (Book *domain.Book, err error) {
	result := store.books.FindOne(context.TODO(), filter)
	err = result.Decode(&Book)
	return
}


func (store *BooksMongoDBStore) filter(filter interface{}) ([]*domain.Book, error) {
	cursor, err := store.books.Find(context.TODO(), filter)
	
	if cursor == nil{
		fmt.Println("aaaaa")
		return nil, err
	}
	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}
	return decode(cursor)
}


func decode(cursor *mongo.Cursor) (books []*domain.Book, err error) {
	for cursor.Next(context.TODO()) {
		var Book domain.Book
		err = cursor.Decode(&Book)
		if err != nil {
			return
		}
		books = append(books, &Book)
	}
	err = cursor.Err()
	return
}