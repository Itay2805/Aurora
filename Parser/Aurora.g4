grammar Aurora;

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
    : expr7
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


/* Precedence 2 */

expr2
    :  
    ( lhsOperator
    )
    | expr1
    ;

lhsOperator
    :
        Op=( '-'
        | '+'
        | '!'
        | '*'
        | '&'
        )
        On=expr
    ;

/* Precedence 3 */

expr3
    : Left=expr3 mulDivMod
    | expr2
    ;

mulDivMod
    :   Op=( '*'
        | '/'
        | '%'
        )
        Right=expr3
    ;

/* Precedence 4 */

expr4
    : Left=expr4 addSub
    | expr3
    ;

addSub
    :   Op=( '-'
        | '+'
        )
        Right=expr4
    ;

/* Precedence 5 */

expr5
    : Left=expr5 logicalAnd
    | expr4
    ;

logicalAnd
    :   '&&' Right=expr5
    ;

/* Precedence 6 */

expr6
    : Left=expr6 logicalOr
    | expr5
    ;

logicalOr
    :   '||' Right=expr6
    ;

/* Precedence 7 */
expr7
    :(
        ( memberAccess
        | identifierImmidiate
        )
        assign
    )
    | expr6
    ;

assign
    :
    '='
    Right=expr
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