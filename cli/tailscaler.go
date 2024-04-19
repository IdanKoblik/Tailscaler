package main

import (
	"fmt"
	"log"
	"tailscaler/client"
	"tailscaler/utils"
)

/*
$$$$$$$$\  $$$$$$\  $$$$$$\ $$\       $$$$$$\   $$$$$$\   $$$$$$\  $$\       $$$$$$$$\ $$$$$$$\
\__$$  __|$$  __$$\ \_$$  _|$$ |     $$  __$$\ $$  __$$\ $$  __$$\ $$ |      $$  _____|$$  __$$\
   $$ |   $$ /  $$ |  $$ |  $$ |     $$ /  \__|$$ /  \__|$$ /  $$ |$$ |      $$ |      $$ |  $$ |
   $$ |   $$$$$$$$ |  $$ |  $$ |     \$$$$$$\  $$ |      $$$$$$$$ |$$ |      $$$$$\    $$$$$$$  |
   $$ |   $$  __$$ |  $$ |  $$ |      \____$$\ $$ |      $$  __$$ |$$ |      $$  __|   $$  __$$<
   $$ |   $$ |  $$ |  $$ |  $$ |     $$\   $$ |$$ |  $$\ $$ |  $$ |$$ |      $$ |      $$ |  $$ |
   $$ |   $$ |  $$ |$$$$$$\ $$$$$$$$\\$$$$$$  |\$$$$$$  |$$ |  $$ |$$$$$$$$\ $$$$$$$$\ $$ |  $$ |
   \__|   \__|  \__|\______|\________|\______/  \______/ \__|  \__|\________|\________|\__|  \__|
*/

const HEADER = "$$$$$$$$\\  $$$$$$\\  $$$$$$\\ $$\\       $$$$$$\\   $$$$$$\\   $$$$$$\\  $$\\       $$$$$$$$\\ $$$$$$$\\\n\\__$$  __|$$  __$$\\ \\_$$  _|$$ |     $$  __$$\\ $$  __$$\\ $$  __$$\\ $$ |      $$  _____|$$  __$$\\\n   $$ |   $$ /  $$ |  $$ |  $$ |     $$ /  \\__|$$ /  \\__|$$ /  $$ |$$ |      $$ |      $$ |  $$ |\n   $$ |   $$$$$$$$ |  $$ |  $$ |     \\$$$$$$\\  $$ |      $$$$$$$$ |$$ |      $$$$$\\    $$$$$$$  |\n   $$ |   $$  __$$ |  $$ |  $$ |      \\____$$\\ $$ |      $$  __$$ |$$ |      $$  __|   $$  __$$<\n   $$ |   $$ |  $$ |  $$ |  $$ |     $$\\   $$ |$$ |  $$\\ $$ |  $$ |$$ |      $$ |      $$ |  $$ |\n   $$ |   $$ |  $$ |$$$$$$\\ $$$$$$$$\\\\$$$$$$  |\\$$$$$$  |$$ |  $$ |$$$$$$$$\\ $$$$$$$$\\ $$ |  $$ |\n   \\__|   \\__|  \\__|\\______|\\________|\\______/  \\______/ \\__|  \\__|\\________|\\________|\\__|  \\__|"

func main() {
	fmt.Println(HEADER)
	fmt.Println("Welcome to Tailscaler please select an option: ")
	fmt.Println("1) Get users")
	fmt.Println("2) Lookup user")
	fmt.Println("3) EXIT")

	var option int
	_, _ = fmt.Scanln(&option)

	switch option {
	case 1:
		nodes, err := client.GetNodes()
		if err != nil {
			log.Fatalf("Error getting nodes: %v\n", err)
			return
		}

		utils.PrintTable(nodes)
		break
	case 2:
		fmt.Println("Please enter hostname to lookup: ")
		var hostName string
		_, _ = fmt.Scanln(&hostName)

		node, err := client.LookupNode(hostName)
		if err != nil {
			log.Fatalf("Error getting node: %v\n", err)
			return
		}

		utils.PrintTable(node)
	case 3:
		break
	}
}
