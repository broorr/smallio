package cmd

import (
	"log"
	"testing"
	"time"
)

func TestSyncOperation(t *testing.T) {
	type args struct {
		f func() error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "got current time",
			args: args{f: func() error {
				log.Printf("CurrentTime ==> %v", time.Now().Format("2006.01.02 15:04:05"))
				return nil
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SyncOperation(tt.args.f)
		})
	}
}
