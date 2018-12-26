package main

type Cat struct {
	Color string
	Name string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name:name,
	}
}

func NewCatByColor(color string) *Cat {
	return &Cat{
		Color:color,
	}
}

type BlackCat struct {
	Cat
}

func NewCat(name string) *Cat{
	return &Cat{
		Name:name,
	}
}

func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}


func main() {
	
}
