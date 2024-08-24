package stream

import (
	"iter"
	"slices"
)

func Stream[T any](values []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range values {
			if !yield(v) {
				break
			}
		}
	}
}

func Collect[T any](it iter.Seq[T]) []T {
	return slices.Collect(it)
}

func Filter[T any](it iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if f(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
}

func Map[T1, T2 any](it iter.Seq[T1], f func(T1) T2) iter.Seq[T2] {
	return func(yield func(T2) bool) {
		for v := range it {
			if !yield(f(v)) {
				break
			}
		}
	}
}
