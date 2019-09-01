grammar Euclid;

// IMPORT: 'import';
// DIGIT: [d]+;
// WORD: [w]+;

// importR : IMPORT;

r   : 'hello' ID;
ID  : [a-z]+ ;
WS  : [ \t\r\n]+ -> skip ;