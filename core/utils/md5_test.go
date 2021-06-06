package utils

import (
	"strings"
	"testing"
)

func TestMD5Sum(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		wantSum string
	}{
		{"tests", "cnbattle", "1126657366ED4F0394B16BC1B5086584"},
		{"tests", "sybmall85214", "dbedb83ed8f5e4e42a004c6c17976c06"},
		{"tests", "cnbattle1", "76588979196866A98BA41F694FB15A24"},
		{"tests", "cnbattle2", "1F4318D5BF8C757AB7C0CD0965C7B394"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := MD5Sum(tt.args); gotSum != strings.ToUpper(tt.wantSum) {
				t.Errorf("MD5Sum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
