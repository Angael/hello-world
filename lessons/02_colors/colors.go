package colors

import "fmt"

type Color interface {
	String() string
	ToHex() string
}

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func NewRGB(r, g, b uint8) RGB {
	return RGB{
		Red:   r,
		Green: g,
		Blue:  b,
	}
}

func (c RGB) String() string {
	return fmt.Sprintf("RGB(%d, %d, %d)", c.Red, c.Green, c.Blue)
}

func (c RGB) ToHex() string {
	return fmt.Sprintf("#%02X%02X%02X", c.Red, c.Green, c.Blue)
}

var (
	Red     = RGB{255, 0, 0}
	Green   = RGB{0, 255, 0}
	Blue    = RGB{0, 0, 255}
	Black   = RGB{0, 0, 0}
	White   = RGB{255, 255, 255}
	Yellow  = RGB{255, 255, 0}
	Cyan    = RGB{0, 255, 255}
	Magenta = RGB{255, 0, 255}
)

func Run() {
	// Use the RGB color struct from the package
	customColor := NewRGB(30, 144, 255) // DodgerBlue
	fmt.Println("Custom color:", customColor)
	fmt.Println("Custom color hex:", customColor.ToHex())

	// Using predefined colors
	fmt.Println("Red:", Red)
	fmt.Println("Green:", Green)
	fmt.Println("Blue:", Blue)
}
