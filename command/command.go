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
	if len(args) > 2 {
		utils.Print_Error("too many arguments.")
		return
	}

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
	default:
		utils.Print_Error("command not recognized.")
	}
}

func build() {
	working_dir, err := os.Getwd()
	if err != nil {
		utils.Print_Error("error getting working dir.")
		return
	}

	generated_dir_path := path.Join(working_dir, ".generated")
	_, err = os.Stat(generated_dir_path)

	if err != nil {
		err = os.Mkdir(generated_dir_path, 0777)
		if err != nil {
			utils.Print_Error(err.Error())
			return
		}
	} else {
		os.RemoveAll(generated_dir_path)
	}

	posts_dir_path := path.Join(working_dir, "posts")
	_, err = os.Stat(posts_dir_path)

	if err != nil {
		err = os.Mkdir(posts_dir_path, 0777)
		if err != nil {
			utils.Print_Error(err.Error())
			return
		}
	}

	image_dir_path := path.Join(working_dir, "images")
	_, err = os.Stat(image_dir_path)

	if err != nil {
		err = os.Mkdir(image_dir_path, 0777)
		if err != nil {
			utils.Print_Error(err.Error())
			return
		}
	}

	files, err := os.ReadDir(posts_dir_path)
	check(err)
	for _, file := range files {
		post_path := path.Join(posts_dir_path, file.Name())
		data, err := os.ReadFile(post_path)
		check(err)
		htmlString := parsePost(string(data))
		fmt.Println(htmlString)
	}
}

func check(e error) {
	if e != nil {
		utils.Print_Error(e.Error())
		panic(e)
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
