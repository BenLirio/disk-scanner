
//line dynamic_buffer.rl:1
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

//line dynamic_buffer.go:26
var _JPEG_scanner_actions []byte = []byte{
	0, 1, 8, 1, 9, 1, 10, 1, 11, 
	1, 14, 1, 17, 1, 18, 1, 19, 
	1, 20, 1, 22, 1, 23, 2, 0, 
	4, 2, 2, 6, 2, 3, 7, 2, 
	13, 15, 2, 16, 21, 3, 1, 5, 
	12, 
}

var _JPEG_scanner_key_offsets []byte = []byte{
	0, 1, 2, 2, 2, 3, 4, 5, 
	6, 7, 7, 7, 7, 7, 7, 7, 
	7, 7, 7, 7, 8, 9, 10, 
}

var _JPEG_scanner_trans_keys []byte = []byte{
	255, 224, 74, 70, 73, 70, 0, 255, 
	217, 255, 216, 
}

var _JPEG_scanner_single_lengths []byte = []byte{
	1, 1, 0, 0, 1, 1, 1, 1, 
	1, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 1, 1, 1, 1, 
}

var _JPEG_scanner_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 
}

var _JPEG_scanner_index_offsets []byte = []byte{
	0, 2, 4, 5, 6, 8, 10, 12, 
	14, 16, 17, 18, 19, 20, 21, 22, 
	23, 24, 25, 26, 28, 30, 32, 
}

var _JPEG_scanner_trans_targs []byte = []byte{
	1, 21, 2, 21, 3, 4, 5, 21, 
	6, 21, 7, 21, 8, 21, 9, 21, 
	10, 11, 12, 13, 14, 15, 16, 17, 
	18, 19, 20, 21, 21, 21, 22, 21, 
	0, 21, 21, 21, 21, 21, 21, 21, 
	21, 21, 21, 21, 21, 21, 21, 21, 
	21, 21, 21, 21, 21, 21, 21, 21, 
	
}

var _JPEG_scanner_trans_actions []byte = []byte{
	0, 21, 1, 21, 0, 3, 0, 21, 
	0, 21, 0, 21, 0, 21, 5, 21, 
	0, 7, 0, 0, 23, 0, 38, 26, 
	29, 32, 0, 21, 35, 21, 15, 17, 
	9, 19, 21, 21, 21, 21, 21, 21, 
	21, 21, 21, 21, 21, 21, 21, 21, 
	21, 21, 21, 21, 21, 21, 21, 19, 
	
}

var _JPEG_scanner_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 11, 0, 
}

var _JPEG_scanner_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 13, 0, 
}

var _JPEG_scanner_eof_trans []byte = []byte{
	55, 55, 55, 55, 55, 55, 55, 55, 
	55, 55, 55, 55, 55, 55, 55, 55, 
	55, 55, 55, 55, 55, 0, 56, 
}

const JPEG_scanner_start int = 21
const JPEG_scanner_first_final int = 21
const JPEG_scanner_error int = -1

const JPEG_scanner_en_JPEG_scanner int = 21


//line dynamic_buffer.rl:67


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
  
//line dynamic_buffer.rl:84
  
//line dynamic_buffer.go:132
	{
	 m.cs = JPEG_scanner_start
	 m.ts = 0
	 m.te = 0
	 m.act = 0
	}

//line dynamic_buffer.rl:85
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
  
//line dynamic_buffer.go:161
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
		case 18:
//line NONE:1
 m.ts = p

//line dynamic_buffer.go:181
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
//line dynamic_buffer.rl:24
 storeUint16ToX() 
		case 1:
//line dynamic_buffer.rl:25
 storeUint16ToY() 
		case 2:
//line dynamic_buffer.rl:26
 storeUint8ToX() 
		case 3:
//line dynamic_buffer.rl:27
 storeUint8ToY() 
		case 4:
//line dynamic_buffer.rl:36
debug(2,m.X16)
		case 5:
//line dynamic_buffer.rl:37
debug(2,m.Y16)
		case 6:
//line dynamic_buffer.rl:42
debug(2,m.X8)
		case 7:
//line dynamic_buffer.rl:43
debug(2,m.Y8)
		case 8:
//line dynamic_buffer.rl:47
debug(2,"APP0_magic")
		case 9:
//line dynamic_buffer.rl:48
debug(2,"uint16")
		case 10:
//line dynamic_buffer.rl:49
debug(2,"JFIF_ascii")
		case 11:
//line dynamic_buffer.rl:50
debug(2,"version")
		case 12:
//line dynamic_buffer.rl:51
debug(2,"density")
		case 13:
//line dynamic_buffer.rl:52
debug(2,"thumbnail")
		case 14:
//line dynamic_buffer.rl:57
debug(2,"SOI")
		case 15:
//line dynamic_buffer.rl:58
debug(2,"APP0")
		case 16:
//line dynamic_buffer.rl:59
debug(2,"EOI")
		case 19:
//line NONE:1
 m.te = p+1

		case 20:
//line dynamic_buffer.rl:63
 m.te = p+1

		case 21:
//line dynamic_buffer.rl:64
 m.te = p+1

		case 22:
//line dynamic_buffer.rl:63
 m.te = p
p--

		case 23:
//line dynamic_buffer.rl:63
p = ( m.te) - 1

//line dynamic_buffer.go:322
		}
	}

_again:
	_acts = int(_JPEG_scanner_to_state_actions[ m.cs])
	_nacts = uint(_JPEG_scanner_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _JPEG_scanner_actions[_acts-1] {
		case 17:
//line NONE:1
 m.ts = 0

//line dynamic_buffer.go:336
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

	}

//line dynamic_buffer.rl:105
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
