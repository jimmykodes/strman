# String Manipulation

Golang string manipulation package

## `ToDelimited`

Split the provided string into its parts and join them with the provided delimiter

```golang
strman.ToDelimited("fromCamel", ".") // from.camel
strman.ToDelimited("FromPascal", ".") // from.pascal
strman.ToDelimited("from-kebab", ".") // from.kebab
strman.ToDelimited("from_snake", ".") // from.snake
strman.ToDelimited("from-superMixed_case", ".") // from.super.mixed.case
```

## `ToScreamingDelimited`

Split the provided string into its parts, make them all uppercase, and join them with the provided delimiter

```golang
strman.ToDelimited("fromCamel", ".") // FROM.CAMEL
strman.ToDelimited("FromPascal", ".") // FROM.PASCAL
strman.ToDelimited("from-kebab", ".") // FROM.KEBAB
strman.ToDelimited("from_snake", ".") // FROM.SNAKE
strman.ToDelimited("from-superMixed_case", ".") // FROM.SUPER.MIXED.CASE
```

## `ToSnake`

Run `ToDelimited` with `"_"` as the delimiter

```golang
strman.ToSnake("fromCamel") // from_camel
strman.ToSnake("from-kebab") // from_kebab
```

## `ToScreamingSnake`

Run `ToScreamingDelimited` with `"_"` as the delimiter

```golang
strman.ToScreamingSnake("fromCamel") // FROM_CAMEL
strman.ToScreamingSnake("from-kebab") // FROM_KEBAB
```

## `ToKebab`

Run `ToDelimited` with `"-"` as the delimiter

```golang
strman.ToKebab("fromCamel") // from-camel
strman.ToKebab("from_snake") // from-snake
```

## `ToScreamingKebab`

Run `ToScreamingDelimited` with `"-"` as the delimiter

```golang
strman.ToScreamingKebab("fromCamel") // FROM-CAMEL
strman.ToScreamingKebab("from_snake") // FROM-SNAKE
```

## `ToPascal`

Split the provided string into its parts and join them by capitalizing the first letter of each word

```golang
strman.ToPascal("fromCamel") // FromCamel
strman.ToPascal("from-kebab") // FromKebab
strman.ToPascal("from_snake") // FromSnake
```

## `ToCamel`

Like `ToPascal` but the first letter remains lowercase

```golang
strman.ToCamel("FromPascal") // fromPascal
strman.ToCamel("from-kebab") // fromKebab
strman.ToCamel("from_snake") // fromSnake
```

