package sliceutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name  string
		slice []int64
		fn    func(i int64) bool
		want  []int64
	}{
		{
			name:  "偶数のみを抽出する",
			slice: []int64{1, 2, 3, 4, 5, 6},
			fn: func(i int64) bool {
				return i%2 == 0
			},
			want: []int64{2, 4, 6},
		},
		{
			name:  "3より大きい値のみを抽出する",
			slice: []int64{1, 2, 3, 4, 5, 6},
			fn: func(i int64) bool {
				return i > 3
			},
			want: []int64{4, 5, 6},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.slice, tt.fn)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		elem  int
		want  bool
	}{
		{
			name:  "Contains 1",
			slice: []int{1, 2, 3, 4, 5},
			elem:  1,
			want:  true,
		},
		{
			name:  "Contains 6",
			slice: []int{1, 2, 3, 4, 5},
			elem:  6,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.slice, tt.elem)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSome(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		fn    func(string) bool
		want  bool
	}{
		{
			name:  "barが含まれている",
			slice: []string{"foo", "bar", "baz"},
			fn: func(s string) bool {
				return s == "bar"
			},
			want: true,
		},
		{
			name:  "quxが含まれていない",
			slice: []string{"foo", "bar", "baz"},
			fn: func(s string) bool {
				return s == "qux"
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Some(tt.slice, tt.fn)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		fn    func(i, j int) bool
		want  []int
	}{
		{
			name:  "昇順でソートする",
			slice: []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			fn: func(i, j int) bool {
				return i < j
			},
			want: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			name:  "降順でソートする",
			slice: []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			fn: func(i, j int) bool {
				return i > j
			},
			want: []int{9, 6, 5, 5, 4, 3, 2, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sort(tt.slice, tt.fn)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		index   int
		element []int
		want    []int
	}{
		{
			name:    "先頭に要素を挿入する",
			slice:   []int{2, 3, 4},
			index:   0,
			element: []int{0, 1},
			want:    []int{0, 1, 2, 3, 4},
		},
		{
			name:    "中間に要素を挿入する",
			slice:   []int{0, 1, 4},
			index:   2,
			element: []int{2, 3},
			want:    []int{0, 1, 2, 3, 4},
		},
		{
			name:    "末尾に要素を挿入する",
			slice:   []int{0, 1, 2},
			index:   3,
			element: []int{3, 4},
			want:    []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Insert(tt.slice, tt.index, tt.element...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		index   int
		want    []int
		wantErr error
	}{
		{
			name:    "先頭の要素を削除",
			slice:   []int{0, 1, 2, 3, 4},
			index:   0,
			want:    []int{1, 2, 3, 4},
			wantErr: nil,
		},
		{
			name:    "中間の要素を削除",
			slice:   []int{0, 1, 2, 3, 4},
			index:   2,
			want:    []int{0, 1, 3, 4},
			wantErr: nil,
		},
		{
			name:    "末尾の要素を削除",
			slice:   []int{0, 1, 2, 3, 4},
			index:   4,
			want:    []int{0, 1, 2, 3},
			wantErr: nil,
		},
		{
			name:    "indexが要素外ならエラー",
			slice:   []int{0, 1, 2, 3, 4},
			index:   5,
			want:    nil,
			wantErr: fmt.Errorf("index out of range: %d", 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Delete(tt.slice, tt.index)
			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
