package interpreter

import (
	"github.com/quantosnetwork/Quantosquantix/token"
	"io/ioutil"
	"unicode"
)

type state int

const nullState state = -1

type Lexer struct {
	I      []rune
	Tokens []*token.Token
}

func (l *Lexer) scan(i int) *token.Token {
	s, typ, rext := nullState, token.Error, i+1
	if i < len(l.I) {
		s = nextState[0](l.I[i])
	}
	for s != nullState {
		if rext >= len(l.I) {
			typ = accept[s]
			s = nullState
		} else {
			typ = accept[s]
			s = nextState[s](l.I[rext])
			if s != nullState || typ == token.Error {
				rext++
			}
		}
	}
	tok := token.NewToken(typ, i, rext, l.I)
	return tok
}

func InitNewLexer(_src []rune) *Lexer {
	lex := new(Lexer)
	lex.I = _src
	lex.Tokens = make([]*token.Token, 0, 2048)

	lext := 0
	for lext < len(lex.I) {
		for lext < len(lex.I) && unicode.IsSpace(lex.I[lext]) {
			lext++
		}
		if lext < len(lex.I) {
			tok := lex.scan(lext)
			lext = tok.Rext()
			if !tok.Suppress() {
			}
			lex.addToken(tok)
		}
	}

	lex.add(token.EOF, len(_src), len(_src))

	return lex
}

func NewInputFile(fname string) *Lexer {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	return InitNewLexer([]rune(string(buf)))
}

// util functions

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t'
}

