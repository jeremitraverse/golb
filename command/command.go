package command

import (
	"fmt"
	"os"
	"path"
	"strings"

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
	var parsedPostUrls []string
	var parsedPostTitles []string

	workingDir, err := os.Getwd()
	utils.Check(err)

	distDirPath := path.Join(workingDir, "public", "dist")
	postsDirPath := path.Join(workingDir, "posts")

	posts, err := os.ReadDir(postsDirPath)
	utils.Check(err)
	
	for _, post := range posts {
		postPath := path.Join(postsDirPath, post.Name())

		data, err := os.ReadFile(postPath)
		utils.Check(err)

		preProcessMarkdownPosts(data)

		parsedPost, postTitle := parsePost(string(data))
		parsedPostUrl := path.Join(distDirPath, post.Name())

		utils.GeneratePost(parsedPostUrl, parsedPost)

		parsedPostUrls = append(parsedPostUrls, postPath)
		parsedPostTitles = append(parsedPostTitles, postTitle)
	}
	fmt.Println(parsedPostTitles, parsedPostUrls)
	if len(parsedPostUrls) > 0 {
		config.WritePosts(&parsedPostUrls, &parsedPostTitles)
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
