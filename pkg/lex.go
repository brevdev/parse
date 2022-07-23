package main


import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type itemType int

const (
	itemError itemType = iota
	itemHttps
	itemHttp
	itemHost
	itemOrg
	itemRepo
	itemGitUsername
	itemEOF
)

type item struct {
	typ itemType
	val string
}

func (i item) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemError:
		return i.val
	}
	return fmt.Sprintf("<%s>", i.val)
}

// a function that returns a statefn
type stateFn func(*lexer) stateFn

type lexer struct {
	name  string    // used in error reports
	input string    // string being scanned
	start int       // start position of this item
	pos   int       // current position of this item
	width int       // width of the last rune read
	items chan item // last scanned item
	state stateFn
}

func lex(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
		state: lexText,
		items: make(chan item, 2),
	}
	go l.run() // concurrently begin lexing
	return l
}

// synchronously receive an item from lexer
func (l *lexer) nextItem() item {
	return <-l.items
}

func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.items) // no more tokens will be delivered
}

func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

const eof = -1

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{
		itemError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

const (
	http           = "http://"
	https          = "https://"
	gitUsername    = "git@"
	githubHostname = "git@"
	semicolon      = ':'
	period         = '.'
	forwardslash   = '/'
)

func lexText(l *lexer) stateFn {
	i := 0
	for i < 1000 { //scan 1000 chars max to prevent infinite loops
		if strings.HasPrefix(l.input[l.pos:], http) {
			return lexHttp
		}
		if strings.HasPrefix(l.input[l.pos:], https) {
			return lexHttps
		}
		if strings.HasPrefix(l.input[l.pos:], gitUsername) {
			return lexGitUsername
		}
		if l.peek() == semicolon {
			return lexHost

		}
		if l.peek() == forwardslash {
			return lexHost

		}
		if l.next() == eof {
			break
		}
		i += 1
	}
	// if len(l.input[l.start:l.pos]) != 0 {
	// 	return l.errorf("unexpected eof")
	// }
	l.emit(itemEOF)
	return nil
}

func lexRepo(l *lexer) stateFn {
	if l.input[l.start] == semicolon || l.input[l.start] == forwardslash {
		l.start += 1
	}
	for {

		switch l.peek() {
		case period, forwardslash:
			l.emit(itemRepo)
			l.next()
			return lexText
		case eof:
			l.emit(itemRepo)
			l.next()
			l.emit(itemEOF)
			return nil

		}
		l.next()
	}

}

func lexOrg(l *lexer) stateFn {
	if l.input[l.start] == semicolon || l.input[l.start] == forwardslash {
		l.start += 1
	}
	for {
		if l.peek() == forwardslash {
			l.emit(itemOrg)
			l.next()
			return lexRepo

		}
		if l.next() == eof {

			return l.errorf("unexpected eof")
		}
	}

}

func lexHost(l *lexer) stateFn {
	l.emit(itemHost)
	l.next()
	return lexOrg

}
func lexHttp(l *lexer) stateFn {
	l.pos += len(http)
	l.emit(itemHttp)
	l.next()
	return lexText

}

func lexHttps(l *lexer) stateFn {
	l.pos += len(https)
	l.emit(itemHttps)
	l.next()
	return lexText

}

func lexGitUsername(l *lexer) stateFn {
	l.pos += len(gitUsername)
	l.emit(itemGitUsername)
	l.next()
	return lexText
}