func isLower(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func isUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func isLetter(r rune) bool {
	return isLower(r) || isUpper(r)
}

func isIdentStartChar(r rune) bool {
	return r == '_' || isLetter(r)
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isIdentChar(r rune) bool {
	return isIdentStartChar(r) || isDigit(r)
}

func escape(r rune) string {
	switch r {
	case '"':
		return "\""
	case '\\':
		return "\\\\"
	case '\r':
		return "\\r"
	case '\n':
		return "\\n"
	case '\t':
		return "\\t"
	}
	return string(r)
}

// GetLineColumn returns the line and column of rune[i] in the input
func (l *Lexer) GetLineColumn(i int) (line, col int) {
	line, col = 1, 1
	for j := 0; j < i; j++ {
		switch l.I[j] {
		case '\n':
			line++
			col = 1
		case '\t':
			col += 4
		default:
			col++
		}
	}
	return
}

// GetLineColumnOfToken returns the line and column of token[i] in the imput
func (l *Lexer) GetLineColumnOfToken(i int) (line, col int) {
	return l.GetLineColumn(l.Tokens[i].Lext())
}

// GetString returns the input string from the left extent of Token[lext] to
// the right extent of Token[rext]
func (l *Lexer) GetString(lext, rext int) string {
	return string(l.I[l.Tokens[lext].Lext():l.Tokens[rext].Rext()])
}

func (l *Lexer) add(t token.Type, lext, rext int) {
	l.addToken(token.NewToken(token.TokenType(t), lext, rext, l.I))
}

func (l *Lexer) addToken(tok *token.Token) {
	l.Tokens = append(l.Tokens, tok)
}

func any(r rune, set []rune) bool {
	for _, r1 := range set {
		if r == r1 {
			return true
		}
	}
	return false
}

func not(r rune, set []rune) bool {
	for _, r1 := range set {
		if r == r1 {
			return false
		}
	}
	return true
}

var accept = []token.Type{
	token.Error,
	token.T_0,
	token.Error,
	token.Error,
	token.T_1,
	token.T_2,
	token.T_3,
	token.T_4,
	token.T_5,
	token.T_6,
	token.T_7,
	token.T_8,
	token.T_9,
	token.Error,
	token.Error,
	token.Error,
	token.Error,
	token.Error,
	token.Error,
	token.T_22,
	token.T_23,
	token.T_24,
	token.Error,
	token.Error,
	token.Error,
	token.Error,
	token.Error,
	token.Error,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_16,
	token.T_19,
	token.T_11,
	token.T_11,
	token.T_10,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_15,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_12,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_20,
	token.T_13,
	token.T_20,
	token.T_17,
	token.T_20,
	token.T_21,
	token.T_14,
	token.T_18,
}

var nextState = []func(r rune) state{
	// Set0
	func(r rune) state {
		switch {
		case r == '!':
			return 1
		case r == '"':
			return 2
		case r == '\'':
			return 3
		case r == '(':
			return 4
		case r == ')':
			return 5
		case r == '.':
			return 6
		case r == ':':
			return 7
		case r == ';':
			return 8
		case r == '<':
			return 9
		case r == '>':
			return 10
		case r == '[':
			return 11
		case r == ']':
			return 12
		case r == 'a':
			return 13
		case r == 'e':
			return 14
		case r == 'l':
			return 15
		case r == 'n':
			return 16
		case r == 'p':
			return 17
		case r == 'u':
			return 18
		case r == '{':
			return 19
		case r == '|':
			return 20
		case r == '}':
			return 21
		case unicode.IsUpper(r):
			return 22
		case unicode.IsLower(r):
			return 23
		}
		return nullState
	},
	// Set1
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set2
	func(r rune) state {
		switch {
		case r == '\\':
			return 24
		case not(r, []rune{'"', '\\'}):
			return 25
		}
		return nullState
	},
	// Set3
	func(r rune) state {
		switch {
		case r == '\\':
			return 26
		case not(r, []rune{'\''}):
			return 27
		}
		return nullState
	},
	// Set4
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set5
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set6
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set7
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set8
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set9
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set10
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set11
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set12
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set13
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'n':
			return 29
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set14
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'm':
			return 30
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set15
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'e':
			return 31
		case r == 'o':
			return 32
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set16
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'o':
			return 33
		case r == 'u':
			return 34
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set17
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'a':
			return 35
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set18
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'p':
			return 36
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set19
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set20
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set21
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set22
	func(r rune) state {
		switch {
		case r == '_':
			return 37
		case unicode.IsLetter(r):
			return 37
		case unicode.IsNumber(r):
			return 37
		}
		return nullState
	},
	// Set23
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set24
	func(r rune) state {
		switch {
		case any(r, []rune{'"', '\\', 'n', 'r', 't'}):
			return 25
		}
		return nullState
	},
	// Set25
	func(r rune) state {
		switch {
		case r == '"':
			return 38
		case r == '\\':
			return 24
		case not(r, []rune{'"', '\\'}):
			return 25
		}
		return nullState
	},
	// Set26
	func(r rune) state {
		switch {
		case any(r, []rune{'\'', '\\', 'n', 'r', 't'}):
			return 39
		case r == '\'':
			return 39
		}
		return nullState
	},
	// Set27
	func(r rune) state {
		switch {
		case r == '\'':
			return 40
		}
		return nullState
	},
	// Set28
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set29
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'y':
			return 41
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set30
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'p':
			return 42
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set31
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 't':
			return 43
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set32
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'w':
			return 44
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set33
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 't':
			return 45
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set34
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'm':
			return 46
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set35
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'c':
			return 47
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set36
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'c':
			return 48
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set37
	func(r rune) state {
		switch {
		case r == '_':
			return 37
		case unicode.IsLetter(r):
			return 37
		case unicode.IsNumber(r):
			return 37
		}
		return nullState
	},
	// Set38
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set39
	func(r rune) state {
		switch {
		case r == '\'':
			return 40
		}
		return nullState
	},
	// Set40
	func(r rune) state {
		switch {
		}
		return nullState
	},
	// Set41
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set42
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 't':
			return 49
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set43
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 't':
			return 50
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set44
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'c':
			return 51
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set45
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set46
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'b':
			return 52
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set47
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'k':
			return 53
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set48
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'a':
			return 54
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set49
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'y':
			return 55
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set50
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'e':
			return 56
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set51
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'a':
			return 57
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set52
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'e':
			return 58
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set53
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'a':
			return 59
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set54
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 's':
			return 60
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set55
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set56
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'r':
			return 61
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set57
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 's':
			return 62
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set58
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'r':
			return 63
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set59
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'g':
			return 64
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set60
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'e':
			return 65
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set61
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set62
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'e':
			return 66
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set63
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set64
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case r == 'e':
			return 67
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set65
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set66
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
	// Set67
	func(r rune) state {
		switch {
		case r == '_':
			return 28
		case unicode.IsLetter(r):
			return 28
		case unicode.IsNumber(r):
			return 28
		}
		return nullState
	},
}
