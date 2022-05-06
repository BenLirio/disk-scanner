package main

import (
  "io"
  "log"
  "os"
  "fmt"
  "encoding/binary"
)

const (
  FROM_BEGINING int = 0
)
var BUF_LEN int = 256
var verb int = 1
var offset int
var PNG_MAGIC_LEN int = 8
var chunkLen int = 0
var cs int
var p int
var pe int
var data []byte
var toSkip bool = false
var pngStart int
var pngEnd int
var err error
var f *os.File

func logAddr(p int) string {
  return fmt.Sprintf("0x%02X", p)
}
func uint32Val(p int) uint32 {
  return binary.BigEndian.Uint32(data[p-3:p+1])
}
func check(e error) {
  if e != nil { log.Fatal(e) }
}
var uidIdx int = 0
func genPNGFileName() string {
  s := fmt.Sprintf("found%d.png", uidIdx)
  uidIdx += 1
  return s
}

%%{
  machine png_extractor;
  action png_found {
    pngStart = offset+p-PNG_MAGIC_LEN+1
    if verb>1 { fmt.Printf("png found at: %s\n", logAddr(pngStart)) }
  }
  action extract_png {
    pngEnd = offset+p
    if verb>0 {
      fmt.Printf("Valid PNG from %s to %s\n", logAddr(pngStart), logAddr(pngEnd))
    }
    f.Seek(int64(pngStart),FROM_BEGINING)
    buf := make([]byte, pngEnd-pngStart)
    f.Read(buf)
    fout,err := os.Create(genPNGFileName())
    check(err)
    _,err = fout.Write(buf)
    check(err)
    f.Seek(int64(offset),FROM_BEGINING)
  }
  action skip_data {
    if p + chunkLen < BUF_LEN {
      p += chunkLen
    } else {
      toSkip = true
      fbreak;
    }
  }



  png_magic = 0x89 0x50 0x4E 0x47 0x0D 0x0A 0x1A 0x0A;
  chunk_len = extend{4} @{ chunkLen=int(uint32Val(p)) };
  chunk_crc = extend{4};
  chunk_data = '' > skip_data;

  header_chunk_type = "IHDR" @{ if verb>1 { fmt.Printf("IHDR at: %s\n", logAddr(p-3)) } };
  end_chunk_type = "IEND" @{ if verb>1 { fmt.Printf("IEND at: %s\n", logAddr(p-3)) } };
  other_chunk_type = "PLTE"|"IDAT" @{ if verb>1 { fmt.Printf("PLTE|IDAT at: %s\n", logAddr(p-3)) } };


  chunk         = chunk_len other_chunk_type  chunk_data chunk_crc;
  end_chunk     = chunk_len end_chunk_type    chunk_data chunk_crc @extract_png;
  header_chunk  = chunk_len header_chunk_type chunk_data chunk_crc;

  png = png_magic @png_found header_chunk chunk+ end_chunk;
  png_extractor := (extend* png)* extend*;
  write data;
}%%


func main() {
  data = make([]byte, BUF_LEN)
  %% write init;


  offset = 0
  f,err = os.Open("data/image.png")
  check(err)

  for {
    n,err := f.Read(data)
    if err == io.EOF { break }
    check(err)
    p = 0
    pe = n
    %% write exec;
    if toSkip {
      toSkip = false
      _, err = f.Seek(int64(offset + p + chunkLen), 0)
      check(err)
      offset += p + chunkLen
    } else {
      offset += n
    }
  }
}
