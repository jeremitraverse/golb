package parser

import (
	"strings"

	"github.com/jeremitraverse/golb/config"
)

// Parse the blog config to update index.html
type ConfigParser struct {
	Config		config.BlogConfig
	IndexPath	string
}

func (cp *ConfigParser) New(blogConfig config.BlogConfig, indexPath string) *ConfigParser {
	return &ConfigParser{ Config: *config.GetConfig(), IndexPath: indexPath }
}

func (cp *ConfigParser) Parse() string {
	var sb strings.Builder

	return sb.String()
}

func (cp *ConfigParser) GetHeader() string {
	var sb strings.Builder

	sb.WriteString("<h1>" + cp.Config.Title + "</h1>")

	return sb.String()
}

/*
func (cp *ConfigParser) GetPosts() string {
	var sb strings.Builder
}
*/
