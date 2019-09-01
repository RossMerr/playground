grammar Euclid;

MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';

NUMBER : [0-9]+;

WS  : [ \t\r\n]+ -> skip ;

operation
    : left=NUMBER operator='+' right=NUMBER 
    | left=NUMBER operator='-' right=NUMBER
    | left=NUMBER operator='*' right=NUMBER
    | left=NUMBER operator='/' right=NUMBER
    ;
r   : 'hello' ID;
ID  : [a-z]+ ;


start : r | operation EOF;