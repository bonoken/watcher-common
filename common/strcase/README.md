# strcase
[![Godoc Reference](https://godoc.org/github.com/iancoleman/strcase?status.svg)](http://godoc.org/github.com/iancoleman/strcase)
[![Build Status](https://travis-ci.org/iancoleman/strcase.svg)](https://travis-ci.org/iancoleman/strcase)
[![Coverage](http://gocover.io/_badge/github.com/iancoleman/strcase?0)](http://gocover.io/github.com/iancoleman/strcase)

strcase is a go package for converting string case to [snake case](https://en.wikipedia.org/wiki/Snake_case) or [camel case](https://en.wikipedia.org/wiki/CamelCase).

## Example

```go
s := "AnyKind of_string"
```

| Function                          | Result               |
|-----------------------------------|----------------------|
| `ToSnake(s)`                      | `any_kind_of_string` |
| `ToUpperSnake(s)`                 | `ANY_KIND_OF_STRING` |
| `ToKebab(s)`                      | `any-kind-of-string` |
| `ToUpperKebab(s)`                 | `ANY-KIND-OF-STRING` |
| `ToDelimited(s, '.')`             | `any.kind.of.string` |
| `ToUpperDelimited(s, '.')`        | `ANY.KIND.OF.STRING` |
| `ToPascal(s)`                     | `AnyKindOfString`    |
| `ToCamel(s)`                      | `anyKindOfString`    |