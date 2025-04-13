package main

import (
	"fmt"
	"strings"
)

type Token struct {
	Type   string
	Lexeme string
	Line   int
	Column int
}

func Tokenizer(source string) []Token {
	tokens := []Token{}
	lines := strings.Split(source, "\n")

	for lineNum, line := range lines {
		columns := strings.Fields(line)
		for colNum, word := range columns {
			tokenType := "IDENTIFIER"

			switch word {
			case "Import":
				tokenType = "IMPORT"
			case "Body:":
				tokenType = "BODY"
			case "<SC>":
				tokenType = "SCRIPT_START"
			case "</SC>":
				tokenType = "SCRIPT_END"
			case "<ST>":
				tokenType = "STYLE_START"
			case "</ST>":
				tokenType = "STYLE_END"
			}

			token := Token{
				Type:   tokenType,
				Lexeme: word,
				Line:   lineNum + 1,
				Column: colNum + 1,
			}
			tokens = append(tokens, token)
		}
	}

	return tokens
}

func PrintTokens(tokens []Token) {
	for _, token := range tokens {
		fmt.Printf("Type: %s, Lexema: '%s', Line: %d, Column: %d\n",
			token.Type, token.Lexeme, token.Line, token.Column)
	}
}

func ParseTokens(tokens []Token) (string, error) {
	htmlCode := "<!DOCTYPE html>\n<html lang=\"en\">\n<meta charset='UTF-8'>\n<meta name='viewport' content='width=device-width, initial-scale=1.0'><head>\n</head>\n"
	bodyContent := ""
	openTags := []string{}
	currentTagContent := ""

	for _, token := range tokens {
		switch token.Type {
		case "IMPORT":

			continue
		case "BODY":
			bodyContent += "<body>\n"
		case "IDENTIFIER":

			if strings.HasPrefix(token.Lexeme, "|") {

				tag := strings.TrimPrefix(token.Lexeme, "|")

				if currentTagContent != "" {
					bodyContent += currentTagContent + "\n"
					currentTagContent = ""
				}

				if len(openTags) > 0 && openTags[len(openTags)-1] == strings.ToLower(tag) {
					bodyContent += fmt.Sprintf("</%s>\n", strings.ToLower(tag))
					openTags = openTags[:len(openTags)-1]
				}
			} else if strings.Contains(token.Lexeme, ":") {

				parts := strings.Split(strings.TrimSuffix(token.Lexeme, ":"), ".")
				tag := strings.ToLower(parts[0])
				class := ""
				src := ""
				if len(parts) > 1 {
					class = parts[1]
				}
				bodyContent += fmt.Sprintf("<%s", tag)
				if class != "" {
					bodyContent += fmt.Sprintf(" class=\"%s\"", class)
				}

				if src != "" {
					bodyContent += fmt.Sprintf(" src=\"%s\"", src)
				}
				bodyContent += ">\n"
				openTags = append(openTags, tag)
			} else if strings.ToUpper(token.Lexeme) == "H1" || strings.ToUpper(token.Lexeme) == "P" || strings.ToUpper(token.Lexeme) == "UL" || strings.ToUpper(token.Lexeme) == "H2" || strings.ToUpper(token.Lexeme) == "LI" || strings.ToUpper(token.Lexeme) == "H3" || strings.ToUpper(token.Lexeme) == "H4" || strings.ToUpper(token.Lexeme) == "H5" || strings.ToUpper(token.Lexeme) == "H6" || strings.ToUpper(token.Lexeme) == "SPAN" || strings.ToUpper(token.Lexeme) == "I" || strings.ToUpper(token.Lexeme) == "U" || strings.ToUpper(token.Lexeme) == "BR" || strings.ToUpper(token.Lexeme) == "HEADER" || strings.ToUpper(token.Lexeme) == "SECTION" || strings.ToUpper(token.Lexeme) == "FOOTER" || strings.ToUpper(token.Lexeme) == "TABLE" || strings.ToUpper(token.Lexeme) == "TR" || strings.ToUpper(token.Lexeme) == "TD" || strings.ToUpper(token.Lexeme) == "TH" || strings.ToUpper(token.Lexeme) == "THEAD" || strings.ToUpper(token.Lexeme) == "TBODY" || strings.ToUpper(token.Lexeme) == "TFOOT" || strings.ToUpper(token.Lexeme) == "CAPTION" || strings.ToUpper(token.Lexeme) == "FORM" || strings.ToUpper(token.Lexeme) == "INPUT" || strings.ToUpper(token.Lexeme) == "TEXTAREA" || strings.ToUpper(token.Lexeme) == "BUTTON" || strings.ToUpper(token.Lexeme) == "SELECT" || strings.ToUpper(token.Lexeme) == "OPTION" || strings.ToUpper(token.Lexeme) == "LABEL" || strings.ToUpper(token.Lexeme) == "MAIN" || strings.ToUpper(token.Lexeme) == "ARTICLE" || strings.ToUpper(token.Lexeme) == "NAV" {

				currentTagContent += fmt.Sprintf("<%s>%s</%s>\n", strings.ToLower(token.Lexeme), "", strings.ToLower(token.Lexeme))
				openTags = append(openTags, strings.ToLower(token.Lexeme))
			} else {

				currentTagContent += fmt.Sprintf("%s ", token.Lexeme)
			}
		default:

			currentTagContent += fmt.Sprintf("%s ", token.Lexeme)
		}
	}

	if currentTagContent != "" {
		bodyContent += currentTagContent + "\n"
	}
	for len(openTags) > 0 {
		lastTag := openTags[len(openTags)-1]
		bodyContent += fmt.Sprintf("</%s>\n", lastTag)
		openTags = openTags[:len(openTags)-1]
	}

	htmlCode += bodyContent + "</body>\n</html>"

	return htmlCode, nil
}
