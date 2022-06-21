package main

import (
	"fmt"

	"github.com/pellizzetti/fsfc-simulator/application/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringJSON, _ := route.ExportJSONPositions()
	fmt.Println(stringJSON[1])
}
