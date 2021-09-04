/*
 * yamlをパースするサンプル。"gopkg.in/yaml.v3"を使っています。
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	var cfgFilePath string
	flag.StringVar(&cfgFilePath, "c", "./config.yaml", "the configuration file path")
	flag.Parse()

	cfgFile, err := ioutil.ReadFile(cfgFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// os.ExpandEnv で環境変数を置き換える
	expaned := os.ExpandEnv(string(cfgFile))

	var cfg Config
	if err := yaml.Unmarshal([]byte(expaned), &cfg); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", cfg)
}

type Config struct {
	DB     DBConfig `yaml:"db"`
	APIKey string   `yaml:"apikey"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
