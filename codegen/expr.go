package codegen

//import "astar"

type Expr interface {
   Equiv(e Expr) bool
}

// Operands are identified by a unique integer
type Operand int

func (me *Operand) Equiv(e Expr) bool {
   op, ok := e.(*Operand)
   if ok {
      return int(*op) == int(*me)
   } else {
      return false
   }
}

type Mul struct {
   left, right Expr
}

func (me *Mul) Equiv(e Expr) bool {
   mul, ok := e.(*Mul)
   if ok {
      a := mul.left.Equiv(me.right) && mul.right.Equiv(me.left)
      b := mul.left.Equiv(me.left) && mul.right.Equiv(me.right)

      return a || b
   } else {
      return false
   }

   panic("bad return path error (x_X)")
}

type Add struct {
   left, right Expr
}

func (me *Add) Equiv(e Expr) bool {
   add, ok := e.(*Add)
   if ok {
      a := add.left.Equiv(me.right) && add.right.Equiv(me.left)
      b := add.left.Equiv(me.left) && add.right.Equiv(me.right)

      return a || b
   } else {
      return false
   }

   panic("bad return path error (x_X)")
}
