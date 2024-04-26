package config

import (
	"encoding/json"
	"os"
	"path"
	"syscall"
	"time"

	"github.com/jeremitraverse/golb/util/error"
)

type BlogConfig struct {
	Posts []Post
}

type Post struct {
	Title       string
	Path        string
	CreatedOn   string
	Description string
	Id          uint64
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

// Gets the Post array from the config file
func GetPosts() *[]Post {
	var config BlogConfig
	configPath := getConfigFilePath()

	f, err := os.ReadFile(configPath)
	error.Check(err)

	err = json.Unmarshal(f, &config)
	error.Check(err)

	return &config.Posts
}

// Gets a json copy of the BlogConfig file
func GetConfig() *BlogConfig {
	var config BlogConfig
	configPath := getConfigFilePath()

	f, err := os.ReadFile(configPath)
	error.Check(err)

	err = json.Unmarshal(f, &config)
	error.Check(err)

	return &config
}

// Updates the list of Posts in the config file. Updates only if the title of an
// existing post has changed or a new post has been created
func UpdateConfigPosts(htmlPostName, postsTitle *[]string) {
	config := GetConfig()

	workingDir, err := os.Getwd()
	error.Check(err)

	// Directory that contains the html posts
	postsDirPath := path.Join(workingDir, "public", "dist")

	titles := *postsTitle
	configFilePath := getConfigFilePath()

	for index, postName := range *htmlPostName {
		postPath := path.Join(postsDirPath, postName)

		fileInfo, _ := os.Stat(postPath)
		fileInfoSys := fileInfo.Sys()

		// Equivalent of doing the stat <filepath> syscall
		fileStat := fileInfoSys.(*syscall.Stat_t)
		fileIno := fileStat.Ino

		postExists, existingPostConfigIndex := postExists(fileIno, config.Posts)

		if postExists {
			existingPost := config.Posts[existingPostConfigIndex]
			// Check if Post title has changed
			if existingPost.Title != (*postsTitle)[existingPostConfigIndex] {
				config.Posts[existingPostConfigIndex].Title = (*postsTitle)[existingPostConfigIndex]
				config.Posts[existingPostConfigIndex].CreatedOn = time.Now().Format("2006-01")
			}
		} else {
			post := Post{
				// Only need the dist since it's we're onyl serving the public folder
				Path:        postName,
				Title:       titles[index],
				CreatedOn:   time.Now().Format("2006-01"), // Format date to MM-YYYY
				Description: "",
				Id:          fileIno,
			}

			config.Posts = append(config.Posts, post)
		}
	}

	jsonConfig, err := json.MarshalIndent(config, " ", " ")
	error.Check(err)
	f, fileOpenErr := os.Create(configFilePath)

	error.Check(fileOpenErr)

	f.Write(jsonConfig)
	f.Close()
}

// Checks if a post existing by comparing it's inode value to each
// existing post in the config file
func postExists(postId uint64, posts []Post) (bool, int) {
	for index, post := range posts {
		if post.Id == postId {
			return true, index
		}
	}

	return false, 0
}
