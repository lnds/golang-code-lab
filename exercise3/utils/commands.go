package utils

import (
	"log"
	"strconv"
	"strings"
)

func ParseCommand(command string) (string, int) {
	p := strings.Split(command, " ")
	if len(p) == 2 {
		arg, err := strconv.Atoi(p[1])
		if err != nil {
			log.Fatal(err)
		}
		return p[0], arg
	} else {
		return p[0], 0
	}
}
