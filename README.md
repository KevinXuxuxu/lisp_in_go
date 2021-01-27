# Lisp-in-Go
> Golang powered Lisp intepreter

### EBNF (WIP)
```
SourceCharacter ::=  #x0009 | #x000A | #x000D | [#x0020-#xFFFF]
StringCharacter ::= SourceCharacter - `"`
StringContent ::= StringCharacter | StringCharacter StringContent
Ignored ::= [\t\n\v\f\r ]+
Symbol ::= [^0-9()'"\t\n\v\f\r ][^()'"\t\n\v\f\r ]+
Int ::= [0-9]+
String ::= " " | " StringContent "
List ::= ' FunctionCall
Element ::= Symbol | Int | String | List | FunctionCall
Elements ::= Element Ignored | Element Ignored Elements Ignored
FunctionCall ::= ( Ignored Elements )
```
