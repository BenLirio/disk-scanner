package main

import (
  "os"
  "io"
  "fmt"
)
var verb int = 100

type Machine struct {
  cs int
  p int
  pe int
  eof int
  data []byte
  BUF_LEN int
  offset int
  verb int
}
func newMachine(n int) *Machine {
  return &Machine{
    BUF_LEN: n,
    data: make([]byte,n),
    verb: verb,
  }
}
func (m *Machine) debug(verbLevel int, msg string) {
  if m.verb>verbLevel {
    loc := m.offset + m.p
    fmt.Printf("[0x%08X] %s\n", loc, msg)
  }
}

func (m *Machine) Run(filename string) {
  %%{
    machine JPEG_scanner;

    SOI = (0xFF 0xD8) @{ m.debug(2,"SOI") };
    EOI = (0xFF 0xD9) @{ m.debug(2,"EOI") };
    JPEG = SOI extend* EOI;
    JPEG_scanner := (extend* JPEG)* extend*;
    write data;
    access m.;
    variable p m.p;
    variable pe m.pe;
    variable eof m.eof;
    variable data m.data;
  }%%

  %% write init;

  f,err := os.Open(filename)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to open %s\n", filename)
  }
  for {
    n,err := f.Read(m.data)
    if err == io.EOF { break }
    if err != nil {
      fmt.Fprintf(os.Stderr, "Failed to read data\n")
    }
    m.p = 0
    m.pe = n
    %% write exec;
    m.offset += n
  }
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Supply a file")
    os.Exit(0)
  }
  m := newMachine(256)
  m.Run(os.Args[1])
}
