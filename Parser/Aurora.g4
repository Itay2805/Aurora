grammar Aurora;

program
    : declarationList? EOF?
    ;

declarationList
    : declaration declaration*
    ;

declaration
    : variableDeclaration
    | functionDeclaration
    ;

statementList
    : stmt stmt*
    ;

codeBlock
    : '{' statementList? '}'
    ;

variableDeclaration
    : variableStmt
    ;

variableStmt
    : 'var' export='export'? optional='optional'?  name=Identifier ':' (variableType=Identifier ptr=pointer? gcptr='^'?)? ('=' expr)?
    ;

functionDeclaration
    : 'func' export='export'? native='native'? name=Identifier '(' functionParameterList ')' code=codeBlock
    ;

functionParameter
    : parameterName=Identifier ':' parameterType=Identifier (ptr=pointer? gcptr='^'?) | ref='&'
    ;

functionParameterList
    : functionParameter (',' functionParameter)*
    ;

pointer
    : '*'+
    ;

stmt
    :   (
            ( variableStmt
            | expressionStmt
            )
            (Newline | ';' | EOF)

        | codeBlock
        )
    ;

expressionStmt
    : expr
    ;

/* Precedence 0 (immidiate values) */

expr
    : expr7
    ;

expr0
    :   ( brackets
        | stringImmidiate
        | integerImmidiate
        | identifierImmidiate
        )
    ;

brackets
    : '(' paramExpr=expr ')'
    ;

stringImmidiate
    : stringVal=StringLiteral
    ;

identifierImmidiate
    : name=Identifier
    ;

integerImmidiate
    : numberVal=IntegerLiteral
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
    : '.' member=Identifier
    ;

functionCall
    : '(' functionCallParam? (',' functionCallParam)* ')'
    ;

functionCallParam
    : paramExpr=expr
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
        op=( '-'
        | '+'
        | '!'
        | '*'
        | '&'
        )
        on=expr
    ;

/* Precedence 3 */

expr3
    : left=expr3 mulDivMod
    | expr2
    ;

mulDivMod
    :   op=( '*'
        | '/'
        | '%'
        )
        right=expr3
    ;

/* Precedence 4 */

expr4
    : left=expr4 addSub
    | expr3
    ;

addSub
    :   op=( '-'
        | '+'
        )
        right=expr4
    ;

/* Precedence 5 */

expr5
    : left=expr5 logicalAnd
    | expr4
    ;

logicalAnd
    :   '&&' right=expr5
    ;

/* Precedence 6 */

expr6
    : left=expr6 logicalOr
    | expr5
    ;

logicalOr
    :   '||' right=expr6
    ;

/* Precedence 7 */
expr7
    :(
        ( expr1 memberAccess
        | identifierImmidiate
        )
        assign
    )
    | expr6
    ;

assign
    :
    '='
    right=expr
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