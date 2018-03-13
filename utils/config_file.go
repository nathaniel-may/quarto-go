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
	Db string
}

type config struct {
	envName string
	envs map[string]env
}

func Load(envName string) Config {
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

	return &config{envName, envMap}
}

type Config interface {
	GetDBConnString() string
	GetDB() string
}

func (c *config) GetDBConnString() string {
	return c.envs[c.envName].DBConnString
}

func (c *config) GetDB() string {
	return c.envs[c.envName].Db
}