package main

import "example.com/basic/package/testlib"

func main() {
	song := testlib.GetMusic("Pl Queen")
	println(song)
}
