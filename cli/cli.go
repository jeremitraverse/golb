package cli

import (
	"fmt"

	"github.com/jeremitraverse/golb/util/error"
)

func Run(args []string) {
	if len(args) == 1 {
		error.Print_Error("missing operand.")
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
			error.Print_Error("missing blog name.")
			return
		}
		initBlog(args[2])
	case "--serve":
		serve()	
	default:
		error.Print_Error("command not recognized.")
	}
}
