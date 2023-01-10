package main

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
5(汉字
6汉字)
7"汉字
8汉字" 
9 '汉字
10 汉字'`,
			},
			want: `1: 汉字：汉字
2; 汉字；汉字
3? 汉字？汉字
4! 汉字！汉字
5（汉字
6汉字）
7“汉字
8汉字” 
9 ‘汉字
10 汉字’`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replace(tt.args.data); !reflect.DeepEqual(strings.TrimSpace(got), strings.TrimSpace(tt.want)) {
				t.Errorf("replace2() = %v, want %v", got, tt.want)
			}
		})
	}
}
