/*
 * Revision History:
 *     Initial: 2018/05/02        Chen Yanchen
 */

package main

import (
	"log"
	"sync"
)

type Mod int8

const (
	ModTest Mod = iota // test version
	ModDev             // develop version
	ModCom             // complete version
)

type Lock struct {
	Switch bool
	Model  Mod
	Size   uint32
	wgMu   sync.RWMutex
	wg     sync.WaitGroup
}

func (l *Lock) SetModel(mod Mod) {
	l.Model = mod
}

// SetSize set lock's size, 0 < size < 32
func (l *Lock) SetSize(size uint32) {
	if size > 0 && size < 32 {
		l.Size = size
	}
}

func (l *Lock) New(f func()) {
	l.wgMu.RLock()
	defer l.wgMu.RUnlock()

	select {
	case l.Switch == false:
		log.Println("Unlock")
		return
	default:
	}

	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		f()
	}()
}
