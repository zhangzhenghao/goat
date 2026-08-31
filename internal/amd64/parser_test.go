package amd64

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gorse-io/goat/internal"
)

func TestGenerateGoAssemblyIncludesFloatResultInArgumentSize(t *testing.T) {
	path := filepath.Join(t.TempDir(), "result.s")
	functions := []internal.Function{
		{
			Name: "distance",
			Type: "float",
			Parameters: []internal.Parameter{
				{Name: "lhs", ParameterType: internal.ParameterType{Pointer: true}},
				{Name: "rhs", ParameterType: internal.ParameterType{Pointer: true}},
				{Name: "size", ParameterType: internal.ParameterType{Type: "int64_t"}},
			},
			Lines: []internal.Line{{Assembly: "retq"}},
		},
	}

	if err := generateGoAssembly("", "", path, functions); err != nil {
		t.Fatal(err)
	}
	generated, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(generated), "TEXT ·distance(SB), $8-28") {
		t.Fatalf("float result missing from argument size:\n%s", generated)
	}
}
