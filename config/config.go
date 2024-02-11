package config

import (
	"encoding/json"
	"os"
	"path"
	"time"

	"github.com/jeremitraverse/golb/utils"
)

type BlogConfig struct {
	Author		string
	Description	string
	Title		string
	Posts			[]Post
}

type Post struct {
	Title		string
	Path		string
	CreatedOn	string	
}

func CreateConfigFile(path string) {
	f, err := os.Create(path)
	utils.Check(err)

	marshaledConfig, marshErr := json.MarshalIndent(BlogConfig{}, " ", " ")
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

func GetConfig() *BlogConfig {
	var config BlogConfig
	configPath := getConfigFilePath()

	f, err := os.ReadFile(configPath)
	utils.Check(err)

	err = json.Unmarshal(f, &config)

	return &config
}

func UpdateConfigPosts(postsUrl, postsTitle *[]string) {
	config := GetConfig()

	titles := *postsTitle	

	configFilePath := getConfigFilePath()

	for index, url := range *postsUrl {
		postTitle := titles[index]
		if !postExists(postTitle, config.Posts) {
			post := Post{
				Path: url,
				Title: titles[index],
				CreatedOn: time.Now().Format(time.UnixDate),
			}

			config.Posts = append(config.Posts, post)
		}
	}

	json, err := json.MarshalIndent(config, " ", " ")
	utils.Check(err)

	f, fileOpenErr := os.OpenFile(configFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	utils.Check(fileOpenErr)
		
	f.Write(json)
	f.Close()
}

func postExists(postTitle string, posts []Post) bool {
	for _, post := range posts {
		if post.Title == postTitle {
			return true
		}
	}

	return false
}
