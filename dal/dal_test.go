package dal

import (
	"testing"
	"quarto-go/utils"
	"fmt"
)

func (dal *mongoDal) mockGetBoardId() string {
	return "test"
}

var d = NilDal()

func getDal() (Dal, error) {
	if d != NilDal() {
		config := utils.LoadConfig("dev")
		d, err := NewDal(config.GetDBConnString(), config.GetDB())
		if err != nil {
			return NilDal(), err
		}
		return d, nil
	}
	return d, nil
}

func TestNewDal(t *testing.T) {
	storage, err := getDal()
	if err != nil {
		t.Errorf(err.Error())
	}

	//not expected to happen, I just have to use the variable to set it here
	if storage == nil {
		t.Fail()
	}
}

func TestCreatesAndLoadsGame(t *testing.T) {
	fmt.Println("started")
	storage, err := getDal()
	if err != nil {
		t.Errorf(err.Error())
	}
	createdGame, err := storage.CreateGame()
	fmt.Println("before err checker")
	if err != nil {
		fmt.Println("inside err checker")
		t.Errorf(err.Error())
	}

	fmt.Println("created")

	loadedGame, err := storage.LoadGame(createdGame.GetId())
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println("loaded")

	if createdGame.GetId() != loadedGame.GetId() {
		t.Fail()
	}

	fmt.Println("finished")
}