package command

import (
	"fmt"
	"os"
	"path"
	"strings"
	"net/http"

	"github.com/jeremitraverse/golb/config"
	"github.com/jeremitraverse/golb/lexer"
	"github.com/jeremitraverse/golb/line"
	"github.com/jeremitraverse/golb/parser"
	"github.com/jeremitraverse/golb/utils"
)

func Run(args []string) {
	if len(args) == 1 {
		utils.Print_Error("missing operand.")
		return
	}

	switch cmd := args[1]; cmd {
	case "--help":
		fmt.Println("Usage: golb [PATH TO YOUR BLOG]")
		fmt.Println()
		fmt.Println("Full documentation <https://www.github.com/jeremitraverse/golb>")
	case "--build":
		config.GetPosts()
		build()
	case "--init":
		if len(args) == 2 {
			utils.Print_Error("missing blog name.")
			return
		}
		initBlog(args[2])
	case "--serve":
		serve()	
	default:
		utils.Print_Error("command not recognized.")
	}
}

func initBlog(blogName string) {
	workingDir, err := os.Getwd()
	utils.Check(err)

	blogPath := path.Join(workingDir, blogName)
	utils.CreateDir(blogPath)

	config.CreateConfigFile(path.Join(blogPath, "config.json"))

	publicDirPath := path.Join(blogPath, "public")
	utils.CreateDir(publicDirPath)

	distDirPath := path.Join(publicDirPath, "dist")
	utils.CreateDir(distDirPath)

	imageDirPath := path.Join(blogPath, "images")
	utils.CreateDir(imageDirPath)

	postsDirPath := path.Join(blogPath, "posts")
	utils.CreateDir(postsDirPath)
}

func build() {
	var parsedPostsPath []string
	var parsedPostTitles []string

	workingDir, err := os.Getwd()
	utils.Check(err)

	distDirPath := path.Join(workingDir, "public", "dist")
	postsDirPath := path.Join(workingDir, "posts")

	postFiles, err := os.ReadDir(postsDirPath)
	utils.Check(err)
	
	for _, postFile := range postFiles {
		postPath := path.Join(postsDirPath, postFile.Name())

		postContent, err := os.ReadFile(postPath)
		utils.Check(err)
	
		preProcessMarkdownPosts(postContent)

		parsedPostContent, postTitle := parsePost(string(postContent))
		parsedPostPath := path.Join(distDirPath, postFile.Name())

		parsedPostPath = utils.CreateParsedPost(parsedPostPath, parsedPostContent)
		
		parsedPostsPath = append(parsedPostsPath, parsedPostPath)
		parsedPostTitles = append(parsedPostTitles, postTitle)
	}

	if len(parsedPostsPath) > 0 {
		config.UpdateConfigPosts(&parsedPostsPath, &parsedPostTitles)
		utils.UpdateIndexFile(path.Join(workingDir, "public", "index.html"))
	}
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
