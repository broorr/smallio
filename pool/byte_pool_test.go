package pool

import (
	"reflect"
	"testing"

	"maggie-chou/tester"
)

func (s) TestBytePool_Get(t *testing.T) {
	type fields struct {
		content chan []byte
		size    int64
		cap     int64
	}
	tests := []struct {
		name   string
		fields fields
		wantC  []byte
	}{
		{
			name: "byte pool get data",
			fields: fields{
				// 初始化一个最大容量8比特的通信管道
				content: make(chan []byte, 8),
				size:    0,
				cap:     8,
			},
			wantC: []byte("hello"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &BytePool{
				content: tt.fields.content,
				size:    tt.fields.size,
				cap:     tt.fields.cap,
			}
			p.Put([]byte("hello"))
			if gotC := p.Get(); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("Get() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func (s) TestBytePool_Put(t *testing.T) {
	type fields struct {
		content chan []byte
		size    int64
		cap     int64
	}
	type args struct {
		c []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantC  []byte
	}{
		{
			name: "byte pool put and get",
			fields: fields{
				content: make(chan []byte, 1),
				size:    1,
				cap:     1,
			},
			args:  args{c: []byte("1234567890")},
			wantC: []byte("1234567890"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &BytePool{
				content: tt.fields.content,
				size:    tt.fields.size,
				cap:     tt.fields.cap,
			}
			p.Put(tt.args.c)
			close(p.content)
			if gotC := p.Get(); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("Get() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func (s) TestNewBytePool(t *testing.T) {
	type args struct {
		maxSize int64
		size    int64
		cap     int64
	}
	tests := []struct {
		name string
		args args
		want *BytePool
	}{
		{
			name: "new byte pool",
			args: args{
				maxSize: 16,
				size:    0,
				cap:     10,
			},
			want: NewBytePool(16, 0, 10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBytePool(tt.args.maxSize, tt.args.size, tt.args.cap); !reflect.DeepEqual(got.size,
				tt.want.size) {
				t.Errorf("NewBytePool() = %v, want %v", got, tt.want)
			}
			if got := NewBytePool(tt.args.maxSize, tt.args.size, tt.args.cap); !reflect.DeepEqual(got.cap,
				tt.want.cap) {
				t.Errorf("NewBytePool() = %v, want %v", got, tt.want)
			}
			if got := NewBytePool(tt.args.maxSize, tt.args.size, tt.args.cap); !reflect.DeepEqual(cap(got.content),
				cap(tt.want.content)) {
				t.Errorf("NewBytePool() = %v, want %v", got, tt.want)
			}
		})
	}
}

type s struct {
	tester.Preparer
}

func TestRun(t *testing.T) {
	tester.RunTests(t, s{})
}
