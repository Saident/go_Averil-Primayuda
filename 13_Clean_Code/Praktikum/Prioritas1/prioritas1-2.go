package main

type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (Mobil *Mobil) Berjalan() {
	KecepatanBaru := 10
	Mobil.tambahKecepatan(KecepatanBaru)
}

func (Mobil *Mobil) tambahKecepatan(KecepatanBaru int) {
	Mobil.KecepatanPerJam += KecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	mobilLamban := Mobil{}
	mobilLamban.Berjalan()
}
