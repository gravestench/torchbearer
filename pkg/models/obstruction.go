package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Obstruction int

func (o Obstruction) String() string {
	return fmt.Sprintf("Ob%d", o)
}

func (Obstruction) FromString(s string) Obstruction {
	if !strings.HasPrefix(s, "Ob") {
		return Obstruction(0)
	}

	s = strings.ReplaceAll(s, "Ob", "")

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return Obstruction(0)
	}

	return Obstruction(int(v))
}

type Factor string

func (f Factor) Value() {

}

func ParseObstruction(s string) (o Obstruction) {
	return o.FromString(s)
}
