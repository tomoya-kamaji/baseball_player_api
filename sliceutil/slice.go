package sliceutil

import (
	"fmt"
	"sort"
)

type User struct {
	ID   int
	Name string
}

type Users []User

// Filter ... 条件に合う要素のみを抽出する
func Filter[T any](srcs []T, fn func(src T) bool) []T {
	dsts := []T{}
	for _, src := range srcs {
		if fn(src) {
			dsts = append(dsts, src)
		}
	}
	return dsts
}

// Contains ... 配列に要素が含まれているか
func Contains[T comparable](srcs []T, e T) bool {
	for _, v := range srcs {
		if e == v {
			return true
		}
	}
	return false
}

// Some ... 配列の要素のいずれかが条件を満たすか
func Some[T any](srcs []T, fn func(src T) bool) bool {
	for _, src := range srcs {
		if fn(src) {
			return true
		}
	}
	return false

}

// Sort ... ソート
func Sort[T any](srcs []T, fn func(i, j int) bool) []T {
	sort.SliceStable(srcs, fn)
	return srcs
}

// Insert ... 配列の任意の場所に挿入する
func Insert[T any](srcs []T, i int, v ...T) []T {
	// TODO: エラー処理
	return append(srcs[:i], append(v, srcs[i:]...)...)
}

// Delete ... 配列の任意の値をindex指定で削除する
func Delete[T any](srcs []T, i int) ([]T, error) {
	if i < 0 || i >= len(srcs) {
		return nil, fmt.Errorf("index out of range: %d", i)
	}
	return append(srcs[:i], srcs[i+1:]...), nil
}

// 以下、まだテスト書けてないのでTODO

// First ... 配列の先頭の要素を取得
func First[T any](srcs []T) T {
	// out of range対策。errorにするかnilを返すかは要検討
	return srcs[0]
}

// Last ... 配列の最後の要素を取得
func Last[T any](srcs []T) T {
	// out of range対策。errorにするかnilを返すかは要検討
	return srcs[len(srcs)-1]
}

// Shift ... 配列の先頭を切り取る (破壊)
func Shift[T any](srcs []T) (T, []T) {
	// out of range対策。errorにするかnilを返すかは要検討 Shiftはnilを返すのが良さそう
	return srcs[0], srcs[1:]
}

// Pop ... 配列の後尾を切り取る (破壊)
func Pop[T any](srcs []T) (T, []T) {
	// out of range対策。errorにするかnilを返すかは要検討 Popはnilを返すのが良さそう
	return srcs[len(srcs)-1], srcs[:len(srcs)-1]
}

// Uniq ... 配列の重複を排除する
func Uniq[T comparable](srcs []T) []T {
	dsts := make([]T, 0, len(srcs))
	m := make(map[T]bool)
	for _, src := range srcs {
		if _, ok := m[src]; !ok {
			m[src] = true
			dsts = append(dsts, src)
		}
	}
	return dsts
}

// TODO：和、積、差などの集合演算
