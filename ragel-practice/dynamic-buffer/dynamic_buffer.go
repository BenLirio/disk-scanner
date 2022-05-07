
//line dynamic_buffer.rl:1
package main
import (
  "os"
  "fmt"
  "bytes"
  "io"
)

var matching bool = false
const (
  BUF_LEN int = 1
  NOT_DONE int = -1
  DONE int = 0
)


//line dynamic_buffer.go:20
var _dynamic_buffer_actions []byte = []byte{
	0, 1, 4, 1, 5, 1, 7, 1, 9, 
	1, 10, 1, 11, 1, 12, 2, 3, 
	8, 2, 6, 0, 3, 6, 1, 2, 
}

var _dynamic_buffer_key_offsets []byte = []byte{
	0, 1, 2, 3, 4, 5, 6, 
}

var _dynamic_buffer_trans_keys []byte = []byte{
	3, 5, 6, 7, 1, 2, 4, 
}

var _dynamic_buffer_single_lengths []byte = []byte{
	1, 1, 1, 1, 1, 1, 1, 
}

var _dynamic_buffer_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 
}

var _dynamic_buffer_index_offsets []byte = []byte{
	0, 2, 4, 6, 8, 10, 12, 
}

var _dynamic_buffer_trans_targs []byte = []byte{
	6, 4, 2, 4, 3, 4, 4, 4, 
	5, 4, 0, 4, 1, 4, 4, 4, 
	4, 4, 4, 4, 
}

var _dynamic_buffer_trans_actions []byte = []byte{
	18, 11, 0, 13, 0, 13, 15, 13, 
	21, 5, 0, 7, 0, 9, 11, 13, 
	13, 13, 7, 9, 
}

var _dynamic_buffer_to_state_actions []byte = []byte{
	0, 0, 0, 0, 1, 0, 0, 
}

var _dynamic_buffer_from_state_actions []byte = []byte{
	0, 0, 0, 0, 3, 0, 0, 
}

var _dynamic_buffer_eof_trans []byte = []byte{
	15, 18, 18, 18, 0, 19, 20, 
}

const dynamic_buffer_start int = 4
const dynamic_buffer_first_final int = 4
const dynamic_buffer_error int = -1

const dynamic_buffer_en_dynamic_buffer int = 4


//line dynamic_buffer.rl:39


type Machine struct {
  ts int
  te int
  act int
  cs int
  data []byte
  offset int
}

func (m *Machine) init() {
  
//line dynamic_buffer.rl:52
  
//line dynamic_buffer.go:94
	{
	 m.cs = dynamic_buffer_start
	 m.ts = 0
	 m.te = 0
	 m.act = 0
	}

//line dynamic_buffer.rl:53
}
func NewMachine() *Machine {
  m := &Machine{}
  m.init()
  return m
}

func (m *Machine) exec(p int, pe int, eof int) {
  
//line dynamic_buffer.go:112
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
	_acts = int(_dynamic_buffer_from_state_actions[ m.cs])
	_nacts = uint(_dynamic_buffer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _dynamic_buffer_actions[_acts - 1] {
		case 5:
//line NONE:1
 m.ts = p

//line dynamic_buffer.go:132
		}
	}

	_keys = int(_dynamic_buffer_key_offsets[ m.cs])
	_trans = int(_dynamic_buffer_index_offsets[ m.cs])

	_klen = int(_dynamic_buffer_single_lengths[ m.cs])
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
			case  m.data[p] < _dynamic_buffer_trans_keys[_mid]:
				_upper = _mid - 1
			case  m.data[p] > _dynamic_buffer_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_dynamic_buffer_range_lengths[ m.cs])
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
			case  m.data[p] < _dynamic_buffer_trans_keys[_mid]:
				_upper = _mid - 2
			case  m.data[p] > _dynamic_buffer_trans_keys[_mid + 1]:
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
	 m.cs = int(_dynamic_buffer_trans_targs[_trans])

	if _dynamic_buffer_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_dynamic_buffer_trans_actions[_trans])
	_nacts = uint(_dynamic_buffer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _dynamic_buffer_actions[_acts-1] {
		case 0:
//line dynamic_buffer.rl:18

    fmt.Println("short_match")
  
		case 1:
//line dynamic_buffer.rl:21

    fmt.Println("begin_short_match")
  
		case 2:
//line dynamic_buffer.rl:24

    fmt.Println("long_long_match")
  
		case 3:
//line dynamic_buffer.rl:27

    fmt.Println("long_match")
  
		case 6:
//line NONE:1
 m.te = p+1

		case 7:
//line dynamic_buffer.rl:34
 m.te = p+1

		case 8:
//line dynamic_buffer.rl:36
 m.te = p+1

		case 9:
//line dynamic_buffer.rl:34
 m.te = p
p--

		case 10:
//line dynamic_buffer.rl:35
 m.te = p
p--

		case 11:
//line dynamic_buffer.rl:34
p = ( m.te) - 1

		case 12:
//line dynamic_buffer.rl:35
p = ( m.te) - 1

//line dynamic_buffer.go:251
		}
	}

_again:
	_acts = int(_dynamic_buffer_to_state_actions[ m.cs])
	_nacts = uint(_dynamic_buffer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _dynamic_buffer_actions[_acts-1] {
		case 4:
//line NONE:1
 m.ts = 0

//line dynamic_buffer.go:265
		}
	}

	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		if _dynamic_buffer_eof_trans[ m.cs] > 0 {
			_trans = int(_dynamic_buffer_eof_trans[ m.cs] - 1)
			goto _eof_trans
		}
	}

	}

//line dynamic_buffer.rl:62
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
    fmt.Printf("offset=%d\tts=%d\tte=%d\tact=%d\tdata=%v\n\n",m.offset,m.ts, m.te, m.act, m.data)
  }
}

func main() {
  m := NewMachine()
  /*
  filename := "data/t1"
  f,err := os.Open(filename)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to open %s\n", filename)
    return
  }
  */
  f := bytes.NewReader([]byte{
    9,1,2,3,4,5,6,8,
  })
  m.Run(f)
  /*
  f.Close()
  */
}
