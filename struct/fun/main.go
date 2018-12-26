package Vec2

import "fmt"

type Bag struct {
	items []int
}

func Insert(b *Bag, itemid int) {
	b.items = append(b.items,itemid)
}

func main() {

	bag := new(Bag)
	Insert(bag,1001)
	fmt.Println(bag)
}
