package binary_search

import "testing"

func Test_binary_search(t *testing.T){
    input:=[]int{1,2,3,4,5,6,7}
    output:=search_v1(input, 7)
    println(output)
}

func Test_binary_search1(t *testing.T){
    input:=[]int{1,2,3,4,5,6,7}
    output:=search_excess_v1(input, 7)
    println(output)
}

func Test_binary_search2(t *testing.T){
    input:=[]int{1,2,3,4,5,6,7}
    output:=search_v2(input, 7)
    println(output)
}