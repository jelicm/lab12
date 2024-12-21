package domain

type User struct {

	UserId  	int                `bson:"userId"`
	Name		string			   `bson:"name"`
	Surname  	string			   `bson:"surname"`
	Address		string			   `bson:"address"`
	Jmbg		string		   	   `bson:"jmbg"`
	BooksNum	int				   `bson:"books"`
}