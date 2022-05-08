package main
import (
  "io"
  "os"
  "fmt"
)
var BUF_LEN int = 256
var ToSkip uint16 = 0
var Skipping bool = false
var Scanning bool = false
var JPEGStart int
var offset int = 0

const (
    SOI_S int = iota
    BLOCK_S
    SOS_S
    EOI_S
)

func SOI(b byte, s []int) {
  for {
    switch s[1] {
    case 0:
      if b==0xFF {
        s[1] += 1
        JPEGStart = offset
      }
      goto FOR_BREAK
    case 1:
      if b==0xD8 {
        s[1] = 0
        s[0] = BLOCK_S
        goto FOR_BREAK
      } else {
        s[1] = SOI_S
      }
    default: goto FOR_BREAK
    }
  }
  FOR_BREAK:
}

func BLOCK(b byte, s []int) {
  switch s[1] {
  case 0:
    if b==0xFF { s[1] += 1 }
  case 1:
    if b==0xDA {
      s[0] = SOS_S
      s[1] = 0
    } else { s[1] += 1 }
  case 2:
    ToSkip += uint16(uint8(b))<<8
    s[1] += 1
    ToSkip -= 1
  case 3:
    ToSkip += uint16(uint8(b))
    s[0] = BLOCK_S
    s[1] = 0
    ToSkip -= 1
    Skipping = true
  default:
    s[0] = SOI_S
    s[1] = 0
  }
}
func EOI(b byte, s []int) {
}

func SOS(b byte, s []int) {
  switch s[1] {
  case 0: if b==0xFF { s[1] += 1 }
  case 1:
    if b==0xD9 {
      fmt.Printf("[0x%02X-0x%02X] size=%dK\n", JPEGStart,offset+1,(offset+1-JPEGStart)/1000)
      s[1] = 0
      s[0] = SOI_S
    } else { s[1] = 0 }
  default:
  }
}

func Next(b byte, s []int) {
  if Skipping && ToSkip > 0 {
    ToSkip -= 1
  } else {
    Skipping = false
    switch s[0] {
    case SOI_S: SOI(b,s)
    case BLOCK_S: BLOCK(b,s)
    case SOS_S: SOS(b,s)
    case EOI_S: EOI(b,s)
    default:
    }
  }
  offset += 1
}

func Run(r io.Reader) {
  b := make([]byte,BUF_LEN)
  s := make([]int,5)
  for {
    n,err := r.Read(b)
    if err == io.EOF {
      break
    }
    if err != nil {
      fmt.Fprintf(os.Stderr, "Failed to read\n")
      break
    }
    for _,x := range b[:n] {
      Next(x,s)
    }
  }
}

func main() {
  //filename := "/home/ben/data/jpgs/2.jpg"
  filename := "/home/ben/data/A.img"
  f,err := os.Open(filename)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to open %s\n", filename)
  }
  Run(f)
  f.Close()
}
