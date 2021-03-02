package css

import (
	"strings"

	"github.com/gorilla/css/scanner"
)

type CSSValue struct {
	Tokens []*scanner.Token
}

func NewCSSValue(csstext string) *CSSValue {
	sc := scanner.New(csstext)
	val := CSSValue{Tokens: make([]*scanner.Token, 0)}
Loop:
	for {
		token := sc.Next()
		switch token.Type {
		case scanner.TokenError, scanner.TokenEOF:
			break Loop
		default:
			val.Tokens = append(val.Tokens, token)
		}
	}
	return &val
}

func NewCSSValueString(data string) *CSSValue {
	data = strings.ReplaceAll(data, `"`, `\\"`)
	data = `"` + data + `"`
	token := scanner.Token{scanner.TokenString, data, 0, 0}
	return &CSSValue{Tokens: []*scanner.Token{&token}}
}

func (v *CSSValue) Text() string {
	var b strings.Builder
	for _, t := range v.Tokens {
		b.WriteString(t.Value)
	}
	return strings.TrimSpace(b.String())
}
