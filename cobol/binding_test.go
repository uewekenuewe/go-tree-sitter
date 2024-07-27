package cobol_test

import (
	"context"
	"testing"

	sitter "github.com/uewekenuwe/go-tree-sitter"
	"github.com/uewekenuwe/go-tree-sitter/cobol"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.ParseCtx(context.Background(), []byte("import cobol.io.*;"), java.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(program (import_declaration (scoped_identifier scope: (identifier) name: (identifier)) (asterisk)))",
		n.String(),
	)
}
