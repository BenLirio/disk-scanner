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

type Png struct {
  Height uint32
  Width uint32
  BitDepth uint8
  Chunks []Chunk
}
var png Png
func uint32Val(data []byte,p int) uint32 {
  return binary.BigEndian.Uint32(data[p-3:p+1])
}
func uint8Val(data []byte,p int) uint8 {
  return data[p]
}
%% machine png;
%% machine chunk;
%% machine ihdr_chunk;

%%{
  machine scanner;
  action header_magic {
    png = Png{}
    if verb { fmt.Println("Png magic") }
  }
  action chunk_length {
    chunk.Length = binary.BigEndian.Uint32(data[p-3:p+1])
    chunk.Length = uint32Val(data,p)
    if verb { fmt.Println("chunk.Length =", chunk.Length) }
  }
  action chunk_type {
    chunk.Type = binary.BigEndian.Uint32(data[p-3:p+1])
  }
  action chunk_data {
    if p+int(chunk.Length) < len(data) {
      chunk.Data = make([]byte, chunk.Length)
      copy(chunk.Data, data[p:p+int(chunk.Length)])
      if verb { fmt.Println("chunk.Data =", chunk.Data) }
    } else {
      if verb { fmt.Println("Bad chunk size", chunk.Data) }
    }
  }
  action ihdr_width {
    png.Width = uint32Val(data,p)
    if verb { fmt.Println("Png Width =", png.Width) }
  }
  action ihdr_height {
    png.Height = uint32Val(data,p)
    if verb { fmt.Println("Png Height =", png.Height) }
  }
  action ihdr_bit_depth {
    png.BitDepth = uint8Val(data,p)
    if verb { fmt.Println("Bit Depth =", png.BitDepth) }
  }
  png_magic = 0x89 0x50 0x4E 0x47 0x0D 0x0A 0x1A 0x0A;
  ihdr_chunk = extend{4} @ ihdr_width
               extend{4} @ ihdr_height
               extend{1} @ ihdr_bit_depth
             ;
  chunk = (extend{4} @ chunk_length)
          ("IHDR" ihdr_chunk)
    ;
  png = png_magic @ header_magic chunk+;
  main := png;
}%%

%% write data;

var BUF_LEN int = 256
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
  copy(data,b)
  cs := 0
  p := 0
  pe := len(data)
  %% write init;
  %% write exec;
}
