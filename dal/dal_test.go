package tests

import (
	"testing"
	//"quarto-go/quarto"
	"quarto-go/dal"
)

func mockGetBoardId() string {
	return "test"
}

func TestCreatesGame(t *testing.T) {
	d, err := dal.New()
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = d.CreateGame()
	if err == nil {
		t.Errorf(err.Error())
	}
}

//TODO how?
func TestNotCreateGameWithDupBoardId(t *testing.T) {

}