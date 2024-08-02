lexer grammar CalcLexer;

ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';
INT: [0-9]+;
WS: [ \t\r\n]+ -> skip;