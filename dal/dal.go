package dal

import (
	"quarto-go/quarto"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
	"errors"
	"context"
	"log"
	"sync"
)

var singletonDal = NilDal()
var singletonNil = mongoDal{}
var once sync.Once

type mongoDal struct {
	client *mongo.Client
	quartoCol *mongo.Collection
}

func NewDal(connString string, database string) (Dal, error) {

	var clientErr error
	once.Do(func() {
		client, err := mongo.NewClient(connString)
		if err != nil {
			clientErr = err
		}
		singletonDal = &mongoDal{client, client.Database(database).Collection("games")}
	})
	if clientErr != nil {
		return NilDal(), clientErr
	}

	return singletonDal, nil
}

func NilDal() Dal {
	return &singletonNil
}

type Dal interface {
	CreateGame() (quarto.Quarto, error)
	LoadGame(boardId string) (quarto.Quarto, error)
	SaveGame(quarto.Quarto) error
}

func CheckNilReceiver(receiver interface{}, fName string) error {
	nilDal := NilDal()
	if receiver == nil || receiver == nilDal {
		return errors.New("create game called with nil receiver")
	}
	return nil
}

func (dal *mongoDal) CreateGame() (quarto.Quarto, error) {
	err := CheckNilReceiver(dal, "CreateGame")
	if err != nil {
		return quarto.NilBoard(), err
	}

	var game quarto.Quarto
	boardId, err := dal.getBoardId()
	if err != nil {
		return quarto.NilBoard(), err
	}
	game = quarto.NewBoard(boardId)
	err = dal.SaveGame(game)
	if err != nil {
		return quarto.NilBoard(), err
	}

	return game, nil
}

func (dal *mongoDal) LoadGame(boardId string) (quarto.Quarto, error) {
	err := CheckNilReceiver(dal, "LoadGame")
	if err != nil {
		return quarto.NilBoard(), err
	}

	//TODO retry
	result := bson.NewDocument()
	filter := bson.NewDocument(bson.EC.String("boardId", boardId))
	err = dal.quartoCol.FindOne(context.Background(), filter).Decode(result)
	if err != nil {
		return quarto.NilBoard(), errors.New("error finding board with boardId " + boardId)
	}

	resultBytes, err := bson.Marshal(result)
	if err != nil {
		return quarto.NilBoard(), err
	}

	var game quarto.Quarto
	err = bson.Unmarshal(resultBytes, &game)
	if err != nil {
		return quarto.NilBoard(), err
	}

	return game, nil
}

// schema is defined here
func (dal *mongoDal) SaveGame(game quarto.Quarto) error {
	err := CheckNilReceiver(dal, "SaveGame")
	if err != nil {
		return err
	}

	boardId, err := dal.getBoardId()
	if err != nil {
		return err
	}
	game = quarto.NewBoard(boardId)

	bsonGameBytes, err := bson.Marshal(game)
	if err != nil {
		log.Fatal(err)
		return err
	}

	bsonGame, err := mongo.TransformDocument(bsonGameBytes)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = dal.quartoCol.InsertOne(context.Background(), bsonGame)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (dal *mongoDal) getBoardId() (string, error){
	err := CheckNilReceiver(dal, "CreateGame")
	if err != nil {
		return "", err
	}

	//TODO randomize
	boardId := "test"

	return boardId, nil
}


