package config

import (
	"encoding/json"
	"os"
	"path"
	"time"

	 "github.com/jeremitraverse/golb/util/error"
)

type BlogConfig struct {
	Author      string
	Description string
	Title       string
	Posts       []Post
}

type Post struct {
	Title       string
	Path        string
	CreatedOn   string
	Description string
}

func CreateConfigFile(path string) {
	f, err := os.Create(path)
	error.Check(err)

	marshaledConfig, marshErr := json.MarshalIndent(BlogConfig{}, " ", " ")
	error.Check(marshErr)

	f.Write(marshaledConfig)
	f.Close()
}

func getConfigFilePath() string {
	workingDir, err := os.Getwd()
	error.Check(err)

	return path.Join(workingDir, "config.json")
}

func GetPosts() *[]Post {
	var config BlogConfig
	configPath := getConfigFilePath()

	f, err := os.ReadFile(configPath)
	error.Check(err)

	err = json.Unmarshal(f, &config)

	return &config.Posts
}

func GetConfig() *BlogConfig {
	var config BlogConfig
	configPath := getConfigFilePath()

	f, err := os.ReadFile(configPath)
	error.Check(err)

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
				Path:        url,
				Title:       titles[index],
				CreatedOn:   time.Now().Format("2006-01"), // Format date to MM-YYYY
				Description: "",
			}

			config.Posts = append(config.Posts, post)
		}
	}

	json, err := json.MarshalIndent(config, " ", " ")
	error.Check(err)

	f, fileOpenErr := os.OpenFile(configFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	error.Check(fileOpenErr)

	f.Write(json)
	f.Close()
}

func postExists(postTitle string, posts []Post) bool {
	for _, post := range posts {
		if post.Path == postTitle {
			return true
		}
	}

	return false
}
