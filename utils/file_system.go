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

func CreateParsedPostList(parsedPostListPath string) {
	var sb strings.Builder


	f, err := os.Create(parsedPostListPath)
	Check(err)

	sb.WriteString(`<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="../posts_styles.css">
		<meta charset="utf-8">
		<title></title>
	</head>
	<body>
		<ul class="post-list">
			<li>
				<a class="post" href="/">
					<div class="post-title">
						Premier Post 
					</div>
					<div class="post-date">
						26 February 2024 EST 12:00 
					</div>
				</a>
			</li>
		</ul>
	</body>
</html>`)

	_, fileWriteError := f.WriteString(sb.String())
	Check(fileWriteError)

	f.Close()
	
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
</script>`)
	
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

.content .posts {
	border: none;
	margin-top: 2rem;
}`)

	f.WriteString(sb.String())
}

func CreatePostsStyleFile(stylesheetPath string) {
	var sb strings.Builder

	f, err := os.Create(stylesheetPath)
	Check(err)
	
	sb.WriteString(`.post-list {
	list-style: none;
}

.post-list .post {
		text-decoration: none;
		display: flex;
}

.post-list .post .post-title {
		color: blue;
		margin-left: 1rem;
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

