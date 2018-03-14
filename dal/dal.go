package dal

import (
	"quarto-go/quarto"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
	"errors"
	"context"
	"log"
	"fmt"
)

type mongoDal struct {
	client *mongo.Client
	quartoCol *mongo.Collection
}

func NewDal(connString string, database string) (Dal, error) {
	client, err := mongo.NewClient(connString)
	if err != nil {
		return &mongoDal{}, err
	}

	return &mongoDal{client, client.Database(database).Collection("games")}, nil
}

func NilDal() Dal {
	return &mongoDal{}
}

type Dal interface {
	CreateGame() (quarto.Quarto, error)
	LoadGame(boardId string) (quarto.Quarto, error)
	SaveGame(quarto.Quarto) error
}

func (dal *mongoDal) CreateGame() (quarto.Quarto, error) {
	fmt.Println("created called")
	var game quarto.Quarto
	game = quarto.NewBoard(dal.getBoardId())
	fmt.Println("game created not saved")
	fmt.Println("col: ", dal.quartoCol)
	err := dal.SaveGame(game)
	fmt.Println("game created, saved, err not checked")
	if err != nil {
		return quarto.NilBoard(), err
	}

	fmt.Println("created and saved")
	return game, nil
}

func (dal *mongoDal) LoadGame(boardId string) (quarto.Quarto, error) {
	//TODO retry
	result := bson.NewDocument()
	filter := bson.NewDocument(bson.EC.String("boardId", boardId))
	err := dal.quartoCol.FindOne(context.Background(), filter).Decode(result)
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
	game = quarto.NewBoard(dal.getBoardId())

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

	fmt.Println("col in save: ", dal.quartoCol)
	_, err = dal.quartoCol.InsertOne(context.Background(), bsonGame)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (dal *mongoDal) getBoardId() string {
	//TODO randomize
	boardId := "test"

	return boardId
}


