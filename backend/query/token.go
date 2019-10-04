package query

import "strings"

type TokenKind int

type Token struct {
	kind TokenKind
	str string
}

func NewToken(kind TokenKind, str string) *Token{
	return &Token{ kind : kind, str : str}
}

// The list of tokens
const(
	ILLEGAL TokenKind = iota
	EOF

	lit_begin
	STRING
	lit_end

	type_begin
	INT
	type_end

	operator_begin
	LBRACE
	RBRACE
	COMMA
	STAR
	operator_end

	keyword_begin
	SELECT
	FROM
	WHERE
	CREATE
	TABLE
	keyword_end
)

var tokens = [...]string{
	ILLEGAL:"ILLEGAL",
	EOF:"EOR",
	STRING:"STRING",
	INT:"INT",
	LBRACE:"{",
	RBRACE:"}",
	COMMA:",",
	SELECT:"Select",
	FROM:"From",
	WHERE:"Where",
	CREATE:"Create",
	TABLE:"Table",
}

func (t TokenKind) String() string{
	return tokens[t]
}

type tokenizer struct {
	input string
	pos int
}

func newTokenizer(input string) *tokenizer{
	return &tokenizer{
		input:input,
		pos:0,
	}
}

func (t *tokenizer) isSpace() bool{
	return t.input[t.pos] == ' ' || t.input[t.pos] == '\n' || t.input[t.pos] == '\t'
}


func (t *tokenizer) skipSpace(){
	for t.isSpace(){
		t.pos++
	}
}

func (t *tokenizer) isEnd() bool{
	return t.pos >= len(t.input)
}

func (t *tokenizer) matchKeyWord(keyword string) bool{
	ok := t.pos + len(keyword) <= len(t.input) &&
		strings.ToLower(t.input[t.pos:t.pos+len(keyword)]) == keyword

	if ok{
		t.pos += len(keyword)
	}
	return ok
}

func (t *tokenizer) isAsciiChar() bool{
	return (t.input[t.pos] >= 'a' && t.input[t.pos] <= 'z') ||
			(t.input[t.pos] >= 'A' && t.input[t.pos] <= 'Z')
}

func (t *tokenizer) scanString() string{
	var out []uint8
	for !t.isEnd() && !t.isSpace(){
		out = append(out, t.input[t.pos])
		t.pos++
	}
	return string(out)
}

func (t *tokenizer) Tokenize() ([]*Token, error){
	var tokens []*Token

	// compatible with ascii.
	for t.pos = 0; t.pos<len(t.input);{
		t.skipSpace()

		if t.matchKeyWord("create"){
			tokens = append(tokens, &Token{ kind : CREATE })
			continue
		}

		if t.matchKeyWord("table"){
			tokens = append(tokens, &Token{ kind : TABLE })
			continue
		}

		if t.matchKeyWord("int"){
			tokens = append(tokens, &Token{ kind : INT })
			continue
		}

		if t.matchKeyWord("select"){
			tokens = append(tokens, &Token{ kind : SELECT })
			continue
		}

		if t.matchKeyWord("from"){
			tokens = append(tokens, &Token{ kind : FROM })
			continue
		}

		if t.matchKeyWord("where"){
			tokens = append(tokens, &Token{ kind : WHERE })
			continue
		}

		if t.isAsciiChar(){
			ascii := t.scanString()
			tkn := NewToken(STRING, ascii)
			tokens = append(tokens, tkn)
			continue
		}

		switch t.input[t.pos] {
		case '{': tokens = append(tokens, &Token{ kind : LBRACE})
		case '}': tokens = append(tokens, &Token{ kind : RBRACE})
		case ',': tokens = append(tokens, &Token{ kind : COMMA})
		case '*': tokens = append(tokens, &Token{ kind : STAR})
		default:
			// error
		}

		t.pos++
	}

	return tokens, nil
}

func IsType(kind TokenKind) bool{
	return kind > type_begin && kind < type_end
}