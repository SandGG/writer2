package main

import (
	"fmt"
	"strconv"
)

type myWrite struct {
	s []byte
	i int
}

type myData struct {
	name string
	age  int
}

func (p *myWrite) Write(b []byte) (n int, err error) {
	p.s = append(p.s, b...)
	n = len(b)
	p.i = n
	return
}

func main() {
	var w = &myWrite{}
	var data = []myData{
		{name: "Sandra Juarez", age: 19},
		{name: "Marco Dias", age: 21},
		{name: "Jorge Alvarez", age: 15},
	}
	for i, v := range data {
		var buf = []byte(v.name)
		var auxAge = []byte(strconv.Itoa(v.age))
		buf = append(buf, ' ')
		buf = append(buf, auxAge...)
		if i < len(data)-1 {
			buf = append(buf, '\n')
		}
		w.Write(buf)
	}
	fmt.Println(string(w.s))
}
