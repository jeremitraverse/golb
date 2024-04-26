package file

import (
	"github.com/jeremitraverse/golb/util/error"
	"os"
	"path"
	"path/filepath"
)

func CreateDir(dirPath string) {
	_, err := os.Stat(dirPath) // os.Stat returns an error if the dir exists

	if err != nil {
		dirErr := os.Mkdir(dirPath, 0777)
		error.Check(dirErr)
	}
}

func CreateHtmlPost(postPath, content string) string {
	parsedPostPath := changeFileExtToHtml(postPath)

	f, err := os.Create(parsedPostPath)
	error.Check(err)

	headerContent := getPostHeaderContent()
	headerContent = tabulateContent(headerContent, 2)
	
	bodyContent := tabulateContent([]byte(content), 3)

	postContent := `<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="../styles/post_styles.css">
		<meta charset="utf-8">
	</head>
	<body>
` + string(headerContent) + `
		<div class="post-content">
` + string(bodyContent) + `
		</div>
	</body>
</html>`

	_, fileWriteError := f.WriteString(postContent)
	error.Check(fileWriteError)

	f.Close()

	return path.Base(parsedPostPath)
}

func FormatConfigPostToHtml(postPath, postTitle, postDate, postDescription string) string {
	return `<li>
			<div class="post-title">
				<a class="post-link" href="` + postPath + `" target="_top">
					` + postTitle + `
				</a>
				<div class="post-date">
					` + postDate + `
				</div>
			</div>
			<div class="post-description">
				` + postDescription + `
			</div>
		</li>
`
}

func CreateParsedPostList(path string) {
	f, err := os.Create(path)
	error.Check(err)
	f.Close()
}

func WritePostsToIndexHtml(htmlPostListPath, htmlPostList string) {
	f, err := os.Create(htmlPostListPath)
	error.Check(err)

	fileContent := `<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="../styles/posts_styles.css">
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
	error.Check(fileWriteError)

	f.Close()
}

func CreateIndexFile(indexPath string) {
	htmlIndexPath := changeFileExtToHtml(indexPath)

	f, err := os.Create(htmlIndexPath)
	error.Check(err)

	fileContent := `<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="./styles/styles.css">
		<meta charset="utf-8">
		<title></title>
	</head>
	<body>
		<div class="main-navbar">
			This is the main header, replace me!
		</div>
		<div class=content>
			<iframe class="posts" src="./dist/posts.html" width="900"></iframe>
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
	error.Check(fileWriteError)
}

func CreateStyleFile(stylesheetPath string) {
	f, err := os.Create(stylesheetPath)
	error.Check(err)

	fileContent := `.content {
	display: flex;
	align-items: center;
	flex-direction: column;
	height: 100%;
}

.main-navbar {
	display: flex;
	justify-content: center;
	align-items: center;
	margin-top: 1.5rem;
}

.main-navbar a {
	color: black;
	text-decoration: none
}

html, body, iframe {
	height: 100%;
}

.content #posts {
	border: none;
	margin-top: 2rem;
}`

	f.WriteString(fileContent)
}

func CreatePostsStyleFile(stylesheetPath string) {
	f, err := os.Create(stylesheetPath)
	error.Check(err)

	fileContent := `.post-list {
	list-style: none;
}

.post-list li {
	margin-bottom: 1rem
}

.post-list li .post-title {
	display: flex;
	align-items: end;
}

.post-list li .post-title .post-date {
	margin-left: 0.7rem;
	font-size: 0.7rem;
	color: gray;
}

.post-list li .post-description {
	margin-left: 2rem;
}
`

	f.WriteString(fileContent)
}

func CreatePostStyleFile(stylesheetPath string) {
	f, err := os.Create(stylesheetPath)
	error.Check(err)

	fileContent := `.post-content {
	max-width: 900px
}

h1 {
	text-align: center;
}

body {
	display: flex;
	justify-content: center;
	align-items: center;
	flex-direction: column;
	margin: 0px;
}

.post-header {
	background-color: #dff;
	width: 100%;
	padding: 0.1em 0.1em 0.2em;
	border-top: 1px solid black;
	border-bottom: 1px solid #8ff;
}
`

	f.WriteString(fileContent)
}

func changeFileExtToHtml(filePath string) string {
	var htmlFilePath string
	htmlExtension := ".html"
	postPathExt := filepath.Ext(filePath)

	// remove post file extension
	if postPathExt != "" {
		dir, fileName := filepath.Split(filePath)
		newPostFileName := fileName[:len(fileName)-len(postPathExt)]
		htmlFilePath = filepath.Join(dir, newPostFileName)
	}

	htmlFilePath += htmlExtension

	return htmlFilePath
}

func getPostHeaderContent() []byte {
	baseDir, err := os.Getwd()
	error.Check(err)

	postHeaderContent, err := os.ReadFile(path.Join(baseDir, "public", "post_header.html"))
	error.Check(err)

	return postHeaderContent
}

func tabulateContent(content []byte, numberOfTab int) []byte {
	var tabArray []byte 

	for i := 0; i < numberOfTab; i++ {
		tabArray = append(tabArray, '\t')
	}

	var finalContent []byte
	var previousReturnLinePos int

	preProcessedContent := PreProcessFileContent(content)

	for i, char := range preProcessedContent {
		if char == '\n' {
			line := content[previousReturnLinePos:i+1]

			tempTabedArray := tabArray
			tempTabedArray = append(tempTabedArray, line...)

			finalContent = append(finalContent, tempTabedArray...)

			previousReturnLinePos = i + 1
		}
	}

	return finalContent
}

// Adds return line at the end of a file content 
func PreProcessFileContent(content []byte) []byte {
	if content[len(content)-1] != 10 {
		content = append(content, 10)
	}

	return content
}
