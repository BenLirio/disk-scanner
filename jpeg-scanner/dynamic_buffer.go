
//line dynamic_buffer.rl:1
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

//line dynamic_buffer.go:27
var _JPEG_scanner_actions []byte = []byte{
	0, 1, 11, 1, 12, 1, 13, 1, 14, 
	1, 17, 1, 23, 1, 27, 1, 28, 
	1, 29, 1, 30, 1, 32, 1, 33, 
	2, 0, 5, 2, 2, 7, 2, 10, 
	4, 2, 26, 31, 3, 1, 6, 15, 
	5, 3, 8, 9, 16, 24, 7, 18, 
	0, 19, 20, 21, 22, 25, 
}

var _JPEG_scanner_key_offsets []byte = []byte{
	0, 1, 2, 2, 2, 3, 4, 5, 
	6, 7, 7, 7, 7, 7, 7, 7, 
	7, 7, 7, 8, 8, 8, 8, 9, 
	10, 11, 
}

var _JPEG_scanner_trans_keys []byte = []byte{
	255, 224, 74, 70, 73, 70, 0, 255, 
	255, 217, 255, 216, 
}

var _JPEG_scanner_single_lengths []byte = []byte{
	1, 1, 0, 0, 1, 1, 1, 1, 
	1, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 1, 0, 0, 0, 1, 1, 
	1, 1, 
}

var _JPEG_scanner_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 
}

var _JPEG_scanner_index_offsets []byte = []byte{
	0, 2, 4, 5, 6, 8, 10, 12, 
	14, 16, 17, 18, 19, 20, 21, 22, 
	23, 24, 25, 27, 28, 29, 30, 32, 
	34, 36, 
}

var _JPEG_scanner_trans_targs []byte = []byte{
	1, 24, 2, 24, 3, 4, 5, 24, 
	6, 24, 7, 24, 8, 24, 9, 24, 
	10, 11, 12, 13, 14, 15, 16, 17, 
	18, 19, 24, 20, 21, 22, 23, 24, 
	24, 24, 25, 24, 0, 24, 24, 24, 
	24, 24, 24, 24, 24, 24, 24, 24, 
	24, 24, 24, 24, 24, 24, 24, 24, 
	24, 24, 24, 24, 24, 24, 24, 
}

var _JPEG_scanner_trans_actions []byte = []byte{
	0, 23, 1, 23, 0, 3, 0, 23, 
	0, 23, 0, 23, 0, 23, 5, 23, 
	0, 7, 0, 0, 25, 0, 37, 28, 
	41, 31, 23, 9, 0, 47, 0, 23, 
	34, 23, 17, 19, 11, 21, 23, 23, 
	23, 23, 23, 23, 23, 23, 23, 23, 
	23, 23, 23, 23, 23, 23, 23, 23, 
	23, 23, 23, 23, 23, 23, 21, 
}

var _JPEG_scanner_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	13, 0, 
}

var _JPEG_scanner_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	15, 0, 
}

var _JPEG_scanner_eof_trans []byte = []byte{
	62, 62, 62, 62, 62, 62, 62, 62, 
	62, 62, 62, 62, 62, 62, 62, 62, 
	62, 62, 62, 62, 62, 62, 62, 62, 
	0, 63, 
}

const JPEG_scanner_start int = 24
const JPEG_scanner_first_final int = 24
const JPEG_scanner_error int = -1

const JPEG_scanner_en_JPEG_scanner int = 24


//line dynamic_buffer.rl:92


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
  
//line dynamic_buffer.rl:110
  
//line dynamic_buffer.go:142
	{
	 m.cs = JPEG_scanner_start
	 m.ts = 0
	 m.te = 0
	 m.act = 0
	}

//line dynamic_buffer.rl:111
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
  
