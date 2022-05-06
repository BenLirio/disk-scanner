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
var offset int
var toSkip int

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
    if p + int(chunk.Length) >= len(data) {
      toSkip = int(chunk.Length) - (len(data)-p)
      goto SKIP
    } else {
      p += int(chunk.Length)
    }
  }
  action png_done {
    pngEnd = p+1
    f,err := os.Create("found.png")
    if err != nil {
      log.Fatal("failed to open file")
    }
    defer f.Close()
    f.Write(data[pngStart:pngEnd])
    if verb { fmt.Println("PNG end:", pngEnd) }
    os.Exit(0)
  }

  end_chunk = extend{4} @ chunk_data @ png_done
            ;
  other_chunk = extend{4} @ chunk_data
              ;
  chunk = (extend{4} @ chunk_length)
          ("IEND" end_chunk | extend{4} other_chunk)
    ;
  png = png_magic @ header_magic chunk+;
  main := extend* png;
}%%
%% write data;

var BUF_LEN int = 1000
func main() {
  f, err := os.Open("/Users/ben/Desktop/2.png")
  if err != nil {
    log.Fatal("failed to open file")
  }
  b := make([]byte, BUF_LEN)
  data := make([]byte, BUF_LEN)
  cs := 0
  p := 0
  pe := len(b)
  for {
    _,err = f.Seek(int64(offset), 0)
    toSkip = 0
    n, err := f.Read(b)
    p = 0
    if err == io.EOF {
      fmt.Println("None found")
      os.Exit(0)
    }
    if err != nil {
      log.Fatal("Error reading")
    }
    copy(data,b[:n])
    for i := n; i < BUF_LEN; i++ {
      b[i] = 0
    }
    %% write init;
    %% write exec;
    SKIP:
    offset += n+toSkip
    fmt.Printf("Read: %d\t[Total=%d]\n",n,offset)
  }
}
