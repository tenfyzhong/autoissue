package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config hexo的配置_config.yml
type Config struct {
	URL         string   `yaml:"url"`
	Owner       string   `yaml:"owner"`
	CommentRepo string   `yaml:"comment_repo"`
	Labels      []string `yaml:"labels"`
}

// NewConfig 从文件加生成Config对象
func NewConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
