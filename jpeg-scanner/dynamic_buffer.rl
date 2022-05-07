package main
import (
  "os"
  "fmt"
  "io"
  "encoding/binary"
)
var verb int = 100

var matching bool = false
const (
  BUF_LEN int = 256
  NOT_DONE int = -1
  DONE int = 0
)

/*
min data_size = 0x02 because of data_size num
*/

//TODO: thumbnail data
%%{
  machine JPEG_scanner;
  action store_uint16_to_X { storeUint16ToX() }
  action store_uint16_to_Y { storeUint16ToY() }
  action store_uint8_to_X { storeUint8ToX() }
  action store_uint8_to_Y { storeUint8ToY() }

  JFIF_ascii = 0x4A 0x46 0x49 0x46 0x00;
  _uint16 = extend{2};
  _uint8 = extend{1};
  version = extend{2};
  xdensity = _uint16;
  ydensity = _uint16;
  density = _uint8
            xdensity        @store_uint16_to_X @{debug(2,m.X16)}
            ydensity        @store_uint16_to_Y @{debug(2,m.Y16)}
            ;
  xthumbnail = _uint8;
  ythumbnail = _uint8;
  thumbnail_data = extend;
  thumbnail = xthumbnail      @store_uint8_to_X @{debug(2,m.X8)}
              ythumbnail      @store_uint8_to_Y @{debug(2,m.Y8)}
              thumbnail_data
              ;
  APP0_magic = 0xFF 0xE0;
  APP0 =  APP0_magic        @{debug(2,"APP0_magic")}
          _uint16           @{debug(2,"uint16")}
          JFIF_ascii        @{debug(2,"JFIF_ascii")}
          version           @{debug(2,"version")}
          density           @{debug(2,"density")}
          thumbnail         @{debug(2,"thumbnail")}
          ;

  SOI_magic = 0xFF 0xD8;
  EOI_magic = 0xFF 0xD9;
  JFIF =  SOI_magic     @{debug(2,"SOI")}
          APP0          @{debug(2,"APP0")}
          EOI_magic     @{debug(2,"EOI")}
          ;
  JPEG_scanner :=
  |*
  extend;
  JFIF;
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
  X16 uint16
  Y16 uint16
  X8 uint8
  Y8 uint8
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
  debug := func(verbLevel int, msg interface{}) {
    if verb>verbLevel {
      loc := m.offset + p
      fmt.Printf("[0x%08X] %v\n", loc, msg)
    }
  }
  getUint16 := func() uint16 { return binary.LittleEndian.Uint16(m.data[p-1:p+1]) }
  storeUint16ToX := func() { m.X16 = getUint16() }
  storeUint16ToY := func() { m.Y16 = getUint16() }
  storeUint8ToX := func() { m.X8 = uint8(m.data[p]) }
  storeUint8ToY := func() { m.Y8 = uint8(m.data[p]) }
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
  }
}

func main() {
  m := NewMachine()
  filename := "/home/ben/data/1.jpeg"
  f,err := os.Open(filename)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to open %s\n", filename)
    return
  }
  m.Run(f)
  f.Close()
}
