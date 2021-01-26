# Lisp-in-Go
> Golang powered Lisp intepreter

### EBNF
```
SourceCharacter ::=  #x0009 | #x000A | #x000D | [#x0020-#xFFFF]
StringCharacter ::= SourceCharacter - `"`
StringContent ::= StringCharacter | StringCharacter StringContent
Ignored ::= [\t\n\v\f\r ]+
Symbol ::= [^0-9()'"\t\n\v\f\r ][^()'"\t\n\v\f\r ]+
Int ::= [0-9]+
String ::= " " Ignored | " StringContent " Ignored
List ::= ' FunctionCall
Element ::= Int | String | List | FunctionCall
Elements ::= Element | Element Elements
FunctionCall ::= ( Elements )
```
