package util

import (
	"reflect"
	"strings"
	"testing"
)

func Test_replace(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "句号",
			args: args{
				data: `1: 汉字: 汉字
2; 汉字; 汉字
3? 汉字? 汉字
4! 汉字! 汉字
5(汉字)
6"汉字"
7'汉字'
8.(1)`,
			},
			want: `1: 汉字：汉字
2; 汉字；汉字
3? 汉字？汉字
4! 汉字！汉字
5（汉字）
6“汉字”
7‘汉字’
8.（1）`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Replace(tt.args.data); !reflect.DeepEqual(strings.TrimSpace(got), strings.TrimSpace(tt.want)) {
				t.Errorf("replace2() = %v, want %v", got, tt.want)
			}
		})
	}
}
