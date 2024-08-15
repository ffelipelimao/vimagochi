package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/ffelipelimao/vimagochi/frames"
	"github.com/ffelipelimao/vimagochi/menu"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for {
		for _, frame := range frames.IdleFrame {
			clearScreen()
			fmt.Println(menu.Title)
			fmt.Println(frame)
			fmt.Println(menu.Actions)
			time.Sleep(600 * time.Millisecond)
		}
	}
}
