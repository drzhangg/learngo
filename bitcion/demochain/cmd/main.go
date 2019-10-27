package main

import "learngo/bitcion/demochain/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to jacky")
	bc.SendData("Send 1 EOS to Jack")
	bc.Print()
}
