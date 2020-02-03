// ex1
package main

import (
	"log"
	"time"

	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/ttaem/puzzle/global"
	"github.com/ttaem/puzzle/scene"
	"github.com/ttaem/puzzle/scenemanager"
)

func update(screen *ebiten.Image) error {
	/* Input Process */
	return nil
}

func main() {
	var err error

	rand.Seed(time.Now().Unix())

	scenemanager.SetScene(&scene.StartScene{})

	err = ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1, "PUZZLE")
	if err != nil {
		log.Fatalf("Run Error: %v", err)
	}
}
