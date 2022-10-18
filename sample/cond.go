package main

import "sync"

type transaction struct {
	txnum int
	cond  sync.Cond
}

func main() {

	tx := transaction{txnum: 0}

	tx.cond.L.Lock()
	for {
		tx.cond.Wait()
	}
	tx.cond.L.Unlock()

}
