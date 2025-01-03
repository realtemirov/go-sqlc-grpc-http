package utils_test

import (
	"testing"

	"github.com/realtemirov/go-sqlc-grpc-http/utils"
)

func Test_HasArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		val  string
		want bool
	}{
		{
			name: "Has array successful",
			arr:  []string{"a", "b", "c"},
			val:  "b",
			want: true,
		},
		{
			name: "Has array failed",
			arr:  []string{"a", "b", "c"},
			val:  "d",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.HasArray(tt.val, tt.arr); got != tt.want {
				t.Errorf("HasArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ArrayFirstOrDefault(t *testing.T) {
	tests := []struct {
		name  string
		arr   []string
		index int
		want  string
	}{
		{
			name:  "Has element in array successful",
			arr:   []string{"a", "b", "c"},
			index: 2,
			want:  "c",
		},
		{
			name:  "Has element in array failed",
			arr:   []string{"a", "b", "c"},
			index: 3,
			want:  "",
		},
		{
			name:  "Has element in array failed",
			arr:   []string{},
			index: 1,
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ArrayIndexOrDefault(tt.arr, tt.index); got != tt.want {
				t.Errorf("ArrayFirstOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
