package console

import "fmt"

func Usage() {
	fmt.Println("Usage:")
	fmt.Println("  " + DockerImages + " | " + App)
	fmt.Println("  " + DockerPs + " [-a] | " + App)
	fmt.Println("  " + DockerComposePs + " [-a] | " + App)
}
