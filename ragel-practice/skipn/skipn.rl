package main

import (
  "fmt"
)

%%{
  machine skipn;
  action accept {
    fmt.Println("accept")
  }
  action skip {
    fmt.Printf("skip %d\n", data[p])
    p += int(data[p])
  }
  action reject {
    fmt.Printf("Reject at position %d.\n", p)
  }
  action done {
    fmt.Printf("Done at %d\n", p)
    fbreak;
  }
  skipn := ((0x09 $done) | ((extend-0x03) $skip $err(reject)))* 0x00 @accept;
  write data;
}%%


var cs int
var p int = 0
var pe int
var data []byte
var eof int = 0
func main() {
  data = []byte{
    0x01,0x03,
    0x04,0x03,0x03,0x03,0x03,
    0x02,0x03,0x03,
    0x09,0x00,0x00,
  }
  pe = len(data)
  %% write init;
  %% write exec;
}
