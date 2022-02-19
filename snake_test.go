package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	t.Run("Up", func(t *testing.T) {
		s := NewSnake(Coordinates{5, 5})
		s.Move(s.Turn(Up).Step())
		assert.Equal(t, NewSnake(Coordinates{5, 4}).Body, s.Body)
	})

	t.Run("Down", func(t *testing.T) {
		s := NewSnake(Coordinates{5, 5})
		s.Move(s.Turn(Down).Step())
		assert.Equal(t, NewSnake(Coordinates{5, 6}).Body, s.Body)
	})

	//	t.Run("Left", func(t *testing.T) {
	//		s := NewSnake(5, 5)
	//		s.Move(Left)
	//		assert.Equal(t, NewSnake(5, 4), s)
	//	})
	//
	//	t.Run("Right", func(t *testing.T) {
	//		s := NewSnake(5, 5)
	//		s.Move(Right)
	//		assert.Equal(t, NewSnake(5, 6), s)
	//	})
}

//func TestEat(t *testing.T) {
//	t.Run("Up", func(t *testing.T) {
//		s := NewSnake(5, 5)
//		s.Eat(Up)
//		exp := NewSnake(5, 5)
//		exp.Body = append(exp.Body, Coordinates{6, 5})
//		assert.Equal(t, exp, s)
//	})
//}