//line dynamic_buffer.go:182
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
_resume:
	_acts = int(_JPEG_scanner_from_state_actions[ m.cs])
	_nacts = uint(_JPEG_scanner_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _JPEG_scanner_actions[_acts - 1] {
		case 28:
//line NONE:1
 m.ts = p

//line dynamic_buffer.go:202
		}
	}

	_keys = int(_JPEG_scanner_key_offsets[ m.cs])
	_trans = int(_JPEG_scanner_index_offsets[ m.cs])

	_klen = int(_JPEG_scanner_single_lengths[ m.cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case  m.data[p] < _JPEG_scanner_trans_keys[_mid]:
				_upper = _mid - 1
			case  m.data[p] > _JPEG_scanner_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_JPEG_scanner_range_lengths[ m.cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case  m.data[p] < _JPEG_scanner_trans_keys[_mid]:
				_upper = _mid - 2
			case  m.data[p] > _JPEG_scanner_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
_eof_trans:
	 m.cs = int(_JPEG_scanner_trans_targs[_trans])

	if _JPEG_scanner_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_JPEG_scanner_trans_actions[_trans])
	_nacts = uint(_JPEG_scanner_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _JPEG_scanner_actions[_acts-1] {
		case 0:
//line dynamic_buffer.rl:25
 storeUint16ToX() 
		case 1:
//line dynamic_buffer.rl:26
 storeUint16ToY() 
		case 2:
//line dynamic_buffer.rl:27
 storeUint8ToX() 
		case 3:
//line dynamic_buffer.rl:28
 storeUint8ToY() 
		case 4:
//line dynamic_buffer.rl:30

    skipn()
    if p == pe { p++; goto _out
 }
  
		case 5:
//line dynamic_buffer.rl:43
debug(7,sf("\t\txdensity=%d",m.X16))
		case 6:
//line dynamic_buffer.rl:45
debug(7,sf("\t\tydensity=%d",m.Y16))
		case 7:
//line dynamic_buffer.rl:51
debug(7,sf("\t\txthumbnail=%d",m.X8))
		case 8:
//line dynamic_buffer.rl:53
debug(7,sf("\t\tythumbnail=%d",m.Y8))
		case 9:
//line dynamic_buffer.rl:54
m.ToSkip=int(m.X8)*int(m.Y8)
		case 10:
//line dynamic_buffer.rl:55
debug(3,sf("\t\tskip thumbnail n=%d",m.ToSkip))
		case 11:
//line dynamic_buffer.rl:59
debug(6,"\tAPP0_magic")
		case 12:
//line dynamic_buffer.rl:60
debug(6,"\tAPP0_len")
		case 13:
//line dynamic_buffer.rl:61
debug(6,"\tJFIF_ascii")
		case 14:
//line dynamic_buffer.rl:62
debug(6,"\tversion")
		case 15:
//line dynamic_buffer.rl:63
debug(6,"\tdensity")
		case 16:
//line dynamic_buffer.rl:64
debug(6,"\tthumbnail")
		case 17:
//line dynamic_buffer.rl:70
debug(4,"\tblock_magic")
		case 18:
//line dynamic_buffer.rl:71
debug(6,"\tblock_len")
		case 19:
//line dynamic_buffer.rl:73
m.ToSkip=int(m.X16-2)
		case 20:
//line dynamic_buffer.rl:74
debug(6,sf("\tBlock Length=%d",m.ToSkip))
		case 21:
//line dynamic_buffer.rl:75
skipn()
		case 22:
//line dynamic_buffer.rl:76
if p==pe { p++; goto _out
 }
		case 23:
//line dynamic_buffer.rl:81
debug(100,"SOI")
		case 24:
//line dynamic_buffer.rl:82
debug(5,"APP0")
		case 25:
//line dynamic_buffer.rl:83
debug(5,"block")
		case 26:
//line dynamic_buffer.rl:84
debug(5,"EOI")
		case 29:
//line NONE:1
 m.te = p+1

		case 30:
//line dynamic_buffer.rl:88
 m.te = p+1

		case 31:
//line dynamic_buffer.rl:89
 m.te = p+1

		case 32:
//line dynamic_buffer.rl:88
 m.te = p
p--

		case 33:
//line dynamic_buffer.rl:88
p = ( m.te) - 1

//line dynamic_buffer.go:378
		}
	}

_again:
	_acts = int(_JPEG_scanner_to_state_actions[ m.cs])
	_nacts = uint(_JPEG_scanner_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _JPEG_scanner_actions[_acts-1] {
		case 27:
//line NONE:1
 m.ts = 0

//line dynamic_buffer.go:392
		}
	}

	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		if _JPEG_scanner_eof_trans[ m.cs] > 0 {
			_trans = int(_JPEG_scanner_eof_trans[ m.cs] - 1)
			goto _eof_trans
		}
	}

	_out: {}
	}

//line dynamic_buffer.rl:142
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
