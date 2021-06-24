package ringbuffer

import "errors"

var ErrIsEmpty = errors.New("ringbuffer is empty")

type RingBuffer struct {
	buf         []interface{}
	initialSize int
	size        int
	r           int
	w           int
}

func New(initialSize int) *RingBuffer {
	if initialSize < 2 {
		initialSize = 2
	}

	return &RingBuffer{
		buf:         make([]interface{}, initialSize),
		initialSize: initialSize,
		size:        initialSize,
		r:           0,
		w:           0,
	}
}

func (rb *RingBuffer) Read() (interface{}, error) {
	if rb.r == rb.w {
		return nil, ErrIsEmpty
	}

	v := rb.buf[rb.r]
	rb.r++
	if rb.r == rb.size {
		rb.r = 0
	}

	return v, nil
}

func (rb *RingBuffer) Pop() interface{} {
	v, err := rb.Read()
	if errors.Is(err, ErrIsEmpty) {
		panic(ErrIsEmpty.Error())
	}

	return v
}

func (rb *RingBuffer) Peek() interface{} {
	if rb.r == rb.w {
		panic(ErrIsEmpty.Error())
	}

	v := rb.buf[rb.r]
	return v
}

func (rb *RingBuffer) Write(v interface{}) {
	rb.buf[rb.w] = v
	rb.w++

	if rb.w == rb.size {
		rb.w = 0
	}

	if rb.w == rb.r {
		rb.grow()
	}
}

func (rb *RingBuffer) grow() {
	var size int
	if rb.size < 1024 {
		size = rb.size << 1
	} else {
		size = rb.size + rb.size>>2
	}

	buf := make([]interface{}, size)
	copy(buf, rb.buf[rb.r:])
	copy(buf[rb.size-rb.r:], rb.buf[:rb.r])

	rb.r = 0
	rb.w = rb.size
	rb.size = size
	rb.buf = buf
}

func (rb *RingBuffer) IsEmpty() bool {
	return rb.r == rb.w
}

func (rb *RingBuffer) Capacity() int {
	return rb.size
}

func (rb *RingBuffer) Len() int {
	if rb.r == rb.w {
		return 0
	}
	if rb.r < rb.w {
		return rb.w - rb.r
	}
	return rb.size - rb.r + rb.w
}

func (rb *RingBuffer) Reset() {
	rb.r = 0
	rb.w = 0
	rb.size = rb.initialSize
	rb.buf = make([]interface{}, rb.initialSize)
}
