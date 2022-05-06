
//line pumping.rl:1
package main

import (
  "fmt"
)

var stack int = 0

//line pumping.rl:27



//line pumping.go:16
var _pumping_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 
}

var _pumping_key_offsets []byte = []byte{
	0, 0, 2, 
}

var _pumping_trans_keys []byte = []byte{
	97, 98, 98, 
}

var _pumping_single_lengths []byte = []byte{
	0, 2, 1, 
}

var _pumping_range_lengths []byte = []byte{
	0, 0, 0, 
}

var _pumping_index_offsets []byte = []byte{
	0, 0, 3, 
}

var _pumping_trans_targs []byte = []byte{
	1, 2, 0, 2, 0, 
}

var _pumping_trans_actions []byte = []byte{
	3, 5, 0, 5, 0, 
}

var _pumping_eof_actions []byte = []byte{
	0, 1, 1, 
}

const pumping_start int = 1
const pumping_first_final int = 1
const pumping_error int = 0

const pumping_en_pumping int = 1


//line pumping.rl:30

var cs int
var p int
var pe int
var data []byte
var eof int
func main() {
  
//line pumping.go:69
	{
	cs = pumping_start
	}

//line pumping.rl:38

  data = []byte("aaabbb")
  
//line pumping.go:78
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
	_keys = int(_pumping_key_offsets[cs])
	_trans = int(_pumping_index_offsets[cs])

	_klen = int(_pumping_single_lengths[cs])
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
			case data[p] < _pumping_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _pumping_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_pumping_range_lengths[cs])
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
			case data[p] < _pumping_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _pumping_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	cs = int(_pumping_trans_targs[_trans])

	if _pumping_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_pumping_trans_actions[_trans])
	_nacts = uint(_pumping_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _pumping_actions[_acts-1] {
		case 1:
//line pumping.rl:13

    fmt.Println("push")
    stack += 1
  
		case 2:
//line pumping.rl:17

    fmt.Println("pop")
    stack -= 1
  
//line pumping.go:168
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
		__acts := _pumping_eof_actions[cs]
		__nacts := uint(_pumping_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _pumping_actions[__acts-1] {
			case 0:
//line pumping.rl:10

    fmt.Println("log")
  
//line pumping.go:192
			}
		}
	}

	_out: {}
	}

//line pumping.rl:41
  fmt.Println(stack)
}
