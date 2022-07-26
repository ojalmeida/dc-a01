package server

import (
	"regexp"
)

var routes []route

func init() {

	routes = append(routes,

		route{
			pattern: regexp.MustCompile(`^/tb01`),
			method:  regexp.MustCompile(`POST|OPTIONS`),
			fn:      insertItemHandler,
		},
	)

}
