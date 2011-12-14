package codegen

type Operand []int

type Production struct {
   // The number of operations that preceded this one ( i.e. parent.depth + 1 )
   Depth  int
   Parent *Production
   Result Operand
   Op     *Operation
   prob   *problemInstance

   // Depth of the operands that were passed to the operation to get this
   A, B int
}

type problemInstance struct {
   goal *Operand
   db   *ExprDB
}

type Operation interface {
   // mnemonic for this operation
   String() string

   // Returns all Productions that could come after 'pred' using this
   // operation. They must use 'pred.Result' as an operand.
   Children(pred *Production) []Production
}

func (me *Production) IsGoal() bool {
   g := *(me.prob.goal)
   if len(me.Result) == len(g) {
      for i := range g {
         if g[i] != me.Result[i] {
            return false
         }
      }
      return true
   } else {
      return false
   }
   panic("Return path error")
}
