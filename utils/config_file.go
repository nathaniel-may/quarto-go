package utils

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

// reflects exact structure for marshalling
type raw struct {
	envs []env
}

type env struct {
	name string
	dBConnString string
}

type config struct {
	envs map[string]env
}

func Load() Config {
	fileContents, err := ioutil.ReadFile("../config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var file raw
	err = json.Unmarshal(fileContents, &file)
	if err != nil {
		fmt.Println(err)
	}

	//TODO DEL
	fmt.Println(file)

	var envMap = make(map[string]env)
	for _, v := range file.envs {
		envMap[v.name] = v
	}

	return &config{envMap}
}

type Config interface {
	GetDBConnString(env string) string
}

func (c *config) GetDBConnString(env string) string {
	return c.envs[env].dBConnString
}