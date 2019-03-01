package go_web

import (
	"go-web/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type Config map[string]interface{}

type CommonConf struct {
	Common Common `yaml:"common"`
}

type Common struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func ConfigInit(path string) *Config {
	var Conf = make(Config)
	confFiles, err := utils.Traversing(path)
	Error(err)
	for _, filename := range confFiles {
		filepath := path + filename
		fileInfo := strings.Split(filename, ".")
		name, _ := fileInfo[0], fileInfo[1]
		yamlFile, err := ioutil.ReadFile(filepath)
		Error(err)
		var c = new(CommonConf)
		err = yaml.Unmarshal(yamlFile,c)
		Conf[name] = c
	}
  return &Conf
}
