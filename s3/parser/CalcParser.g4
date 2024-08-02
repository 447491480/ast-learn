parser grammar CalcParser;

options { tokenVocab=CalcLexer; }

expression
    : expression op=('*'|'/') expression
    | expression op=('+'|'-') expression
    | INT
    ;