package example

import (
	"fmt"
	"github.com/bonoken/watcher-common/common/strcase"
	"testing"
)

//  go test -run Case
func TestCase(t *testing.T) {
	s := "anyKind of_string"
	fmt.Println("ToSnake : ", strcase.ToSnake(s))
	fmt.Println("ToUpperSnake : ", strcase.ToUpperSnake(s))
	fmt.Println("ToKebab : ", strcase.ToKebab(s))
	fmt.Println("ToUpperKebab : ", strcase.ToUpperKebab(s))
	fmt.Println("ToDelimited : ", strcase.ToDelimited(s, '.'))
	fmt.Println("ToUpperDelimited : ", strcase.ToUpperDelimited(s, '.'))
	fmt.Println("ToPascal : ", strcase.ToPascal(s))
	fmt.Println("ToCamel : ", strcase.ToCamel(s))
}
