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

func CreateParsedPost(postPath, content string) string {
	parsedPostPath := changeFileExtToHtml(postPath)

	f, err := os.Create(parsedPostPath)
	Check(err)

	_, fileWriteError := f.WriteString(content)
	Check(fileWriteError)

	f.Close()

	return parsedPostPath 
}

func UpdateIndexFile(indexPath string) {

}

func changeFileExtToHtml(filePath string) string {
	var htmlFilePath string
	htmlExtension := ".html"
	postPathExt := filepath.Ext(filePath)

	// remove post file extension
	if postPathExt != "" {
		dir, fileName := filepath.Split(filePath)
		newPostFileName := fileName[:len(fileName) -len(postPathExt)]
		htmlFilePath = filepath.Join(dir, newPostFileName)
	}
	
	htmlFilePath += htmlExtension

	return htmlFilePath
}

