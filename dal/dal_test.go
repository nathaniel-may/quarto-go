package dal

import (
	"testing"
	//"quarto-go/quarto"
)

func (dal *mongoDal) mockGetBoardId() string {
	return "test"
}

func TestCreatesAndLoadsGame(t *testing.T) {
	d, err := NewDal()
	if err != nil {
		t.Errorf(err.Error())
	}

	createdGame, err := d.CreateGame()
	if err != nil {
		t.Errorf(err.Error())
	}

	loadedGame, err := d.LoadGame(createdGame.GetId())
	if err != nil {
		t.Errorf(err.Error())
	}

	if createdGame.GetId() != loadedGame.GetId() {
		t.Fail()
	}
}

// mongo driver (alpha) has no client.Close() function so I am not creating new dal objects
func TestNotCreateGameWithDupBoardId(t *testing.T) {
	var save func(mongoDal) string
	save = getBoardId
	getBoardId = mongoDal.mockGetBoardId
	d, err := NewDal()
	_, err = d.CreateGame()
	if err == nil {
		t.Errorf(err.Error())
	}
	_, err = d.CreateGame()
	//expect error
	if err != nil {
		t.Errorf(err.Error())
	}

	//TODO does this work?
	getBoardId = save
}