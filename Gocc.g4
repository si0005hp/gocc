grammar Gocc;

@parser::header {
import "github.com/si0005hp/gocc/ast"
}

expr returns [ast.ExprNode n]
  : l=expr op=('*'|'/') r=expr      { $n = &ast.BinOpNode{$op.type, $l.n, $r.n} }
	| l=expr op=('+'|'-') r=expr      { $n = &ast.BinOpNode{$op.type, $l.n, $l.n} }
  | INT                             { $n = &ast.IntNode{$INT.int} }
  ;

INT: [0-9]+ ;
ADD: '+' ;
SUB: '-' ;
MUL: '*' ;
DIV: '/' ;

NEWLINE : ('\r' '\n'?|'\n') -> skip ;
WS : [ \t]+ -> skip ;