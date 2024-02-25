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
	"github.com/jeremitraverse/golb/util"
)

func Run(args []string) {
	if len(args) == 1 {
		util.Print_Error("missing operand.")
		return
	}

	switch cmd := args[1]; cmd {
	case "--help":
		fmt.Println("Usage: golb [PATH TO YOUR BLOG]")
		fmt.Println()
		fmt.Println("Full documentation <https://www.github.com/jeremitraverse/golb>")
	case "--build":
		build()
	case "--init":
		if len(args) == 2 {
			util.Print_Error("missing blog name.")
			return
		}
		initBlog(args[2])
	case "--serve":
		serve()	
	default:
		util.Print_Error("command not recognized.")
	}
}

func initBlog(blogName string) {
	workingDir, err := os.Getwd()
	util.Check(err)

	blogPath := path.Join(workingDir, blogName)
	util.CreateDir(blogPath)

	config.CreateConfigFile(path.Join(blogPath, "config.json"))

	publicDirPath := path.Join(blogPath, "public")
	util.CreateDir(publicDirPath)
	
	publicStylesDir := path.Join(publicDirPath, "styles")
	util.CreateDir(publicStylesDir)

	util.CreateIndexFile(path.Join(publicDirPath, "index.html"))
	util.CreateStyleFile(path.Join(publicStylesDir, "styles.css"))
	util.CreatePostsStyleFile(path.Join(publicStylesDir, "posts_styles.css"))
	util.CreatePostStyleFile(path.Join(publicStylesDir, "post_styles.css"))

	os.Create(path.Join(publicDirPath, "post_header.html"))
	os.Create(path.Join(publicStylesDir, "post_header_styles.css"))

	distDirPath := path.Join(publicDirPath, "dist")
	util.CreateDir(distDirPath)

	util.CreateParsedPostList(path.Join(distDirPath, "posts.html"))

	imageDirPath := path.Join(blogPath, "images")
	util.CreateDir(imageDirPath)

	postsDirPath := path.Join(blogPath, "posts")
	util.CreateDir(postsDirPath)
}

func build() {
	var parsedPostsPath []string
	var parsedPostTitles []string

	workingDir, err := os.Getwd()
	util.Check(err)

	postsDirPath := path.Join(workingDir, "posts")
	distDirPath := path.Join(workingDir, "public", "dist")
	htmlPostListPath := path.Join(distDirPath, "posts.html")

	postFiles, err := os.ReadDir(postsDirPath)
	util.Check(err)
	
	for _, postFile := range postFiles {
		mdPostPath := path.Join(postsDirPath, postFile.Name())

		mdPostContent, err := os.ReadFile(mdPostPath)
		util.Check(err)
	
		preProcessMarkdownPosts(mdPostContent)

		htmlPostContent, postTitle := parsePost(string(mdPostContent))
		htmlPostPath := util.CreateHtmlPost(path.Join(distDirPath, postFile.Name()), htmlPostContent)
		
		parsedPostsPath = append(parsedPostsPath, htmlPostPath)
		parsedPostTitles = append(parsedPostTitles, postTitle)
	}

	if len(parsedPostsPath) > 0 {
		config.UpdateConfigPosts(&parsedPostsPath, &parsedPostTitles)
	}

	var htmlPostList strings.Builder
	
	for _, configPost := range *config.GetPosts() {
		htmlPostList.WriteString(util.FormatConfigPostToHtml(
			configPost.Path,
			configPost.Title,
			configPost.CreatedOn,
			configPost.Description),
		)
	}

	util.WritePostsToIndexHtml(htmlPostListPath, htmlPostList.String())
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

// Making sure that the last line of the post is an empty line
func preProcessMarkdownPosts(content []byte) []byte {
	if content[len(content)-1] != 10 {
		content = append(content, 10)
	}

	return content
}

func serve() {
	fmt.Println("golb: serving your blog on http://localhost:3000")
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":3000", nil)
}
