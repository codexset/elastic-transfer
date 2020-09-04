package manage

import (
	"elastic-transfer/app/types"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var manager *ElasticManager

func TestMain(m *testing.M) {
	os.Chdir("../..")
	var err error
	if _, err := os.Stat("./config/autoload"); os.IsNotExist(err) {
		os.Mkdir("./config/autoload", os.ModeDir)
	}
	cfgByte, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Fatalln("Failed to read service configuration file", err)
	}
	config := types.Config{}
	err = yaml.Unmarshal(cfgByte, &config)
	if err != nil {
		log.Fatalln("Service configuration file parsing failed", err)
	}
	manager, err = NewElasticManager(config.Elastic, config.Mq)
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}

func TestElasticManager_Put(t *testing.T) {
	err := manager.Put(types.PipeOption{
		Identity: "task",
		Service:  "schedule",
		Validate: `{"type":"object","properties":{"name":{"type":"string"}}}`,
		Topic:    "sys.schedule",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestElasticManager_Delete(t *testing.T) {
	err := manager.Delete("task")
	if err != nil {
		t.Error(err)
	}
}
