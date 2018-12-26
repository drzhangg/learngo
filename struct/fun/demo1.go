package Vec2

import "fmt"

//定义属性结构
type Property struct {
	value int
}

//设置属性值
func (p *Property) SetValue(v int) {
	p.value = v
}

func (p *Property) Value() int {
	return p.value
}

func main() {

	//实例化属性
	p:=new(Property)

	//设置值
	p.SetValue(100)

	fmt.Println(p.Value())
}
