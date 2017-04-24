package main

import(
"fmt"
"errors"
"log"
)

type Node struct{
Value string
Data string
Left *Node
Right *Node
}
//insert a node into binary search tree
func (n *Node) Insert(value,data string) error {

if n==nil{
return errors.New("cant insert a value into nil tree")
}
switch{

case value==n.Value:
      return nil
case value<n.Value:
   if n.Left==nil{
     n.Left= &Node{Value:value,Data:data}
     return nil
   }
   return n.Left.Insert(value,data)


case value>n.Value:
   if n.Right==nil{
     n.Right=&Node{Value:value,Data:data}
     return nil
   }
   return n.Right.Insert(value,data)
} 
return nil
}


//find if node exists in BST
func (n *Node) Find(s string) (string,bool){
if n==nil{
  return "",false
}
switch{
case n.Value==s:
  return n.Data,true
case n.Value > s:
  return n.Left.Find(s)
default:
 return n.Right.Find(s)

}

}

type Tree struct{
Root *Node
}
//actual BST tree functions

func (t *Tree) Insert(value,data string) error{
 if t.Root==nil{
    t.Root=&Node{Value:value,Data:data}
    return nil
 }
 return t.Root.Insert(value,data)

}

func (t *Tree) Find(s string) (string,bool) {
     if t.Root==nil{
        return "",false
     }
     return t.Root.Find(s)
}

func (t *Tree) printTree(n *Node){
if n==nil{
return
}
t.printTree(n.Left)
//inorder
fmt.Println("tree node: "+n.Value);
t.printTree(n.Right)


}

func main(){

values:= []string{"a","b","c","d","e"}
data := []string{"bc","he","nn","la","va"}

tree:=&Tree{}

for i:=0;i<len(values);i++{
err:=tree.Insert(values[i],data[i])
if err!=nil{
   log.Fatal("Error inserting value: ",values[i],":",err)
}


}

tree.printTree(tree.Root);
s:="d"
fmt.Print("finding node: ",s)
d,found:= tree.Find(s)

if !found{
log.Fatal("cant find element: ",s);
} else {
fmt.Println("Found element - "+s+" : "+d)
}

}
