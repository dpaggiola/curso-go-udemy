package book

import "fmt"

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	title  string
	author string
	pages  int
}

func NewBook(title, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\n Author: %s\n Pages: %d\n", b.title, b.author, b.pages)
}

type TextBook struct {
	Book
	editorial string
	level string
}

func NewTextBook(title, author, editorial, level string, pages int) *TextBook {
	return &TextBook{
		Book: Book{
			title:  title,
			author: author,
			pages:  pages,
		},
		editorial: editorial,
		level: level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\n Author: %s\n Pages: %d\n Editorial: %s\n Nivel: %s\n", b.title, b.author, b.pages, b.editorial, b.level)
}