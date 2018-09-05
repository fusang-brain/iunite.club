package client

import (
	"strconv"
)

type clientKey int

func (ck clientKey) getKey() string {
	// if ck == 1 {

	// }
	return "iron.service.ck_" + strconv.Itoa(int(ck))
}

const (
	userKey clientKey = iota
)
