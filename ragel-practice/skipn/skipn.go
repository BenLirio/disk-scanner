
//line skipn.rl:1
package main

import (
  "fmt"
)


//line skipn.go:11
var _skipn_actions []byte = []byte{
	0, 1, 1, 1, 2, 2, 1, 0, 2, 
	3, 1, 
}

var _skipn_key_offsets []byte = []byte{
	0, 0, 3, 
}

var _skipn_trans_keys []byte = []byte{
	0, 3, 9, 0, 3, 9, 
}

var _skipn_single_lengths []byte = []byte{
	0, 3, 3, 
}

var _skipn_range_lengths []byte = []byte{
	0, 0, 0, 
}

var _skipn_index_offsets []byte = []byte{
	0, 0, 4, 
}

var _skipn_trans_targs []byte = []byte{
	2, 0, 1, 1, 2, 0, 1, 1, 
	
}

var _skipn_trans_actions []byte = []byte{
	5, 3, 8, 1, 5, 3, 8, 1, 
	
}

var _skipn_eof_actions []byte = []byte{
	0, 3, 0, 
}

const skipn_start int = 1
const skipn_first_final int = 2
const skipn_error int = 0

const skipn_en_skipn int = 1


//line skipn.rl:25



var cs int
var p int = 0
var pe int
var data []byte
var eof int = 0
func main() {
  data = []byte{
    0x01,0x03,
    0x04,0x03,0x03,0x03,0x03,
    0x02,0x03,0x03,
    0x09,0x00,0x00,
  }
  pe = len(data)
  
//line skipn.go:76
	{
	cs = skipn_start
	}

//line skipn.rl:42
  
//line skipn.go:83
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
	_keys = int(_skipn_key_offsets[cs])
	_trans = int(_skipn_index_offsets[cs])

	_klen = int(_skipn_single_lengths[cs])
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
			case data[p] < _skipn_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _skipn_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_skipn_range_lengths[cs])
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
			case data[p] < _skipn_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _skipn_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	cs = int(_skipn_trans_targs[_trans])

	if _skipn_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_skipn_trans_actions[_trans])
	_nacts = uint(_skipn_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _skipn_actions[_acts-1] {
		case 0:
//line skipn.rl:9

    fmt.Println("accept")
  
		case 1:
//line skipn.rl:12

    fmt.Printf("skip %d\n", data[p])
    p += int(data[p])
  
		case 2:
//line skipn.rl:16

    fmt.Printf("Reject at position %d.\n", p)
  
		case 3:
//line skipn.rl:19

    fmt.Printf("Done at %d\n", p)
    p++; goto _out

  
//line skipn.go:184
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
	if p == eof {
		__acts := _skipn_eof_actions[cs]
		__nacts := uint(_skipn_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _skipn_actions[__acts-1] {
			case 2:
//line skipn.rl:16

    fmt.Printf("Reject at position %d.\n", p)
  
//line skipn.go:208
			}
		}
	}

	_out: {}
	}

//line skipn.rl:43
}
