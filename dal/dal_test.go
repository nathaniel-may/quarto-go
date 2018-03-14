package dal

import (
	"testing"
	"quarto-go/utils"
)

var config = utils.LoadConfig("dev")

func (dal *mongoDal) mockGetBoardId() string {
	return "test"
}

func TestNewDal(t *testing.T) {
	d, err := NewDal(config.GetDBConnString(), config.GetDB())
	if err != nil {
		t.Errorf(err.Error())
	}

	if d == NilDal() {
		t.Fail()
	}
}

func TestNewDalTwice(t *testing.T) {
	d, err := NewDal(config.GetDBConnString(), config.GetDB())
	if err != nil {
		t.Errorf(err.Error())
	}

	if d == NilDal() {
		t.Fail()
	}
}

func TestCreatesAndLoadsGame(t *testing.T) {
	d, err := NewDal(config.GetDBConnString(), config.GetDB())
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