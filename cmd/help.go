package cmd

import "fmt"

func Help() {
	help := `
blomple is a simple blog solution. It supports SSR or SSG for client-side interaction.

Usage:

	blomple <command> [arguments]

Commands:
	
	help		show this guideline
	seed		seed database with some data
	ssg		use SSG (Server Side Generation) for frontend
	ssr		use SSR (Server Side Rendering) for frontend
`

	fmt.Println(help)
}
