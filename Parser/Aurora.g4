grammar Aurora;

program
    : expressionList EOF
    ;

expressionList
    : expr0 expr0*
    ;

/* Precedence 0 (immidiate values) */

expr0
    : stringImmidiate
    | integerImmidiate
    | variableImmidiate
    | expr1
    ;

stringImmidiate
    : StringVal=StringLiteral expr0?
    ;

variableImmidiate
    : Name=Identifier expr0?
    ;

integerImmidiate
    : NumberVal=IntegerLiteral expr0?
    ;

/* Precedence 1 */

expr1
    : functionCall
    | memberAccess
    ;

memberAccess
    : '.' Member=Identifier expr1?
    ;

functionCall
    : '(' functionCallParam? (',' functionCallParam)* ')' expr1?
    ;

functionCallParam
    : ParamExpr=expr0
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

        -> skip
    ;

BlockComment
    : '/*' .*? '*/'
        -> skip
    ;

LineComment
    : '//' ~[\r\n]*
        -> skip
    ;