package service

import (
	"bytes"
	"ccproject/domain"
	"ccproject/store"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type BooksService struct {
	store    store.BooksMongoDBStore
	mainHost string
	mainPort string
}

func NewBooksService(store store.BooksMongoDBStore) (BooksService, error) {
	return BooksService{
		store:    store,
		mainHost: os.Getenv("MAIN_HOST"),
		mainPort: os.Getenv("MAIN_PORT"),
	}, nil
}

func (service *BooksService) GetAll() ([]*domain.Book, error) {
	return service.store.GetAll()
}

func (service *BooksService) Borrow(book *domain.Book) (string, error) {
	userID := strconv.Itoa(book.UserID)
	mainLibraryURL := "http://" + service.mainHost + ":" + service.mainPort + "/" + userID

	resp, err := http.Get(mainLibraryURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return "", err
		}

		if err != nil {
			return "", err
		}

		log.Println(string(body))

		return "", errors.New(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	fmt.Println("Number of books:", string(body))
	return string(body), service.store.Insert(book)
}

func (service *BooksService) Register(user *domain.User) (string, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	mainLibraryURL := "http://" + service.mainHost + ":" + service.mainPort + "/register"

	resp, err := http.Post(mainLibraryURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return "", err
		}

		if err != nil {
			return "", err
		}

		log.Println(string(body))

		return "", errors.New(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}

func (service *BooksService) Return(userId int, isbn string) (string, error) {

	book, err := service.store.GetByUserIdIsbn(userId, isbn)
	if err != nil {
		return "", err
	}

	userID := strconv.Itoa(book.UserID)
	mainLibraryURL := "http://" + service.mainHost + ":" + service.mainPort + "/return/" + userID

	resp, err := http.Get(mainLibraryURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return "", err
		}

		if err != nil {
			return "", err
		}

		log.Println(string(body))

		return "", errors.New(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	return "The book was successfully returned. Total number of books owned by the user is: " + string(body), service.store.DeleteOne(userId, isbn)

}
