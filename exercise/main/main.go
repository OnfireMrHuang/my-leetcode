package main

import "fmt"

var count int 

func totalNQueens(n int) int {
	putQueen(n,0,0,0,0)
	return count
}

func putQueen(n,row,col,pia,na int) {
	if row >= n {
		count++
		return
	}

	// 查看是否有空位
	bits := (^(col | pia | na)) & ((1<<n)-1)
	for bits > 0 {
		// 取出最近的一个空位
		p := bits & -bits
		putQueen(n,row+1,(col|p),((pia|p)<<1),((na|p)>>1))
		bits &= bits-1
	}
	return
}

func main() {
	res := totalNQueens(4)
	fmt.Println(res)
}