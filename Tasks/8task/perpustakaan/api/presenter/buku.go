// presenter/buku_presenter.go
package presenter

import "perpustakaan_app/pkg/entities"

// Struct khusus untuk response JSON
type BukuResponse struct {
	ID      uint   `json:"id"`
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
	Tahun   int    `json:"tahun"`
}

// Fungsi untuk mapping entity ke response
func BukuToResponse(b entities.Buku) BukuResponse {
	return BukuResponse{
		ID:      b.ID,
		Judul:   b.Judul,
		Penulis: b.Penulis,
		Tahun:   b.Tahun,
	}
}

func ListBukuToResponse(bukus []entities.Buku) []BukuResponse {
	responses := make([]BukuResponse, len(bukus))
	for i, b := range bukus {
		responses[i] = BukuToResponse(b)
	}
	return responses
}
