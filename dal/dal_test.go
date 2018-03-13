package dal

import (
	"testing"
	"quarto-go/utils"
)

func (dal *mongoDal) mockGetBoardId() string {
	return "test"
}

var d Dal

func TestNewDal(t *testing.T) {
	config := utils.LoadConfig("dev")
	d, err := NewDal(config.GetDBConnString(), config.GetDB())
	if err != nil {
		t.Errorf(err.Error())
	}

	//not expected to happen, I just have to use the variable to set it here
	if d == nil {
		t.Fail()
	}
}

func TestCreatesAndLoadsGame(t *testing.T) {
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