package main

import "fmt"
import "strconv"


// 1 * 2 ....
func execute(statement []string) int {
  
  first := statement[0]
  op := statement[1]
  var second string
  
  
  for s, i := range statement {
    
      var op string
      var first, second int
      var r int
      
      switch s {
      case "+":
        op = s
        r = operate(first, op, statement[i+1])
      case "-":
        op = s
        r = operate(first, op, statement[i+1])
      case "*":
        op = s
        r = operate(first, op, statement[i+1])
      case "/":
        op = s
        r = operate(first, op, statement[i+1])
      else 
        first, _ = strconv.Atoi(10, s)
        
      
        
      
  }
  
  if op == "*" || op == "/" {
    second = statement[2]
  } else {
    nextOp := []string{statement[2], statement[3], statement[4]}
    second = strconv.FormatInt(10, execute(nextOp))
  }
  
  result := operate(first, op, second)
   
   
  return excute(result, )
  
}

func operate(f, op, s string) int {
  
    first, _ := strconv.Atoi(f)
    second, _ := strconv.Atoi(s)  
    fmt.Printf("%s %s %s \n", f, op, s)
    var result int
    switch op {
      case "+":
        result =  first + second
      case "-":
        result  = first - second
      case "*":
        result  = first * second
      case "/":
        result  = first / second
    }  
    
    return result
}


func main(){
  
  a := []string{"4", "/", "2"}
  fmt.Println(execute(a))

  b := []string{"1","*","2","+","3"}
  fmt.Println(execute(b))
  
}

