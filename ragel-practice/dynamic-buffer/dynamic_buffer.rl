package main
import (
  "os"
  "fmt"
  "bytes"
  "io"
)

var matching bool = false
const (
  BUF_LEN int = 1
  NOT_DONE int = -1
  DONE int = 0
)

%%{
  machine dynamic_buffer;
  action short_match {
    fmt.Println("short_match")
  }
  action begin_short_match {
    fmt.Println("begin_short_match")
  }
  action begin_long_match {
    fmt.Println("long_long_match")
  }
  action long_match {
    fmt.Println("long_match")
  }
  short_dfa = 1 2 3;
  long_dfa = 1 2 3 4 5 6 7;
  dynamic_buffer :=
  |*
  extend;
  short_dfa >begin_short_match @short_match;
  long_dfa >begin_long_match @long_match;
  *|;
  write data;
}%%

type Machine struct {
  ts int
  te int
  act int
  cs int
  data []byte
  offset int
}

func (m *Machine) init() {
  %% access m.;
  %% write init;
}
func NewMachine() *Machine {
  m := &Machine{}
  m.init()
  return m
}

func (m *Machine) exec(p int, pe int, eof int) {
  %% write exec;
}

func (m *Machine) Run(f io.Reader) {
  m.data = make([]byte, BUF_LEN)
  for {
    n,err := f.Read(m.data)
    if err == io.EOF {
      m.exec(0,0,0)
      break
    }
    if err != nil {
      fmt.Fprintf(os.Stderr, "Failed to read")
      break
    }
    m.exec(0,n,NOT_DONE)
    m.offset += n
    fmt.Printf("offset=%d\tts=%d\tte=%d\tact=%d\tdata=%v\n\n",m.offset,m.ts, m.te, m.act, m.data)
  }
}

func main() {
  m := NewMachine()
  /*
  filename := "data/t1"
  f,err := os.Open(filename)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to open %s\n", filename)
    return
  }
  */
  f := bytes.NewReader([]byte{
    9,1,2,3,4,5,6,8,
  })
  m.Run(f)
  /*
  f.Close()
  */
}
