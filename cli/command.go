package cli

import (
	"fmt"
	"os"
	"path"
	"strings"
	"net/http"

	"github.com/jeremitraverse/golb/util/config"
	"github.com/jeremitraverse/golb/lexer"
	"github.com/jeremitraverse/golb/line"
	"github.com/jeremitraverse/golb/parser"
	"github.com/jeremitraverse/golb/util/error"
	"github.com/jeremitraverse/golb/util/file"
)

func initBlog(blogName string) {
	workingDir, err := os.Getwd()
	error.Check(err)

	blogPath := path.Join(workingDir, blogName)
	file.CreateDir(blogPath)

	config.CreateConfigFile(path.Join(blogPath, "config.json"))

	publicDirPath := path.Join(blogPath, "public")
	file.CreateDir(publicDirPath)
	
	publicStylesDir := path.Join(publicDirPath, "styles")
	file.CreateDir(publicStylesDir)

	file.CreateIndexFile(path.Join(publicDirPath, "index.html"))
	file.CreateStyleFile(path.Join(publicStylesDir, "styles.css"))
	file.CreatePostsStyleFile(path.Join(publicStylesDir, "posts_styles.css"))
	file.CreatePostStyleFile(path.Join(publicStylesDir, "post_styles.css"))

	os.Create(path.Join(publicDirPath, "post_header.html"))
	os.Create(path.Join(publicStylesDir, "post_header_styles.css"))

	distDirPath := path.Join(publicDirPath, "dist")
	file.CreateDir(distDirPath)

	file.CreateParsedPostList(path.Join(distDirPath, "posts.html"))

	imageDirPath := path.Join(blogPath, "images")
	file.CreateDir(imageDirPath)

	postsDirPath := path.Join(blogPath, "posts")
	file.CreateDir(postsDirPath)
}

func build() {
	var parsedPostsPath []string
	var parsedPostTitles []string

	workingDir, err := os.Getwd()
	error.Check(err)

	postsDirPath := path.Join(workingDir, "posts")
	distDirPath := path.Join(workingDir, "public", "dist")
	htmlPostListPath := path.Join(distDirPath, "posts.html")

	postFiles, err := os.ReadDir(postsDirPath)
	error.Check(err)
	
	for _, postFile := range postFiles {
		mdPostPath := path.Join(postsDirPath, postFile.Name())

		mdPostContent, err := os.ReadFile(mdPostPath)
		error.Check(err)
	
		file.PreProcessFileContent(mdPostContent)

		htmlPostContent, postTitle := parsePost(string(mdPostContent))
		htmlPostPath := file.CreateHtmlPost(path.Join(distDirPath, postFile.Name()), htmlPostContent)
		
		parsedPostsPath = append(parsedPostsPath, htmlPostPath)
		parsedPostTitles = append(parsedPostTitles, postTitle)
	}

	if len(parsedPostsPath) > 0 {
		config.UpdateConfigPosts(&parsedPostsPath, &parsedPostTitles)
	}

	var htmlPostList strings.Builder
	
	for _, configPost := range *config.GetPosts() {
		htmlPostList.WriteString(file.FormatConfigPostToHtml(
			configPost.Path,
			configPost.Title,
			configPost.CreatedOn,
			configPost.Description),
		)
	}

	file.WritePostsToIndexHtml(htmlPostListPath, htmlPostList.String())
}

func parsePost(input string) (string, string) {
	var sb strings.Builder
	var postTitle string

	lex := lexer.New(input)
	li := lex.GetLine()	

	for li.Type != line.EOF  {
		p := parser.New(li)
		
		parsedLine := p.ParseLine()

		// removing the h1 tag
		if postTitle == "" && li.Type == line.FIRST_TITLE {
			postTitle = parsedLine[4:len(parsedLine)-5] 
		}

		parsedLine += string('\n')

		sb.WriteString(parsedLine)
		li = lex.GetLine()
	}

	return sb.String(), postTitle
}


func serve() {
	fmt.Println("golb: serving your blog on http://localhost:3000")
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":3000", nil)
}
