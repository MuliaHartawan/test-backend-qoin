package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLemparDadu(t *testing.T) {
	p := &Pemain{id: 1, dadu: make([]int, 4)}
	p.lemparDadu()

	for _, dadu := range p.dadu {
		assert.GreaterOrEqual(t, dadu, 1)
		assert.LessOrEqual(t, dadu, 6)
	}
}

func TestEvaluasiPemain(t *testing.T) {
	pemain1 := &Pemain{id: 1, dadu: []int{6, 1, 3}}
	pemain2 := &Pemain{id: 2, dadu: []int{2, 5, 1}}
	pemain3 := &Pemain{id: 3, dadu: []int{1, 6, 4}}

	pemain := []*Pemain{pemain1, pemain2, pemain3}
	evaluasiPemain(pemain)

	assert.Equal(t, 1, pemain1.skor)
	assert.Equal(t, 4, len(pemain1.dadu))

	assert.Equal(t, 2, len(pemain2.dadu))

	assert.Equal(t, 1, len(pemain3.dadu))
}

func TestBuatPemain(t *testing.T) {
	pemain := buatPemain(3, 2)

	assert.Equal(t, 3, len(pemain))
	for _, p := range pemain {
		assert.Equal(t, 2, len(p.dadu))
	}
}
