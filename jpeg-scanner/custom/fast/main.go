package main
import (
  "io"
  "os"
  "fmt"
  "log"
)
var jpgBuf []byte
var BUF_LEN int = 1<<20
var ToSkip uint16 = 0
var Skipping bool = false
var JPEGStart int
var offset int = 0
var guid int = 0

func IntToSizeString(v int) string {
  s := ""
  if v>>40 > 0 {
    s = fmt.Sprintf("%dT", v>>40)
    v = v%(1<<40)
  }
  if v>>30 > 0 {
    s = fmt.Sprintf("%s %dG", s, v>>30)
    v = v%(1<<30)
  }
  if v>>20 > 0 {
    s = fmt.Sprintf("%s %dM", s, v>>20)
    v = v%(1<<20)
  }
  if v>>10 > 0 {
    s = fmt.Sprintf("%s %dK", s, v>>10)
    v = v%(1<<10)
  }
  if v > 0 {
    s = fmt.Sprintf("%s %dB", s, v)
  }
  if len(s) == 0 {
    s = "0B"
  }
  return s
}

const (
    SOI_S int = iota
    APP0_S
    BLOCK_S
    SOS_S
    EOI_S
)

func reset(s []int) {
  jpgBuf = []byte{}
  ToSkip = 0
  s[0] = SOI_S
  s[1] = 0
}
func extract(s []int) {
  outName := fmt.Sprintf("found/image%d.jpg", guid)
  guid += 1
  fmt.Printf("Found #%d\n at %s", guid, IntToSizeString(offset))
  f,err := os.Create(outName)
  if err != nil {
    log.Fatal(err)
  }
  f.Write(jpgBuf)
  f.Close()
  reset(s)
}
func setHeader(header int, s []int) {
  s[0] = header
  s[1] = 0
}
func SOI(b byte, s []int) {
  switch s[1] {
  case 0: if b==0xFF { 
    s[1] += 1 
    JPEGStart = offset
  } else { reset(s) }
  case 1: if b==0xD8 { setHeader(APP0_S,s) } else { reset(s) }
  }
}

func BLOCK(b byte, s []int) {
  switch s[1] {
  case 0: if b==0xFF { s[1] += 1 } else { reset(s) }
  case 1: if b==0xDA { setHeader(SOS_S, s) } else if b==0xD9 { extract(s) } else { s[1] += 1 }
  case 2:
    ToSkip = uint16(uint8(b))<<8
    s[1] += 1
  case 3:
    ToSkip += uint16(uint8(b))
    setHeader(BLOCK_S,s)
    Skipping = true
  }
}
func EOI(b byte, s []int) {
}
func SOS(b byte, s []int) {
  switch s[1] {
  case 0: if b==0xFF { s[1] += 1 }
  case 1:
    if b==0xFF {
      // do nothing
    } else if b==0x00 || (b >= 0xD0 && b <= 0xD7) {
      s[1] = 0
    } else if b==0xD9 {
      extract(s)
    } else if b==0xDA {
      setHeader(SOS_S,s)
    } else {
      setHeader(BLOCK_S,s)
      s[1] = 2
    }
  }
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
  case 9: setHeader(BLOCK_S, s)
          Skipping = true
  }
}

func Next(b byte, s []int) {
  jpgBuf = append(jpgBuf,b)
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
  if len(jpgBuf) > 1<<10 { reset(s) }
}

func Find(r io.Reader) {
  b := make([]byte,BUF_LEN)
  s := make([]int,5)
  jpgBuf = []byte{}
  for {
    fmt.Printf("%s [Press Enter]",IntToSizeString(offset))
    fmt.Scanln()
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
  //filename := "/Users/ben/Desktop/scifi.jpeg"
  filename := "/Volumes/4TB Drive/disk.img"
  f,err := os.Open(filename)
  defer f.Close()
  fmt.Printf("Opened %s.\n", filename)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to open %s\n", filename)
  }
  /*
  go func() {
    for {
      time.Sleep(time.Second)
      fmt.Printf("%s\n",IntToSizeString(offset))
    }
  }()
  */
  Find(f)
  f.Close()
}
