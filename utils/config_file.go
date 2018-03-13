package utils

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
)

// reflects exact structure for marshalling
type raw struct {
	Envs []env
}

type env struct {
	Name string
	DBConnString string
}

type config struct {
	envs map[string]env
}

func Load() Config {
	fileContents, err := ioutil.ReadFile("../config.json")
	if err != nil {
		log.Fatal(err.Error())
		return &config{}
	}

	var file raw
	err = json.Unmarshal(fileContents, &file)
	if err != nil {
		fmt.Println(err)
	}

	var envMap = make(map[string]env)
	for _, v := range file.Envs {
		envMap[v.Name] = v
	}

	return &config{envMap}
}

type Config interface {
	GetDBConnString(env string) string
}

func (c *config) GetDBConnString(env string) string {
	return c.envs[env].DBConnString
}