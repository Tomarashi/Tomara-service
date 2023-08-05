package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"sort"
	"strings"
)

var methodIndexes map[string]int

func init() {
	methodIndexes = map[string]int{}
	methodIndexes["GET"] = 0
	methodIndexes["POST"] = 1
	methodIndexes["PUT"] = 2
	methodIndexes["HEAD"] = 3
	methodIndexes["DELETE"] = 4
}

func PrintRoutes(engine *gin.Engine) {
	routes := make(gin.RoutesInfo, 0)
	for _, route := range engine.Routes() {
		routes = append(routes, route)
	}
	sort.Slice(routes, func(i, j int) bool {
		iMethIndex := math.MaxInt
		jMethIndex := math.MaxInt
		if _, ok := methodIndexes[routes[i].Method]; ok {
			iMethIndex = methodIndexes[routes[i].Method]
		}
		if _, ok := methodIndexes[routes[j].Method]; ok {
			jMethIndex = methodIndexes[routes[j].Method]
		}
		if iMethIndex != jMethIndex {
			return iMethIndex < jMethIndex
		}
		return strings.Compare(routes[i].Path, routes[j].Path) < 0
	})

	fmt.Println("Routes:")
	for _, route := range routes {
		toRepeat := strings.Repeat(" ", 6-len(route.Method))
		fmt.Printf("[%s%s] %s\n", toRepeat, route.Method, route.Path)
	}
}
