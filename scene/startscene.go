package scene

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/ttaem/puzzle/font"
	"github.com/ttaem/puzzle/global"
	"github.com/ttaem/puzzle/scenemanager"
)

type StartScene struct {
	startImg *ebiten.Image
}

func (s *StartScene) StartUp() {
	var err error
	s.startImg, _, err = ebitenutil.NewImageFromFile("images/monalisa.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error %v\n", err)

	}
}

func (s *StartScene) Update(screen *ebiten.Image) error {
	screen.DrawImage(s.startImg, nil)

	size := font.TextWidth(global.StartSceneText, 2)
	font.DrawTextWithShadow(screen, global.StartSceneText,
		global.ScreenWidth/2-size/2, global.ScreenHeight/2, 2, color.Black)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		scenemanager.SetScene(&GameScene{})

	}

	return nil
}
