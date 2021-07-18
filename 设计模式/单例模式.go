package 设计模式

import "sync"

type Singlon struct{

}

var singlon *Singlon
var one sync.Once
func (single *Singlon) GetSinglon()*Singlon{
	one.Do(func(){
		singlon = &Singlon{}

	})
	return singlon
}