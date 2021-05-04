package packages

import (
	"GoBaseExercise/utils"
	"fmt"
	"testing"
)
func TestPackages(t *testing.T){
	fmt.Println(utils.GetFibonacci(10))
	utils.Add(1,2)
}
