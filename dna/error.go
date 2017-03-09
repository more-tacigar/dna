package dna

import (
	"errors"
)

var (
	ErrInvalidPath          = errors.New("dna.node: invalid path")
	ErrInvalidParameterPath = errors.New("dna.node: invalid parameter path")
	ErrNotExistPath         = errors.New("dna.node: the path doesn't exist")
)

var (
	ErrInvalidMethod = errors.New("dna.Method: invalid method")
)

var (
	ErrJsonInvalidToken   = errors.New("dna.jsonLexer: invalid token")
	ErrJsonInvalidNumber  = errors.New("dna.jsonLexer: invalid number token")
	ErrJsonInvalidString  = errors.New("dna.jsonLexer: invalid string token")
	ErrJsonInvalidKeyword = errors.New("dna.jsonLexer: invalid keyword token")
)

var (
	ErrJsonInvalidValue   = errors.New("dna.jsonDecoder: invalid value")
	ErrJsonCannotDecode   = errors.New("dna.jsonDecoder: cannot decode to such value")
	ErrJsonCannotSetValue = errors.New("dna.jsonDecoder: unsettable value")
	ErrJsonInvalidSyntax  = errors.New("dna.jsonDecoder: invalid syntax")
)
