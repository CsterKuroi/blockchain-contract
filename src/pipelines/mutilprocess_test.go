package pipelines

import (
	"testing"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func TestFor(t *testing.T) {
	var a = 10
	for i := 0; i < a; i++ {
		fmt.Print("1")
	}

}

func TestSun(t *testing.T) {
	var a int = 1
	var b int = 2
	sum(a, b)
}

func sum(nums ... interface{}) {
	//fmt.Print(nums, " ")

	//for{
	//	logs.Info("1")
	//}

	//s := []int{1, 2, 3}
	//s1 := s[:1]
	//logs.Info(s1)
	//s2 := s[1:]
	//logs.Info(s2)

	//total := 0
	//for _, num := range nums {
	//	total += num
	//}
	//fmt.Println(total)

	//var ss []string=[]string{"1","2","3","4","5","6","7","8"};
	//index:=0;
	//rear:=append([]string{},ss[index:]...)
	//ss=append(ss[0:0],"inserted")
	//ss=append(ss,rear...)
	//logs.Info(ss)

	var ss []string=[]string{"1","2","3","4","5","6","7","8"};

	var s = []string{"!!"}
	ss = append(s,ss...)
	logs.Info(ss)
}
