package command

import (
	"fmt"
	"os"
	"path"
	"strings"

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
	blogPath := path.Join(workingDir, blogName)

	utils.CreateDir(blogPath)
	utils.Check(err)

	publicDirPath := path.Join(blogPath, "public")
	utils.CreateDir(publicDirPath)
	
	utils.CreateConfigFile(path.Join(publicDirPath, "config.json"))

	distDirPath := path.Join(publicDirPath, "dist")
	utils.CreateDir(distDirPath)

	imageDirPath := path.Join(blogPath, "images")
	utils.CreateDir(imageDirPath)

	postsDirPath := path.Join(blogPath, "posts")
	utils.CreateDir(postsDirPath)
}

func build() {
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

		htmlString := parsePost(string(data))
		utils.GeneratePost(path.Join(distDirPath, post.Name()), htmlString)
	}
}

func parsePost(input string) string {
	var sb strings.Builder

	lex := lexer.New(input)
	li := lex.GetLine()	

	for li.Type != line.EOF  {
		p := parser.New(li)

		parsedLine := p.ParseLine()
		parsedLine += string('\n')

		sb.WriteString(parsedLine)
		li = lex.GetLine()
	}

	return sb.String()
}

// Making sure that the last line of the post is an empty line
func preProcessMarkdownPosts(content []byte) []byte {
	if content[len(content)-1] != 10 {
		content = append(content, 10)
	}

	return content
}
