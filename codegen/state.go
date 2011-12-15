package codegen

import "astar"

type Operand []int

type Production struct {
   // The number of operations that preceded this one ( i.e. parent.depth + 1 )
   Depth  int
   Parent *Production
   Result Operand
   Op     *Operation
   prob   *ProblemInstance

   // Depth of the operands that were passed to the operation to get this
   A, B int
}

type ProblemInstance struct {
   goal *Operand
   db   *ExprDB
   ops  []Operation
}

type Operation interface {
   // mnemonic for this operation
   String() string

   // Returns all Productions that could come after 'pred' using this
   // operation. They must use 'pred.Result' as an operand. Note that
   // '*Production' satisfies the 'astar.State' interface.
   Children(pred *Production) []astar.State
}



func (me *Production) GetNext() []astar.State {
   // Start with 16, since we're likely to get far more than that, but if we
   // don't get that many it isn't too much wasted memory.
   rv := make([]astar.State, 16)
   for i := range me.prob.ops {
      rv = append(rv, me.prob.ops[i].Children(me)...)
   }

   return rv
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
