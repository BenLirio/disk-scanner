
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
var png []byte
var pngStart int
var pngEnd int

func uint32Val(data []byte,p int) uint32 {
  return binary.BigEndian.Uint32(data[p-3:p+1])
}
func uint8Val(data []byte,p int) uint8 {
  return data[p]
}

//line scanner.rl:32

//line scanner.rl:33


//line scanner.rl:66



//line scanner.go:44
var _scanner_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 2, 2, 
	3, 
}

var _scanner_key_offsets []byte = []byte{
	0, 0, 1, 2, 3, 4, 5, 6, 
	7, 8, 8, 8, 8, 8, 9, 9, 
	9, 9, 9, 9, 9, 9, 10, 11, 
	12, 12, 12, 12, 12, 
}

var _scanner_trans_keys []byte = []byte{
	137, 80, 78, 71, 13, 10, 26, 10, 
	73, 69, 78, 68, 
}

var _scanner_single_lengths []byte = []byte{
	0, 1, 1, 1, 1, 1, 1, 1, 
	1, 0, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 0, 0, 1, 1, 1, 
	0, 0, 0, 0, 0, 
}

var _scanner_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 
}

var _scanner_index_offsets []byte = []byte{
	0, 0, 2, 4, 6, 8, 10, 12, 
	14, 16, 17, 18, 19, 20, 22, 23, 
	24, 25, 26, 27, 28, 29, 31, 33, 
	35, 36, 37, 38, 39, 
}

var _scanner_trans_targs []byte = []byte{
	2, 0, 3, 0, 4, 0, 5, 0, 
	6, 0, 7, 0, 8, 0, 9, 0, 
	10, 11, 12, 13, 21, 14, 15, 16, 
	17, 18, 19, 20, 28, 22, 15, 23, 
	16, 24, 17, 25, 26, 27, 28, 10, 
	
}

var _scanner_trans_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 1, 0, 
	0, 0, 0, 3, 0, 0, 0, 0, 
	0, 0, 0, 0, 5, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 7, 0, 
	
}

const scanner_start int = 1
const scanner_first_final int = 28
const scanner_error int = 0

const scanner_en_main int = 1


//line scanner.rl:69

var BUF_LEN int = 1<<20
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
//  data := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
//  0x00,0x00,0x00,0x01,  0x00,0x00,0x00,0x00,  0x99,             0x00,0x00,0x00,0x00,
//  0x00,0x00,0x00,0x02,  0x00,0x00,0x00,0x00,  0x99,0x88,        0x00,0x00,0x00,0x00,
//  0x00,0x00,0x00,0x03,  0x49,0x45,0x4E,0x44,  0x99,0x88,0x77,   0x00,0x00,0x00,0x00,
//}
  copy(data,b)
  cs := 0
  p := 0
  pe := len(data)
  
//line scanner.go:135
	{
	cs = scanner_start
	}

//line scanner.rl:95
  
//line scanner.go:142
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
//line scanner.rl:36

    pngStart = p-7
    if verb { fmt.Println("PNG start:", pngStart) }
  
		case 1:
//line scanner.rl:40

    chunk.Length = uint32Val(data,p)
    if verb { fmt.Println("chunk.Length =", chunk.Length) }
  
		case 2:
//line scanner.rl:46

    if p + int(chunk.Length) > len(data) {
      log.Fatal("Buffered chunk not implemented")
    }
    p += int(chunk.Length)
  
		case 3:
//line scanner.rl:52

    pngEnd = p+1
    if verb { fmt.Println("PNG end:", pngEnd) }
  
//line scanner.go:246
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

//line scanner.rl:96
}
