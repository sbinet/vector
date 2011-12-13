package astar

import "testing"

type testState struct {
   fVal int
}

func (me *testState) GetNext() []State {
   return []State{}
}

func (me *testState) FVal() int {
   return me.fVal
}

func (me *testState) IsGoal() bool {
   return false
}

func TestHeapMethods(t *testing.T) {
   states := []State{
      &testState{1},
      &testState{2},
      &testState{5},
      &testState{4},
      &testState{3},
      &testState{15},
      &testState{11},
      &testState{3}}

   heap := newStateHeap()

   // Let's try adding the states one at a time.
   for i := range states {
      heap.Push(states[i])
   }

   // Make sure that the elements are in order
   last := 0
   count := 0
   for !heap.Empty() {
      cur := heap.Pop().FVal()
      if cur < last {
         t.Errorf("Heap ordering incorrect.")
         return
      }
      last = cur
      count++
   }

   if count != len(states) {
      t.Errorf("Fewer elements were removed from the heap than were inserted.")
   }

   // Now we try using the '...' operator to add a while slice of elements to
   // the list at once.
   heap.Push(states...)
   heap.Push(states...)

   // Again, make sure that the elements are in order
   last = 0
   count = 0
   for !heap.Empty() {
      cur := heap.Pop().FVal()
      if cur < last {
         t.Errorf("Heap ordering incorrect.")
         return
      }
      last = cur
      count++
   }

   if count != len(states)*2 {
      t.Errorf("Fewer elements were removed from the heap than were inserted.")
   }

}
