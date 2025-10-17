//go:generate swag init
package main

import "github.com/intraware/rodan-admin/cmd"

func main() {
	cmd.Run()
}