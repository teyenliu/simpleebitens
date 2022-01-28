package component

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Square struct {
	bgc   color.RGBA
	h     float64
	w     float64
	x     float64
	y     float64
	angle float64
	stepX float64
	stepY float64
	alive bool
	IsRun bool
	Image *ebiten.Image
	Opts  *ebiten.DrawImageOptions
}

func NewSquare(bgc color.RGBA, h, w int, x, y, step float64) *Square {
	image := ebiten.NewImage(w, h)
	image.Fill(bgc)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x, y)
	angle := float64(rand.Intn(120) + 30)
	return &Square{
		bgc:   bgc,
		h:     float64(h),
		w:     float64(w),
		x:     x,
		y:     y,
		angle: angle,
		stepX: step * math.Cos(angle),
		stepY: -step * math.Sin(angle),
		alive: true,
		Image: image,
		Opts:  opts,
	}
}

func (s *Square) GetCoordinate() string {
	return fmt.Sprintf("X:%f, Y:%f", s.x, s.y)
}

func (s *Square) IsAlive() bool {
	return s.alive
}

func (s *Square) CollisionDetection(w, h float64) {
	//假設移動步伐後的座標
	x, y := s.x+s.stepX, s.y+s.stepY

	//移動步伐
	tx, ty := s.stepX, s.stepY

	//移動步伐需做碰撞測試，遇到邊界要反彈
	if x <= 0 { //碰撞左邊垂直邊 (-stepX,stepY)
		tx = -s.x
		ty = tx * math.Tan(s.angle)
		s.stepX *= -1
	} else if x+s.w >= w { //碰撞右邊垂直邊 (-stepX,stepY)
		tx = 320 - s.w - s.x
		ty = tx * math.Tan(s.angle)
		s.stepX *= -1
	} else if y <= 0 { //碰撞上方水平邊 (stepX,-stepY)
		ty = -s.y
		tx = ty / math.Tan(s.angle)
		s.stepY *= -1
	} else if y+s.h >= h { //碰撞下方水平邊 (stepX,-stepY)
		ty = 240 - s.h - s.y
		tx = ty / math.Tan(s.angle)
		s.stepY *= -1
	}

	//將真正的移動步伐更新到球內
	s.x += tx
	s.y += ty

	//將球的圖形利用移動步伐來推移
	s.Opts.GeoM.Translate(tx, ty)
}
