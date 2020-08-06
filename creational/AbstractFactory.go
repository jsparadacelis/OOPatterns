package main

import "fmt"

type IBook interface {
	myNameIs() string
}

type DramaBook struct {
	nameBook string
}

func (drama DramaBook) myNameIs() string {
	return drama.nameBook
}

type FantasyBook struct {
	nameBook string
}

func (fantasy FantasyBook) myNameIs() string {
	return fantasy.nameBook
}

type ILanguage interface {
	printBook() IBook
}

type Spanish struct {
}

func (spanish Spanish) printBook(topicBook string) IBook {
	switch topicBook {
	case "drama":
		fmt.Println("Creating a Drama Book...")
		dramaBook := DramaBook{"Book of drama in spanish"}
		return dramaBook
	case "fantasy":
		fmt.Println("Creating a Fantasy Book...")
		dramaBook := FantasyBook{"Book of fantasy in spanish"}
		return dramaBook
	default:
		fmt.Println("Creating a Fantasy Book...")
		dramaBook := DramaBook{"Book of drama in spanish"}
		return dramaBook
	}
}

type Portuguese struct {
}

func (portuguese Portuguese) printBook(topicBook string) IBook {
	switch topicBook {
	case "drama":
		fmt.Println("Creating a Drama Book...")
		dramaBook := DramaBook{"Book of drama ao portuguese"}
		return dramaBook
	case "fantasy":
		fmt.Println("Creating a Fantasy Book...")
		dramaBook := FantasyBook{"Book of fantasy ao portuguese"}
		return dramaBook
	default:
		fmt.Println("Creating a Fantasy Book...")
		dramaBook := DramaBook{"Book of drama ao portuguese"}
		return dramaBook
	}
}

type BookFactory struct {
}

func (bf BookFactory) createBook(bookLang string, bookTopic string) IBook {
	switch bookLang {
	case "spanish":
		return Spanish{}.printBook(bookTopic)
	case "portuguese":
		return Portuguese{}.printBook(bookTopic)
	default:
		return Portuguese{}.printBook(bookTopic)
	}
}
