package main

type StatusA struct {
}

type StatusB struct {
}

type StatusDefault struct {
}

func (s *StatusDefault) switchStatusToA() {
	s = new(StatusA)
}

func main() {

}
