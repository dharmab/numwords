# numwords <br/>

**numwords** is a package for Go that converts natural language numbers
to their actual numeric values. The numbers can be parsed out as strings,
integers, or floats as desired.

```go
func Example() {
  s := "I've got three apples and two and a half bananas"
  fmt.Println(ParseString(s))

  s = "My chili won second place at the county fair"
  fmt.Println(ParseString(s))

  i, _ := ParseInt("fourteen ninety two")
  fmt.Println(i)

  f, _ := ParseFloat("eight and three quarters")
  fmt.Println(f)

  // Output:
  // I've got 3 apples and 2.5 bananas
  // My chili won 2nd place at the county fair
  // 1492
  // 8.75
}
```

This package is a fork of [Chris Roche's original](https://github.com/rodaine/numwords), which was inspired by the [Numerizer Ruby gem](https://github.com/jduff/numerizer).

New features in my version include:

- `IncludeIndefiniteArticle(bool)` — toggle parsing of indefinite articles ("a", "an") as `1`
- `IncludeFractions(bool)` — toggle parsing of fraction words ("half", "thirds", "quarter", etc.)
- Modernized to Go 1.26 with golangci-lint, a Makefile, and updated GitHub Actions CI

## Some Valid Conversions

| String Input | Output Value |
| :------ | ----: |
| fifteen | 15 |
| twenty five | 25 |
| twenty-five | 25 |
| eleven hundred | 1100 |
| three hundred twenty five | 325 |
| three hundred thousand | 300000 |
| one hundred twenty one | 121 |
| fourteen hundred sixty seven | 1467 |
| nineteen eighty-eight | 1988 |
| nine hundred and ninety nine | 999 |
| a half | 0.5 |
| three halves | 1.5 |
| a quarter | 0.25 |
| three quarters | 0.75 |
| one ninth | 0.111111 |
| two thirds | 0.666667 |
| two and three eighths | 2.375 |
| zeroth | 0th |
| twenty second | 22nd |
| 5 hundred | 500 |
| one million two hundred fifty thousand and seven | 1250007 |

## License

This package is released under the MIT [License](https://github.com/dharmab/numwords/blob/master/LICENSE).
