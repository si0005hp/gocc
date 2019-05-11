grammar Gocc;

@parser::header {
import "github.com/si0005hp/gocc/ast"
}

// expr returns [Expr n]
//   : l=expr op=('*'|'/'|'%') r=expr
// 	| l=expr op=('+'|'-') r=expr    
//   | INT                             { $n = IntNode{$INT.int} }
//   ;

expr returns [ast.Expr n]
  // : l=expr op=('*'|'/'|'%') r=expr
	// | l=expr op=('+'|'-') r=expr    
  : INT                             { $n = &ast.IntNode{$INT.int} }
  ;

INT: [0-9]+ ;
ADD: '+' ;
SUB: '-' ;
MUL: '*' ;
DIV: '/' ;