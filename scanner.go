
//line scanner.rl:1
package main

import (
  "log"
  "fmt"
  "io"
  "os"
  "encoding/binary"
)

var eof int = 0

type Chunk struct {
  Length uint32
  Type uint32
  Data []byte
  CRC uint32
}
var chunk Chunk
var verb bool = true

type Png struct {
  Height uint32
  Width uint32
  BitDepth uint8
  Chunks []Chunk
}
var png Png
func uint32Val(data []byte,p int) uint32 {
  return binary.BigEndian.Uint32(data[p-3:p+1])
}
func uint8Val(data []byte,p int) uint8 {
  return data[p]
}

//line scanner.rl:36

//line scanner.rl:37

//line scanner.rl:38


//line scanner.rl:84



//line scanner.go:50
var _scanner_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 
}

var _scanner_key_offsets []byte = []byte{
	0, 0, 1, 2, 3, 4, 5, 6, 
	7, 8, 8, 8, 8, 8, 9, 10, 
	11, 12, 12, 12, 12, 12, 12, 12, 
	12, 12, 12, 
}

var _scanner_trans_keys []byte = []byte{
	137, 80, 78, 71, 13, 10, 26, 10, 
	73, 72, 68, 82, 
}

var _scanner_single_lengths []byte = []byte{
	0, 1, 1, 1, 1, 1, 1, 1, 
	1, 0, 0, 0, 0, 1, 1, 1, 
	1, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 
}

var _scanner_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 
}

var _scanner_index_offsets []byte = []byte{
	0, 0, 2, 4, 6, 8, 10, 12, 
	14, 16, 17, 18, 19, 20, 22, 24, 
	26, 28, 29, 30, 31, 32, 33, 34, 
	35, 36, 37, 
}

var _scanner_trans_targs []byte = []byte{
	2, 0, 3, 0, 4, 0, 5, 0, 
	6, 0, 7, 0, 8, 0, 9, 0, 
	10, 11, 12, 13, 14, 0, 15, 0, 
	16, 0, 17, 0, 18, 19, 20, 21, 
	22, 23, 24, 25, 26, 10, 
}

var _scanner_trans_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 1, 0, 
	0, 0, 0, 3, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 5, 
	0, 0, 0, 7, 9, 0, 
}

const scanner_start int = 1
const scanner_first_final int = 26
const scanner_error int = 0

const scanner_en_main int = 1


//line scanner.rl:87

var BUF_LEN int = 256
func main() {
  f, err := os.Open("./sample.png")
  if err != nil {
    log.Fatal("failed to open file")
  }
  b := make([]byte, BUF_LEN)
  n, err := f.Read(b)
  if err == io.EOF {
    os.Exit(0)
  }
  if err != nil {
    log.Fatal("Error reading")
  }
  data := make([]byte, n)
  copy(data,b)
  cs := 0
  p := 0
  pe := len(data)
  
//line scanner.go:134
	{
	cs = scanner_start
	}

//line scanner.rl:108
  
//line scanner.go:141
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_scanner_key_offsets[cs])
	_trans = int(_scanner_index_offsets[cs])

	_klen = int(_scanner_single_lengths[cs])
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
			case data[p] < _scanner_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _scanner_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_scanner_range_lengths[cs])
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
			case data[p] < _scanner_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _scanner_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	cs = int(_scanner_trans_targs[_trans])

	if _scanner_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_scanner_trans_actions[_trans])
	_nacts = uint(_scanner_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _scanner_actions[_acts-1] {
		case 0:
//line scanner.rl:41

    png = Png{}
    if verb { fmt.Println("Png magic") }
  
		case 1:
//line scanner.rl:45

    chunk.Length = binary.BigEndian.Uint32(data[p-3:p+1])
    chunk.Length = uint32Val(data,p)
    if verb { fmt.Println("chunk.Length =", chunk.Length) }
  
		case 2:
//line scanner.rl:62

    png.Width = uint32Val(data,p)
    if verb { fmt.Println("Png Width =", png.Width) }
  
		case 3:
//line scanner.rl:66

    png.Height = uint32Val(data,p)
    if verb { fmt.Println("Png Height =", png.Height) }
  
		case 4:
//line scanner.rl:70

    png.BitDepth = uint8Val(data,p)
    if verb { fmt.Println("Bit Depth =", png.BitDepth) }
  
//line scanner.go:250
		}
	}

_again:
	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	_out: {}
	}

//line scanner.rl:109
}
