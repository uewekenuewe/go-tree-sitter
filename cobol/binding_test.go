package cobol_test

import (
	"context"
	"testing"
    "os"
	sitter "github.com/uewekenuewe/go-tree-sitter"
	"github.com/uewekenuewe/go-tree-sitter/cobol"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

    dat, err := os.ReadFile("./hiwo.cbl")

	n, err := sitter.ParseCtx(context.Background(), dat, cobol.GetLanguage())
	assert.NoError(err)
	assert.Equal(
"(start (program_definition (identification_division (program_name)) (procedure_division (display_statement (string)) (period) (stop_statement) (period))))",
		n.String(),
	)
}
