package main

import (
  "log"
  "fmt"
  "io"
  "os"
  "encoding/binary"
)

var eof int = 0

type Chunk struct {
  Length uint32
  Type uint32
  Data []byte
  CRC uint32
}
var chunk Chunk
var verb bool = true
var png []byte
var pngStart int
var pngEnd int

func uint32Val(data []byte,p int) uint32 {
  return binary.BigEndian.Uint32(data[p-3:p+1])
}
func uint8Val(data []byte,p int) uint8 {
  return data[p]
}
%% machine png;
%% machine chunk;

%%{
  machine scanner;
  action header_magic {
    pngStart = p-7
    if verb { fmt.Println("PNG start:", pngStart) }
  }
  action chunk_length {
    chunk.Length = uint32Val(data,p)
    if verb { fmt.Println("chunk.Length =", chunk.Length) }
  }
  png_magic = 0x89 0x50 0x4E 0x47 0x0D 0x0A 0x1A 0x0A;

  action chunk_data {
    if p + int(chunk.Length) > len(data) {
      log.Fatal("Buffered chunk not implemented")
    }
    p += int(chunk.Length)
  }
  action png_done {
    pngEnd = p+1
    if verb { fmt.Println("PNG end:", pngEnd) }
  }

  end_chunk = extend{4} @ chunk_data @ png_done
            ;
  other_chunk = extend{4} @ chunk_data
              ;
  chunk = (extend{4} @ chunk_length)
          ("IEND" end_chunk | extend{4} other_chunk)
    ;
  png = png_magic @ header_magic chunk+;
  main := png;
}%%

%% write data;

var BUF_LEN int = 1<<20
func main() {
  f, err := os.Open("./sample.png")
  if err != nil {
    log.Fatal("failed to open file")
  }
  b := make([]byte, BUF_LEN)
  n, err := f.Read(b)
  if err == io.EOF {
    os.Exit(0)
  }
  if err != nil {
    log.Fatal("Error reading")
  }
  data := make([]byte, n)
//  data := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
//  0x00,0x00,0x00,0x01,  0x00,0x00,0x00,0x00,  0x99,             0x00,0x00,0x00,0x00,
//  0x00,0x00,0x00,0x02,  0x00,0x00,0x00,0x00,  0x99,0x88,        0x00,0x00,0x00,0x00,
//  0x00,0x00,0x00,0x03,  0x49,0x45,0x4E,0x44,  0x99,0x88,0x77,   0x00,0x00,0x00,0x00,
//}
  copy(data,b)
  cs := 0
  p := 0
  pe := len(data)
  %% write init;
  %% write exec;
}
