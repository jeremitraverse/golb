package utils

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func CreateDir(dirPath string) {
	_, err := os.Stat(dirPath) // os.Stat returns an error if the dir exists
	
	if err != nil {
		dirErr := os.Mkdir(dirPath, 0777)
		Check(dirErr)
	} 
}

func CreateHtmlPost(postPath, content string) string {
	parsedPostPath := changeFileExtToHtml(postPath)
	fmt.Println(postPath, parsedPostPath)
	f, err := os.Create(parsedPostPath)
	Check(err)

	header := getPostHeaderContent()
	
	// appending the post header to the post content
	_, fileWriteError := f.WriteString(header + content)
	Check(fileWriteError)

	f.Close()

	return path.Base(parsedPostPath)
}

func FormatConfigPostToHtml(postPath, postTitle, postDate string) string {
	return `<li>
	<a class="post" href="` + postPath + `">
		<div class="post-title">
			` + postTitle + `
		</div>
		<div class="post-date">
			`+ postDate + `
		</div>
	</a>
</li>
`
}

func CreateParsedPostList(path string) {
	f, err := os.Create(path)
	Check(err)
	f.Close()
}

func WritePostsToIndexHtml(htmlPostListPath, htmlPostList string) {
	f, err := os.Create(htmlPostListPath)
	Check(err)

	fileContent := `<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="../posts_styles.css">
		<meta charset="utf-8">
		<title></title>
	</head>
	<body>
		<ul class="post-list">
		` + htmlPostList + `
		</ul>
	</body>
</html>`

	_, fileWriteError := f.WriteString(fileContent)
	Check(fileWriteError)

	f.Close()
}

func CreateIndexFile(indexPath string) {
	htmlIndexPath := changeFileExtToHtml(indexPath)

	f, err := os.Create(htmlIndexPath)
	Check(err)

	fileContent := `<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="styles.css">
		<meta charset="utf-8">
		<title></title>
	</head>
	<body>
		<div class=content>
			yo!
			<iframe class="posts" src="./dist/posts.html" width="1200"></iframe>
		</div>
	</body>
</html>
<script>
	// Overriding anchor click within the iframe
    function handleAnchorClick(event) {
        event.preventDefault();

        var href = event.target.parentNode.href;
        window.parent.location.href = href;
    }

    var iframe = document.getElementById('myFrame');
    iframe.onload = function() {
        var iframeContent = iframe.contentWindow;
        var anchors = iframeContent.document.getElementsByTagName('a');
        for (var i = 0; i < anchors.length; i++) {
            anchors[i].addEventListener('click', handleAnchorClick);
        }
    };
</script>`
	
	_, fileWriteError := f.WriteString(fileContent)
	Check(fileWriteError)
}

func CreateStyleFile(stylesheetPath string) {
	f, err := os.Create(stylesheetPath)
	Check(err)
	
	fileContent := `.content {
	display: flex;
	justify-content: center;
	align-items: center;
	flex-direction: column;
}

.content .posts {
	border: none;
	margin-top: 2rem;
}`

	 f.WriteString(fileContent)
}

func CreatePostsStyleFile(stylesheetPath string) {

	f, err := os.Create(stylesheetPath)
	Check(err)
	
	fileContent := `.post-list {
	list-style: none;
}

.post-list .post {
		text-decoration: none;
		color: black;
		margin-bottom: 1rem;
		display: flex;
}

.post-list .post .post-title {
		color: blue;
		margin-left: 1rem;
}`

	f.WriteString(fileContent)
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

func getPostHeaderContent() string {
	baseDir, err := os.Getwd()
	Check(err)

	postHeaderContent, err := os.ReadFile(path.Join(baseDir, "public", "post_header.html"))
	Check(err)

	return string(postHeaderContent)
}

