package client_one

import (
	"desingPatterns/creacionales/singleton1/singleton"
)

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
