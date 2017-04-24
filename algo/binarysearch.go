package main

import (
"fmt"
)
func binarySearch(arr [5]int,l int,r int,x int) int {
if(r>=l){
mid:=l+(r-l)/2;
if (arr[mid]==x) {
  return mid
}

if(arr[mid]>x){
return binarySearch(arr,l,mid-1,x)
} else {
return binarySearch(arr,mid+1,r,x)
}

}

return -1
}

func main(){


arr:=[5]int{3,5,8,12,14}
fmt.Println(arr)
x:=12
n:=len(arr)
var res int
res=binarySearch(arr,0,n-1,x)

if (res == -1) {
fmt.Println("Element not found")
} else {
fmt.Printf("Element found at index %d",res)
}



}
