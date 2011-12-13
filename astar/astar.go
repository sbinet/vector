package astar

// Basically a node of the graph being searched. These nodes should know their
// own distance from the start node ( so they can compute their own FVal() )
type State interface {
   // Get a list of states that we can transition to from here
   GetNext() []State

   // A conservative estimate at how long the search path should be if this
   // state is on the route.
   FVal() int

   // Returns true if this is the goal state.
   IsGoal() bool
}

func Search( initial State ) State {
   list := newStateHeap();

   // Start out with the list of states only containing the initial state
   list.Push( initial )
   
   // While there are still states to explore...
   for !list.Empty() {
      // Take the top element off of the list (it's a heap, so we're guarunteed
      // to get the element with the lowest FVal().
      cur := list.Pop()

      if cur.IsGoal() {
         // If this is our goal state SWEET, keep it.
         return cur
      } else {
         // Otherwise add all of its children to the list and keep searchin
         list.Push( cur.GetNext()... )
      }
   }

   return nil
}
