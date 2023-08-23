// BEGIN: 6d7f5a3d4b1c
package q1

import (
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name    string
		list    []int
		want    int
		wantErr bool
	}{
		{
			name:    "empty list",
			list:    []int{},
			want:    0,
			wantErr: true,
		},
		{
			name:    "single element",
			list:    []int{1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "multiple elements 1",
			list:    []int{1, 2, 3},
			want:    6,
			wantErr: false,
		},
		{
			name:    "multiple elements 2",
			list:    []int{10, 23, 22, 7, 10, 33, 103, 4},
			want:    212,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sum(tt.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sum() got = %v, want %v", got, tt.want)
			}
		})
	}
}
// END: 6d7f5a3d4b1c