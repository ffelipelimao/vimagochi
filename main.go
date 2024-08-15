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
	// reader := bufio.NewReader(os.Stdin)
	for {
		IdleAnimation()
		/*input, _ := reader.ReadString('\n')
		if input == "0\n" {
			EatAnimation()
		}
		*/
	}
}

func IdleAnimation() {
	for _, frame := range frames.IdleFrame {
		clearScreen()
		fmt.Println(menu.Title)
		fmt.Println(frame)
		fmt.Println(menu.Actions)
		time.Sleep(500 * time.Millisecond)
	}
}

func EatAnimation() {
	for _, frame := range frames.EatFrame {
		clearScreen()
		fmt.Println(menu.Title)
		fmt.Println(frame)
		fmt.Println(menu.Actions)
		time.Sleep(500 * time.Millisecond)
	}
}
