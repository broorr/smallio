package channel

// BytePool 字节缓冲池
type BytePool struct {
	// 字节池内容
	content chan []byte
	// 字节池长度
	size int64
	// 字节池容量
	cap int64
}

// NewBytePool 构造一个字节池
func NewBytePool(maxSize, size, cap int64) *BytePool {
	return &BytePool{
		content: make(chan []byte, maxSize),
		size:    size,
		cap:     cap,
	}
}

// Put 写入数据到字节池
func (p *BytePool) Put(c []byte) {
	select {
	case p.content <- c:
	default:
	}
}

// Get 从字节池读取数据
func (p *BytePool) Get() (c []byte) {
	select {
	case c = <-p.content:
	default:
		if p.cap > 0 {
			c = make([]byte, p.size, p.cap)
		} else {
			c = make([]byte, p.size)
		}
	}
	return
}
