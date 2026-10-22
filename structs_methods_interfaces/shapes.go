package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}

type Circle struct{
	Radius float64
}