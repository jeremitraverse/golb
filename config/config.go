package config

import (
	"encoding/json"
	"os"
	"path"

	"github.com/jeremitraverse/golb/utils"
)

type BlogConfig struct {
	BlogAuthor		string
	BlogDescription	string
	BlogTitle		string
	Posts			[]Post
}

type Post struct {
	PostTitle	string
	PostPath	string
}

func CreateConfigFile(path string) {
	f, err := os.Create(path)
	utils.Check(err)

	marshaledConfig, marshErr := json.MarshalIndent(BlogConfig{}, "", "")
	utils.Check(marshErr)

	f.Write(marshaledConfig)
	f.Close()
}

func getConfigFilePath() string {
	workingDir, err := os.Getwd()
	utils.Check(err)

	return path.Join(workingDir, "config.json")
}

func GetPosts() *[]Post {
	var config BlogConfig
	configPath := getConfigFilePath()

	f, err := os.ReadFile(configPath)
	utils.Check(err)

	err = json.Unmarshal(f, &config)

	return &config.Posts
}

func getConfig() *BlogConfig {
	var config BlogConfig
	configPath := getConfigFilePath()

	f, err := os.ReadFile(configPath)
	utils.Check(err)

	err = json.Unmarshal(f, &config)

	return &config
}

func WritePosts(postsUrl, postsTitle *[]string) {
	config := getConfig()

	titles := *postsTitle	

	for index, url := range *postsUrl {
		post := Post{
			PostPath: url,
			PostTitle: titles[index],
		}
		config.Posts = append(config.Posts, post)
	}
	
	configFilePath := getConfigFilePath()

	json, err := json.MarshalIndent(config, "", "")
	utils.Check(err)

	f, fileOpenErr := os.OpenFile(configFilePath, os.O_WRONLY, 0644)
	utils.Check(fileOpenErr)
		
	f.Write(json)
	f.Close()
}
