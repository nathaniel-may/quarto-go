package dal

import (
	"quarto-go/quarto"
	"errors"
	"github.com/mongodb/mongo-go-driver/mongo"
	"context"
	"log"
	"github.com/mongodb/mongo-go-driver/bson"
)

type mongoDal struct {
	client *mongo.Client
	quartoCol *mongo.Collection
}

//makes this function mockable for tests
var getBoardId = mongoDal.getBoardId

func NewDal() (Dal, error) {
	//TODO get from file using aws absolute path from built-in sys var
	client, err := mongo.NewClient("mongodb://localhost:27017,localhost:27018,localhost:27019/?replicaSet=repl0&w=majority")
	if err != nil {
		return &mongoDal{}, err
	}
	return &mongoDal{client, client.Database("quartoTest").Collection("quartoTest")}, nil
}

type Dal interface {
	CreateGame() (quarto.Quarto, error)
	LoadGame(boardId string) (quarto.Quarto, error)
	SaveGame(quarto.Quarto) error
}

func (dal *mongoDal) CreateGame() (quarto.Quarto, error) {
	//TODO randomize
	var boardId string
	boardId = getBoardId(*dal)
	_, err := dal.quartoCol.InsertOne(context.Background(), map[string]string{"boardId": boardId})
	if err != nil {
		log.Fatal(err)
		return quarto.NilBoard(), err
	}

	return quarto.NewBoard(boardId), nil
}

func (dal *mongoDal) LoadGame(boardId string) (quarto.Quarto, error) {
	//TODO retry
	result := bson.NewDocument()
	filter := bson.NewDocument(bson.EC.String("boardId", boardId))
	err := dal.quartoCol.FindOne(context.Background(), filter).Decode(result)
	if err != nil {
		return quarto.NilBoard(), errors.New("error finding board with boardId " + boardId)
	}

	var game quarto.Quarto
	game, err = bsonDocToQuarto(*result)
	if err != nil {
		return quarto.NilBoard(), errors.New("error creating quarto board from document with boardId " + boardId)
	}

	return game, nil
}

// schema is defined here
func (dal *mongoDal) SaveGame(game quarto.Quarto) error {

	//TODO STUB
	return nil
}

func (dal *mongoDal) getBoardId() string {
	//TODO randomize
	boardId := "test"

	return boardId
}

// schema is defined here
func bsonDocToQuarto(doc bson.Document) (quarto.Quarto, error) {

	//TODO STUB
	return quarto.NilBoard(), nil
}

// schema is defined here
func quartoToBsonDoc(game quarto.Quarto) (bson.Document, error) {
	var doc *bson.Document

	var piecesArray []bson.Document
	var pieceDoc bson.Document
	for square, piece := range game.GetSquares() {
		//TODO figure out how to save ints
		pieceDoc = *bson.NewDocument(bson.EC.Int32("H", square.H))
	}


	boardId := bson.EC.String("boardId", game.GetId())
	pieces := bson.EC.Array()
	doc = bson.NewDocument(boardId, pieces)
	doc.Append()
	//TODO STUB

	return *doc, nil
}


