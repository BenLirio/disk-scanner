
//line jpeg-scanner.rl:1
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
  
//line jpeg-scanner.go:39
var _JPEG_scanner_actions []byte = []byte{
	0, 1, 0, 1, 1, 
}

var _JPEG_scanner_key_offsets []byte = []byte{
	0, 1, 3, 4, 
}

var _JPEG_scanner_trans_keys []byte = []byte{
	255, 216, 255, 255, 216, 217, 255, 
}

var _JPEG_scanner_single_lengths []byte = []byte{
	1, 2, 1, 3, 
}

var _JPEG_scanner_range_lengths []byte = []byte{
	0, 0, 0, 0, 
}

var _JPEG_scanner_index_offsets []byte = []byte{
	0, 2, 5, 7, 
}

var _JPEG_scanner_trans_targs []byte = []byte{
	1, 0, 2, 1, 0, 3, 2, 2, 
	2, 3, 2, 
}

var _JPEG_scanner_trans_actions []byte = []byte{
	0, 0, 1, 0, 0, 0, 0, 1, 
	3, 0, 0, 
}

const JPEG_scanner_start int = 0
const JPEG_scanner_first_final int = 0
const JPEG_scanner_error int = -1

const JPEG_scanner_en_JPEG_scanner int = 0


//line jpeg-scanner.rl:48


  
//line jpeg-scanner.go:85
	{
	 m.cs = JPEG_scanner_start
	}

//line jpeg-scanner.rl:51

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
    
//line jpeg-scanner.go:105
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if ( m.p) == ( m.pe) {
		goto _test_eof
	}
_resume:
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
			case ( m.data)[( m.p)] < _JPEG_scanner_trans_keys[_mid]:
				_upper = _mid - 1
			case ( m.data)[( m.p)] > _JPEG_scanner_trans_keys[_mid]:
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
			case ( m.data)[( m.p)] < _JPEG_scanner_trans_keys[_mid]:
				_upper = _mid - 2
			case ( m.data)[( m.p)] > _JPEG_scanner_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
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
//line jpeg-scanner.rl:38
 m.debug(2,"SOI") 
		case 1:
//line jpeg-scanner.rl:39
 m.debug(2,"EOI") 
//line jpeg-scanner.go:186
		}
	}

_again:
	( m.p)++
	if ( m.p) != ( m.pe) {
		goto _resume
	}
	_test_eof: {}
	}

//line jpeg-scanner.rl:65
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
