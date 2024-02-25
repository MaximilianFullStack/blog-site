package utils

import (
	"strconv"

	"github.com/a-h/templ"
)

func GetArticleUrl(articleId int) templ.SafeURL {
	return templ.SafeURL("/article?id=" + strconv.Itoa(articleId))
}
