package config

import (
	"io/ioutil"
	"log"

	"github.com/hashicorp/raft"
	jsonIter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"

	"github.com/penk110/interview_ext_go/logging_zap"
)

func NewConfig() *Config {
	return &Config{}
}

func LoadConfig(path string) *Config {
	config := NewConfig()
	if path == "" {
		log.Fatal("config path not allow empty.")
	}
	fileBytes := loadConfigFile(path)

	err := yaml.Unmarshal(fileBytes, config)
	if err != nil {
		log.Fatalf("unmarshal config file failed, err: %s", err.Error())
	}

	jsonData, _ := jsonIter.Marshal(config)
	log.Printf("unmarshal config file success, config: %s\n", string(jsonData))
	return config
}

func loadConfigFile(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return b
}

type Config struct {
	Service *Service `json:"service" yaml:"service"`
	Nodes   []*Node  `json:"nodes" yaml:"nodes"`

	CacheType string `json:"cacheType" yaml:"cacheType"`

	Log *logging_zap.LogConfig `yaml:"log"`

	RaftConfig *RaftConfig `json:"raftConfig" yaml:"raftConfig"`
}

type Service struct {
	ServiceID   string `json:"serviceId" yaml:"serviceId"`
	ServiceName string `json:"serviceName" yaml:"serviceName"`
	IP          string `json:"ip" yaml:"ip"`
	Port        string `json:"port" yaml:"port"`
	Mode        string `json:"mode" yaml:"mode"`
}

type Node struct {
	ID      raft.ServerID      `json:"id" yaml:"id"`
	Http    string             `json:"http" yaml:"http"`
	Address raft.ServerAddress `json:"address" yaml:"address"`
}

type RaftConfig struct {
	Transport   string `json:"transport" yaml:"transport"`
	LogStore    string `json:"logStore" yaml:"logStore"`
	LocalCache  string `json:"localCache" yaml:"localCache"`
	StableStore string `json:"stableStore" yaml:"stableStore"`
	Snapshot    string `json:"snapshot" yaml:"snapshot"` //快照保存的位置
}
