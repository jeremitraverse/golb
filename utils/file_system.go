package utils

import (
	"os"
	"path/filepath"
	"strings"
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

func CreateIndexFile(indexPath string) {
	var sb strings.Builder
	htmlIndexPath := changeFileExtToHtml(indexPath)

	f, err := os.Create(htmlIndexPath)
	Check(err)

	sb.WriteString(`<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="styles.css">
		<meta charset="utf-8">
		<title></title>
	</head>
	<body>
		<div class=content>
			yo!
			<iframe class="post-list" src="./posts.html"></iframe>
		</div>
	</body>
</html>`)
	
	_, fileWriteError := f.WriteString(sb.String())
	Check(fileWriteError)
}

func CreateStyleFile(stylesheetPath string) {
	var sb strings.Builder

	f, err := os.Create(stylesheetPath)
	Check(err)
	
	sb.WriteString(`.content {
	display: flex;
	justify-content: center;
	align-items: center;
	flex-direction: column;
}

.content .post-list {
	border: none;
	margin-top: 2rem;
}`)

	f.WriteString(sb.String())
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

