package demo1

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"testing"
)
var bufPool sync.Pool
type Buffer interface {
	//Delimiter 用于数据块之间的定界符
	Delimiter() byte
	//Write 用于写一个数据块
	Write(content string)(err error)
	//Read 用于读一个数据块
	Read() (content string,err error)
	//Free用于释放当前的缓冲区
	Free()
}
type mybuffer struct {
	delimiter byte //定界符
	buf bytes.Buffer
}

func (b *mybuffer) Delimiter() byte{
	return b.delimiter
}

func (b *mybuffer) Write(content string)(err error){
	if _,err := b.buf.WriteString(content); err != nil{
		return err
	}
	return b.buf.WriteByte(b.delimiter)
}

func (b *mybuffer) Read()(content string,err error){
	return b.buf.ReadString(b.delimiter)
}

func (b *mybuffer) Free(){
	bufPool.Put(b)
}

//delimiter 预定义定界符
var delimiter = byte('\n')

func init(){
	bufPool = sync.Pool{
		New: func() interface{} {
			return &mybuffer{delimiter:delimiter}
		},
	}
}

func GetBuffer()Buffer{
	return bufPool.Get().(Buffer)
}

func TestBuffer(t *testing.T){
	buf := GetBuffer()
	defer buf.Free()
	buf.Write("A Pool is a set of temporary objects that" +
		"may be individually saved and retrieved.")
	buf.Write("A Pool is safe for use by multiple goroutines simultaneously.")
	buf.Write("A Pool must not be copied after first use.")

	fmt.Println("The data blocks in buffer:")

	for {
		block, err := buf.Read()
		if err != nil {
			if err == io.EOF{
				break
			}
			panic(fmt.Errorf("unexpected error:%s",err))
		}
		fmt.Println(block)
	}
}