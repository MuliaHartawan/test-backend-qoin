package main

import (
	"fmt"
	"math/rand"
)

type Pemain struct {
	id   int
	dadu []int
	skor int
}

func buatPemain(jumlahPemain, jumlahDadu int) []*Pemain {
	pemain := make([]*Pemain, jumlahPemain)
	for i := 0; i < jumlahPemain; i++ {
		pemain[i] = &Pemain{id: i + 1, dadu: make([]int, jumlahDadu)}
	}
	return pemain
}

func (p *Pemain) lemparDadu() {
	for i := range p.dadu {
		p.dadu[i] = rand.Intn(6) + 1
	}
}

func evaluasiPemain(pemain []*Pemain) {
	for i := range pemain {
		if len(pemain[i].dadu) == 0 {
			continue
		}

		daduBaru := []int{}
		for _, nilai := range pemain[i].dadu {
			switch nilai {
			case 6:
				pemain[i].skor++
			case 1:
				indeksBerikutnya := (i + 1) % len(pemain)
				pemain[indeksBerikutnya].dadu = append(pemain[indeksBerikutnya].dadu, 1)
			default:
				daduBaru = append(daduBaru, nilai)
			}
		}
		pemain[i].dadu = daduBaru
	}
}

func tampilkanHasil(pemain []*Pemain, giliran int) {
	fmt.Printf("==================\nGiliran %d lempar dadu:\n", giliran)
	for _, p := range pemain {
		if len(p.dadu) == 0 {
			fmt.Printf("Pemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu)\n", p.id, p.skor)
			continue
		}
		p.lemparDadu()
		fmt.Printf("Pemain #%d (%d): %v\n", p.id, p.skor, p.dadu)
	}
}

func tampilkanEvaluasi(pemain []*Pemain) {
	fmt.Println("Setelah evaluasi:")
	pemainAktif := 0
	for _, p := range pemain {
		if len(p.dadu) == 0 {
			fmt.Printf("Pemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu)\n", p.id, p.skor)
		} else {
			pemainAktif++
			fmt.Printf("Pemain #%d (%d): %v\n", p.id, p.skor, p.dadu)
		}
	}
}

func main() {
	var N, M int

	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&N)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&M)

	pemain := buatPemain(N, M)

	giliran := 1
	for {
		tampilkanHasil(pemain, giliran)
		evaluasiPemain(pemain)
		tampilkanEvaluasi(pemain)

		pemainAktif := 0
		for _, p := range pemain {
			if len(p.dadu) > 0 {
				pemainAktif++
			}
		}
		if pemainAktif <= 1 {
			break
		}
		giliran++
	}

	var pemenang *Pemain
	for _, p := range pemain {
		if len(p.dadu) > 0 && (pemenang == nil || p.skor > pemenang.skor) {
			pemenang = p
		}
	}

	fmt.Printf("==================\nGame berakhir karena hanya pemain #%d yang memiliki dadu.\n", pemenang.id)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", pemenang.id)
}
