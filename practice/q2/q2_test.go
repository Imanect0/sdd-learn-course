package q2

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name    string
		list    []string
		want    []string
		wantErr bool
	}{
		{
			name:    "empty list",
			list:    []string{},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "no duplicates",
			list:    []string{"a", "b", "c"},
			want:    []string{"a", "b", "c"},
			wantErr: false,
		},
		{
			name:    "duplicates",
			list:    []string{"a", "b", "a", "c", "b"},
			want:    []string{"a", "b", "c"},
			wantErr: false,
		},
		{
			name:    "members",
			list:    []string{"go", "rust", "go", "go", "python", "python", "rust", "go", "go", "rust"},
			want:    []string{"go", "rust", "python"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveDuplicates(tt.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveDuplicates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicates() got = %v, want %v", got, tt.want)
			}
		})
	}
}