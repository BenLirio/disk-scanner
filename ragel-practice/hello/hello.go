
//line hello.rl:1
package main

import (
)


//line hello.rl:7

//line hello.rl:8

//line hello.go:14
var _url_prot_key_offsets []byte = []byte{
	0, 0, 1, 2, 3, 4, 5, 6, 
}

var _url_prot_trans_keys []byte = []byte{
	102, 111, 111, 98, 97, 114, 
}

var _url_prot_single_lengths []byte = []byte{
	0, 1, 1, 1, 1, 1, 1, 0, 
}

var _url_prot_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
}

var _url_prot_index_offsets []byte = []byte{
	0, 0, 2, 4, 6, 8, 10, 12, 
}

var _url_prot_trans_targs []byte = []byte{
	2, 0, 3, 0, 4, 0, 5, 0, 
	6, 0, 7, 0, 0, 
}

const url_prot_start int = 1
const url_prot_first_final int = 7
const url_prot_error int = 0

const url_prot_en_main int = 1


//line hello.rl:9


//line hello.rl:14


var cs int = 0  //cur_state
var p int = 0   //pos
var pe int = 0  //pos_end
var eof int = 0
var data []byte

func main() {
  
//line hello.go:61
	{
	cs = url_prot_start
	}

//line hello.rl:24
  for {
    
//line hello.go:69
	{
	var _klen int
	var _trans int
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_url_prot_key_offsets[cs])
	_trans = int(_url_prot_index_offsets[cs])

	_klen = int(_url_prot_single_lengths[cs])
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
			case data[p] < _url_prot_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _url_prot_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_url_prot_range_lengths[cs])
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
			case data[p] < _url_prot_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _url_prot_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	cs = int(_url_prot_trans_targs[_trans])

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

//line hello.rl:26
    break
  }
}
