package lexer

import "slices"
//import "fmt"
import "strings"
import "errors"

// Lexer Token, with cat being the type of token and value being the data
type Token struct {
	cat string // The Catagory the token is in E.g open group or atom
	value string // The value the token holds
}

func NewToken(buff string) (Token,error) {
    buff = strings.TrimSpace(buff)
    if buff == "(" {
        return Token{cat: "open_group",value: buff }, nil
    } else if buff == ")" {
		return Token{cat: "close_group", value: buff}, nil
    } else if buff=="" {
		return Token{cat: "err", value: "err"},  errors.New("EmptyBuff")
    } else {
        return Token{cat: "atom", value: buff}, nil
    }
}

func Lex(text string) []Token {
	closers := []string {"(", ")", " "}
	final := make([]Token, 0)
	var buff string  = ""
	for i := 0; i<len(text); i++ {
		buff += string(text[i])

		if slices.Contains(closers,string(text[i])) {
    		buff = buff[:len(buff)-1]
    		tok1, err := NewToken(buff)
    		if err == nil {
				final = append(final,tok1)
				buff = ""
    		}
    		if (string(text[i])) != " " {

        		buff = buff + string(text[i])
        		tok2, _ := NewToken(buff)
				final = append(final, tok2) 
    		}
    		buff = ""
		}
	}
	return final
}
