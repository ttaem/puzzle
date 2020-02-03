package scene

import (
	"image"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/ttaem/puzzle/global"
)

type GameScene struct {
	bgImg          *ebiten.Image
	subImg         [global.PuzzleColumns * global.PuzzleRows]*ebiten.Image
	board          [global.PuzzleColumns][global.PuzzleRows]int
	blankX, blankY int
}

func (g *GameScene) StartUp() {
	var err error
	g.bgImg, _, err = ebitenutil.NewImageFromFile("images/monalisa.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error %v\n", err)

	}

	width := global.ScreenWidth / global.PuzzleColumns
	height := global.ScreenHeight / global.PuzzleRows

	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; j < global.PuzzleRows; j++ {
			g.subImg[global.PuzzleColumns*j+i] = g.bgImg.SubImage(image.Rect(i*width, j*height, i*width+width, j*height+height)).(*ebiten.Image)

		}
	}

	arr := make([]int, global.PuzzleColumns*global.PuzzleRows)

	var idx int = 0
	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; j < global.PuzzleRows; j++ {
			arr[global.PuzzleColumns*j+i] = idx
			idx++
		}
	}

	g.blankX = global.PuzzleColumns - 1
	g.blankY = global.PuzzleRows - 1

	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; j < global.PuzzleRows; j++ {
			if g.blankX == i && g.blankY == j {
				g.board[i][j] = -1
			} else {
				idx = rand.Intn(len(arr) - 1)
				g.board[i][j] = arr[idx]
				arr = append(arr[:idx], arr[idx+1:]...)
			}

		}
	}

}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (g *GameScene) Update(screen *ebiten.Image) error {

	width := global.ScreenWidth / global.PuzzleColumns
	height := global.ScreenHeight / global.PuzzleRows

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i := x / width
		j := y / height

		if abs(i-g.blankX) == 1 && j == g.blankY {
			g.board[g.blankX][g.blankY] = g.board[i][g.blankY]
			g.board[i][g.blankY] = -1
			g.blankX = i
		} else if abs(j-g.blankY) == 1 && i == g.blankX {
			g.board[g.blankX][g.blankY] = g.board[g.blankX][j]
			g.board[g.blankX][j] = -1
			g.blankY = j
		}

	}

	for i := 0; i < global.PuzzleColumns; i++ {
		for j := 0; j < global.PuzzleRows; j++ {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(i*width), float64(j*height))
			if g.board[i][j] != -1 {
				screen.DrawImage(g.subImg[g.board[i][j]], opts)
			}

		}
	}
	return nil
}
