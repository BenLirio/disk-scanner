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
var JPEGS [][2]int

const (
    SOI_S int = iota
    APP0_S
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
        s[0] = APP0_S
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
    ToSkip = uint16(uint8(b))<<8
    s[1] += 1
  case 3:
    ToSkip += uint16(uint8(b))
    s[0] = BLOCK_S
    s[1] = 0
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
      JPEGS = append(JPEGS, [2]int{JPEGStart, offset+1-JPEGStart})
      s[1] = 0
      s[0] = SOI_S
    } else { s[1] = 0 }
  default:
  }
}
func reset(s []int) {
  ToSkip = 0
  s[0] = SOI_S
  s[1] = 0
}
func APP0(b byte, s []int) {
  switch s[1] {
  case 0: if b==0xFF { s[1] += 1 } else { reset(s) }
  case 1: if b==0xE0 { s[1] += 1 } else { reset(s) }
  case 2: ToSkip = uint16(uint8(b))<<8
          s[1] += 1
  case 3: ToSkip += uint16(uint8(b))
          s[1] += 1
  case 4: if b==0x4A { s[1] += 1 } else { reset(s) }
  case 5: if b==0x46 { s[1] += 1 } else { reset(s) }
  case 6: if b==0x49 { s[1] += 1 } else { reset(s) }
  case 7: if b==0x46 { s[1] += 1 } else { reset(s) }
  case 8: if b==0x00 { s[1] += 1 } else { reset(s) }
  case 9: s[0] = BLOCK_S
          s[1] = 0
          Skipping = true
  default:
  }
}

func Next(b byte, s []int) {
  if !Skipping || ToSkip <= 0 {
    Skipping = false
    switch s[0] {
    case SOI_S: SOI(b,s)
    case APP0_S: APP0(b,s)
    case BLOCK_S: BLOCK(b,s)
    case SOS_S: SOS(b,s)
    case EOI_S: EOI(b,s)
    default:
    }
  }
  ToSkip -= 1
  offset += 1
}

func Find(r io.Reader) {
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
func Extract(r io.ReadSeeker) {
  for i,JPEG := range JPEGS {
    _,err := r.Seek(int64(JPEG[0]),0)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Failed to Seek\n")
      return
    }
    b := make([]byte, JPEG[1])
    n,err := r.Read(b)
    if err != nil || n != JPEG[1] {
      fmt.Fprintf(os.Stderr, "Failed to Read or did not read all\n")
      return
    }
    filename := fmt.Sprintf("found/image%d.jpg", i)
    f,err := os.Create(filename)
    f.Write(b)
    f.Close()
  }
}

func main() {
  //filename := "/home/ben/data/jpgs/2.jpg"
  filename := "/home/ben/data/A.img"
  f,err := os.Open(filename)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to open %s\n", filename)
  }
  Find(f)
  Extract(f)
  f.Close()
}
