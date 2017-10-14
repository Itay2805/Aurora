grammar Aurora;

program
    : expressionList EOF?
    ;

expressionList
    : expression expression*
    ;

expression
    : expr0 (Newline | ';' | EOF)
    ;

/* Precedence 0 (immidiate values) */

expr0
    : 
        ( brackets
        | stringImmidiate
        | integerImmidiate
        | identifierImmidiate
        | expr1
        )

        expr0?
    ;

brackets
    : '(' expr0 ')'
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
    : 
        (functionCall
        | memberAccess
        | expr2
        )
        expr1?
    ;

memberAccess
    : '.' Member=Identifier
    ;

functionCall
    : '(' functionCallParam? (',' functionCallParam)* ')'
    ;

functionCallParam
    : ParamExpr=expr0
    ;

/* Precedence 2 */

expr2
    : 
        Op=( '-'
        | '+'
        | '!'
        | '*'
        | '&'
        )
        Expr=expr0
        expr2?
    ;

/* Literals */

IntegerLiteral
    : (Digit | '_')+
    ;

StringLiteral
    : '"' [^"]* '"'
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