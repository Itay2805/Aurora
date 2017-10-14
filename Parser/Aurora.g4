grammar Aurora;

@lexer::members {
var IdSeen bool = false
 }

program
    : expressionList? EOF?
    ;

expressionList
    : expression expression*
    ;

expression
    : expr (Newline | ';' | EOF)
    ;

expr
    : expr1
    ;


/* Precedence 0 (immidiate values) */

expr0
    :   ( brackets
        | stringImmidiate
        | integerImmidiate
        | identifierImmidiate
        )
    ;

brackets
    : '(' ParamExpr=expr ')'
    ;

stringImmidiate
    : StringVal=StringLiteral
    ;

identifierImmidiate
    : Name=Identifier
    ;

integerImmidiate
    : NumberVal=IntegerLiteral
    ;

/* Precedence 1 */

expr1
    : expr1 
    ( memberAccess
    | functionCall
    )
    | expr0
    ;

memberAccess
    : '.' Member=Identifier
    ;

functionCall
    : '(' functionCallParam? (',' functionCallParam)* ')'
    ;

functionCallParam
    : ParamExpr=expr
    ;

/* Literals */

IntegerLiteral
    : (Digit | '_')+
    ;

StringLiteral
    : '"' ~["]* '"'
    ;

Identifier
    : IdentifierNondigit
        (
            IdentifierNondigit
            | Digit
        )*
    ;

fragment IdentifierNondigit
    : Nondigit
    ;

fragment Nondigit
    : [a-zA-Z_]
    ;

fragment Digit
    : [0-9]
    ;

Whitespace
    : [ \t]+
        -> skip
    ;

Newline
    :   ( '\r' '\n'? 
        | '\n'
        )
    ;

BlockComment
    : '/*' .*? '*/'
        -> skip
    ;

LineComment
    : '//' ~[\r\n]*
        -> skip
    ;