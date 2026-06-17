package main

import "log"

type Song struct {
	Title  string
	Artist string
	Next   *Song
}

type Playlist struct {
	Head *Song
}

func main() {

	// song := &Song{
	// 	Title:  "Musica1",
	// 	Artist: "Artista1",
	// }
	// playlist := Playlist{Head: song}

	playlist := Playlist{}

	found := playlist.Find("One")
	if found != nil {
		log.Printf("%s encontrada", found.Title)
	} else {
		log.Printf("Não encontrada")
	}

	playlist.Print()
	log.Printf("Número de músicas: %d", playlist.Count())
	playlist.Add("Numb", "Linkin Park")
	playlist.Add("Crazy", "Aerosmith")
	playlist.Add("Hail to the king", "A7X")
	playlist.Add("Bohemian Rhapsody", "Queen")
	playlist.Add("One", "Metallica")
	playlist.Print()
	log.Printf("Número de músicas: %d", playlist.Count())

	found = playlist.Find("One")
	if found != nil {
		log.Printf("%s encontrada", found.Title)
	} else {
		log.Printf("Não encontrada")
	}

	log.Printf("Removeu? %t", playlist.Remove("Numb"))
	playlist.Print()

}

func (p *Playlist) Add(title, artist string) {

	newSong := &Song{
		Title:  title,
		Artist: artist,
	}

	if p.Head == nil {
		p.Head = newSong
		return
	}

	lastItem := p.Head
	for {
		if lastItem.Next != nil {
			lastItem = lastItem.Next
		} else {
			lastItem.Next = newSong
			return
		}
	}

}

func (p *Playlist) Print() {
	idx := p.Head
	if idx == nil {
		log.Println("Playlist vazia!!!")
		return
	}
	log.Println("Playlist:")
	for {
		log.Printf("%s - %s", idx.Title, idx.Artist)
		if idx.Next != nil {
			idx = idx.Next
		} else {
			log.Println("-----")
			return
		}
	}
}

func (p *Playlist) Count() int {
	count := 0
	idx := p.Head
	for {
		if idx != nil {
			count++
			if idx.Next != nil {
				idx = idx.Next
				continue
			}
		}
		return count
	}
}

func (p *Playlist) Find(title string) *Song {
	idx := p.Head

	for {
		if idx != nil {
			if idx.Title == title {
				return idx
			} else if idx.Next != nil {
				idx = idx.Next
				continue
			}
		}
		return nil
	}
}

func (p *Playlist) Remove(title string) bool {
	if p.Head == nil {
		return false
	} else if p.Head.Title == title {
		if p.Head.Next != nil {
			p.Head = p.Head.Next
		} else {
			p.Head = nil
		}
		return true
	}
	last := p.Head
	idx := p.Head.Next
	for {
		if idx != nil {
			if idx.Title == title {
				if idx.Next != nil {
					last.Next = idx.Next
				} else {
					last.Next = nil
				}
				return true
			}
			last = idx
			idx = idx.Next
			continue
		}
		return false
	}
}
