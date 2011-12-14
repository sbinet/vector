package astar

// A min-heap based on the FVal of each state
type stateHeap []State

func newStateHeap() *stateHeap {
   s := stateHeap(make([]State, 1))
   return &s
}

// Return true if the heap is empty
func (me *stateHeap) Empty() bool {
   return len(*me) <= 1
}

// Add some states to the heap
func (me *stateHeap) Push(states ...State) {
   s := *me
   firstNew := len(s)
   s = append(s, states...)
   *me = s
   for i := range states {
      me.HeapifyUp(firstNew + i)
   }
}

// Fix the position of a given state in the heap
func (me *stateHeap) HeapifyUp(idx int) {
   s := *me

   // Get the index of the parent of the node at 'idx'
   parent := idx / 2

   if parent < 1 {
      return
   } else if s[parent].FVal() > s[idx].FVal() {
      // If the node at 'parent' has a larger FVal(), swap it with the node at
      // 'idx'
      s[parent], s[idx] = s[idx], s[parent]

      // Then continue heapifying up
      me.HeapifyUp(parent)
   }
}

func (me *stateHeap) Pop() State {
   s := *me

   // Save the top state, we need to return it later
   ret := s[1]

   // Move the lowest rank item to the top of the list
   s[1] = s[len(s)-1]
   s = s[:len(s)-1]

   // Then fix the heap using heapify down
   *me = s
   me.HeapifyDown(1)

   // Now we have a happy heap and we've got the element we need, return it!
   return ret
}

func (me *stateHeap) HeapifyDown(idx int) {
   var min int

   // Index of the left and right children of the state at 'idx'
   left := idx * 2
   right := idx*2 + 1

   s := *me

   if len(s) <= left {
      // If the left child is out of bounds then there is nothing to be done
      return
   } else if len(s) <= right {
      // If the right child is out of bounds ( but not the left ) then we know
      // the left child has the minimum FVal(), since the right child does not
      // exist.
      min = left
   } else {
      // Both children exist, we need to calculate which one has the lowest
      // FVal(), since it needs to be above the other in the heap.
      if s[right].FVal() < s[left].FVal() {
         min = right
      } else {
         min = left
      }
   }

   // Now we see if the lowest FVal() is less than that of the node at 'idx'.
   // If so then we swap them and continue to HeapifyDown().
   if s[min].FVal() < s[idx].FVal() {
      s[min], s[idx] = s[idx], s[min]
      me.HeapifyDown(min)
   }
}
