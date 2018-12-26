package Vec2

import "fmt"

type Player struct {
	currPos Vec2	//当前位置
	targetPos Vec2 //目标位置
	speed float32
}

//移动到某个点就是设置目标位置
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

//获取当前的位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

//是否到达
func (p *Player) IsArrived() bool {
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

//逻辑更新
func (p *Player) Update() {

	if !p.IsArrived() {
		dir := p.targetPos.Sub(p.currPos).Normalize()

		newPis := p.currPos.Add(dir.Scale(p.speed))

		p.currPos = newPis
	}
}

//创建新玩家
func NewPlayer(speed float32) *Player {
	return &Player{
		speed:speed,
	}
}

func main() {
	// 实例化玩家对象，并设速度为0.5
	p := NewPlayer(0.5)
	// 让玩家移动到3,1点
	p.MoveTo(Vec2{3, 1})
	// 如果没有到达就一直循环
	for !p.IsArrived() {
		// 更新玩家位置
		p.Update()
		// 打印每次移动后的玩家位置
		fmt.Println(p.Pos())
	}
}
