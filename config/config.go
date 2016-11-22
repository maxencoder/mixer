package config

import (
	"io/ioutil"

	"github.com/siddontang/go-yaml/yaml"
)

type NodeConfig struct {
	Name             string `yaml:"name"`
	DownAfterNoAlive int    `yaml:"down_after_noalive"`
	IdleConns        int    `yaml:"idle_conns"`

	User     string `yaml:"user"`
	Password string `yaml:"password"`

	Master string `yaml:"master"`
	Slave  string `yaml:"slave"`
}

type SchemaConfig struct {
	DB          string      `yaml:"db"`
	Nodes       []string    `yaml:"nodes"`
	RulesConifg RulesConfig `yaml:"rules"`
}

type RulesConfig struct {
	Default   string        `yaml:"default"`
	ShardRule []ShardConfig `yaml:"shard"`
}

type ShardConfig struct {
	Table  string       `yaml:"table"`
	Key    string       `yaml:"key"`
	Nodes  []string     `yaml:"nodes"`
	Type   string       `yaml:"type"`
	Range  string       `yaml:"range"`
	Lookup LookupConfig `yaml:"lookup"`
}

type LookupConfig struct {
	Node  string `yaml:node`
	Query string `yaml:query`
}

type Config struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`

	Schemas []SchemaConfig `yaml:"schemas"`

	Nodes []NodeConfig `yaml:"nodes"`
}

func (cfg *Config) Marshal() (data []byte, err error) {
	if data, err = yaml.Marshal(cfg); err != nil {
		return nil, err
	}
	return data, nil
}

func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ParseConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return ParseConfigData(data)
}
