/* EXERCISE: https://tour.golang.org/concurrency/8 */

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// PREORDER WALK
func Walk(t *tree.Tree, ch chan int) {
	if(t.Left != nil) {
		Walk(t.Left, ch);
	}
	ch <- t.Value
	if(t.Right != nil){
		Walk(t.Right, ch);
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int);
	ch2 := make(chan int);
	
	go Walk(t2, ch1);
	go Walk(t1, ch2);
	
	return <-ch1 == <-ch2
}


func PrintResult(same bool){
	if (same) {
		fmt.Println("SAME");
		return;
	} 
	fmt.Println("NOT SAME");
}


func main() {
	firstTree := tree.New(1);
	seccondTree := tree.New(2);

	// SAME
	PrintResult(Same(firstTree,firstTree))

	// NOT SAME
	PrintResult(Same(firstTree,seccondTree))
}