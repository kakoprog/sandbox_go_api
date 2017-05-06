package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Dynamic map[interface{}]interface{}

func Read(keys ...interface{}) interface{} {
	buf, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	m := make(Dynamic)
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}

	confval := m[keys[0]]
	for _, key := range keys[1:] {
		confval = confval.(Dynamic)[key]
	}

	return confval
}
