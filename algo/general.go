package main

import(
"fmt"
"strconv"

)

func mul_strings(num1 string,num2 string) int {

n1,err:=strconv.Atoi(num1)
n2,err2:=strconv.Atoi(num2)

if (err!=nil) || (err2!=nil) {
fmt.Printf("Error converting to int %s",err)
return -1;
}
return n1*n2;

}
//used by rotateArrayToRight function below
func reverse(arr []int,start int,end int) {
var temp int
for start < end {
temp = arr[start]
arr[start] = arr[end]
arr[end]=temp
start++
end--
}

}
//rotates 'k' elements to the right
func rotArrayToRight(arr []int,k int)  {
k=k%len(arr)
reverse(arr,0,len(arr)-1)
reverse(arr,0,k-1)
reverse(arr,k,len(arr)-1)

}
func main(){
arr_slice:=[]int{2,3,5,7,11,13}

rotArrayToRight(arr_slice,3);
fmt.Println(arr_slice)


}
