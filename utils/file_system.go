package utils

import (
	"os"
	"path/filepath"
)

func CreateDir(dirPath string) {
	_, err := os.Stat(dirPath) // os.Stat returns an error if the dir exists
	
	if err != nil {
		dirErr := os.Mkdir(dirPath, 0777)
		Check(dirErr)
	} 
}

func GeneratePost(generatedPostPath, content string) {
	htmlExtension := ".html"
	generatedPostPathExt := filepath.Ext(generatedPostPath)

	if generatedPostPathExt != "" {
		dir, fileName := filepath.Split(generatedPostPath)
		newPostFileName := fileName[:len(fileName) -len(generatedPostPathExt)]
		generatedPostPath = filepath.Join(dir, newPostFileName)
	}

	generatedPostPath = generatedPostPath + htmlExtension

	f, err := os.Create(generatedPostPath)
	Check(err)

	_, fileWriteError := f.WriteString(content)
	Check(fileWriteError)

	f.Close()
}
