grammar Aurora;

program
    : expr_list EOF
    ;

functionDeclaration
    : 'func' Name=Identifier 
        '(' 
        
        functionParamList?
        
        ')' 
        
        (
            ':'
            ReturnType=Identifier
        )?

        codeBlock?
    ;

functionParam
    : 
        ParamName=Identifier 
        ':' 
        ParamType=Identifier
    ;

functionParamList
    : functionParam
    | functionParamList ',' functionParam
    ;

codeBlock
    : '{' expr_list '}'
    ;

expr_list
    : expr*
    ;

expr
    : codeBlock
    | functionDeclaration
    |
        ( 
        )
        ';'
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