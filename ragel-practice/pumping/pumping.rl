package main

import (
  "fmt"
)

var stack int = 0
%%{
  machine pumping;
  action log {
    fmt.Println("log")
  }
  action push {
    fmt.Println("push")
    stack += 1
  }
  action pop {
    fmt.Println("pop")
    stack -= 1
  }
  A = [a]* $ push
    ;
  B = [b]* $ pop
    ;
  pumping := A . B % log
          ;
}%%

%% write data;

var cs int
var p int
var pe int
var data []byte
var eof int
func main() {
  %% write init;

  data = []byte("aaabbb")
  %% write exec;
  fmt.Println(stack)
}
