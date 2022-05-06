package main

import (
)

%%machine url;
%%machine url_prot;
%%write data;

%%{
  url_prot = 'bar';

  main := 'foo' . url_prot;
}%%

var cs int = 0  //cur_state
var p int = 0   //pos
var pe int = 0  //pos_end
var eof int = 0
var data []byte

func main() {
  %%write init;
  for {
    %%write exec;
    break
  }
}
