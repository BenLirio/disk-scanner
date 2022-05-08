package main
import (
  "os"
  "fmt"
  "io"
  "encoding/binary"
)
var verb int = 10

var matching bool = false
const (
  BUF_LEN int = 256
  NOT_DONE int = -1
  DONE int = 0
)
func min(a int, b int) int { if a < b { return a } else { return b } }
/*
min data_size = 0x02 because of data_size num
*/
var sf (func(string, ...interface{}) string) = fmt.Sprintf

//TODO: thumbnail data
%%{
  machine JPEG_scanner;
  action store_uint16_to_X { storeUint16ToX() }
  action store_uint16_to_Y { storeUint16ToY() }
  action store_uint8_to_X { storeUint8ToX() }
  action store_uint8_to_Y { storeUint8ToY() }

  skipn = zlen >{
    skipn()
    if p == pe { fbreak; }
  };

  JFIF_ascii = 0x4A 0x46 0x49 0x46 0x00;
  _uint16 = extend{2};
  _uint8 = extend{1};
  version = extend{2};
  xdensity = _uint16;
  ydensity = _uint16;
  density = _uint8
            xdensity        @store_uint16_to_X
                            @{debug(7,sf("\t\txdensity=%d",m.X16))}
            ydensity        @store_uint16_to_Y
                            @{debug(7,sf("\t\tydensity=%d",m.Y16))}
            ;
  xthumbnail = _uint8;
  ythumbnail = _uint8;
  thumbnail_data = skipn;
  thumbnail = xthumbnail      @store_uint8_to_X
                              @{debug(7,sf("\t\txthumbnail=%d",m.X8))}
              ythumbnail      @store_uint8_to_Y
                              @{debug(7,sf("\t\tythumbnail=%d",m.Y8))}
                              @{m.ToSkip=int(m.X8)*int(m.Y8)}
              thumbnail_data  >{debug(3,sf("\t\tskip thumbnail n=%d",m.ToSkip))}
              ;
  APP0_magic = 0xFF 0xE0;
  APP0_len = _uint16;
  APP0 =  APP0_magic        @{debug(6,"\tAPP0_magic")}
          APP0_len          @{debug(6,"\tAPP0_len")}
          JFIF_ascii        @{debug(6,"\tJFIF_ascii")}
          version           @{debug(6,"\tversion")}
          density           @{debug(6,"\tdensity")}
          thumbnail         @{debug(6,"\tthumbnail")}
          ;

  header_magic = 0xFF extend;
  block_len = _uint16;
  block_data = skipn;
  block =   header_magic    @{debug(4,"\tblock_magic")}
            block_len       @{debug(6,"\tblock_len")}
                            @store_uint16_to_X
                            @{m.ToSkip=int(m.X16-2)}
                            @{debug(6,sf("\tBlock Length=%d",m.ToSkip))}
            block_data      >{debug(3,sf("\tskip block_data n=%d",m.ToSkip))}
            ;

  SOI_magic = 0xFF 0xD8;
  EOI_magic = 0xFF 0xD9;
  JFIF =  SOI_magic     @{debug(100,"SOI")}
          APP0          @{debug(5,"APP0")}
          block         @{debug(5,"block")}
          EOI_magic     @{debug(5,"EOI")}
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
  ToSkip int
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
  skipn := func() {
    nextP := min(p+m.ToSkip,pe)
    m.ToSkip -= nextP-p
    p = nextP
  }
  if m.ToSkip > 0 {
    nextP := min(p+m.ToSkip,pe)
    m.ToSkip -= nextP-p
    p = nextP
    if p == pe { return }
  }
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
  filename := "/home/ben/data/A.img"
  f,err := os.Open(filename)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to open %s\n", filename)
    return
  }
  m.Run(f)
  f.Close()
}
