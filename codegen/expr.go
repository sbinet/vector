package codegen

// Uses bit-shifty tricks to map the labels of two nodes to their parent
type ExprDB struct {
   // Look-Up-Table for expressions
   lut map[int]map[int][]int

   // Actual list of expressions
   list []Expr
}

func NewExprDB() *ExprDB {
   return &ExprDB{
      lut:  make(map[int]map[int][]int),
      list: make([]Expr, 1, 256)}
}

// Expressions made up of binary operations
type Expr interface {
   // How many operations are required to calculate this expression
   NumOps() int

   // Labels uniquely identify expressions based on what they compute. Nodes
   // with equivalent labels compute equivalent sums/products
   Label() int
}

const (
   OpAdd = 0
   OpMul = 1
)

type Literal int

func (me *ExprDB) NewLiteral() int {
   rv := len(me.list)
   me.list = append(me.list, Literal(rv))
   return rv
}

func (me Literal) NumOps() int {
   return 0
}

func (me Literal) Label() int {
   return int(me)
}

// Represents a multiplication
type Mul struct {
   db                *ExprDB
   left, right, self int
}

func (me *ExprDB) NewMul(left, right int) int {

   if left > right {
      left, right = right, left
   }

   mp := me.lut[left]
   if mp == nil {
      mp = make(map[int][]int)
      me.lut[left] = mp
   }

   pair := mp[right]
   if pair == nil {
      pair = make([]int, 2)
      mp[right] = pair
   }

   self := pair[OpMul]
   if self == 0 {
      self = len(me.list)
      mul := &Mul{db: me, left: left, right: right, self: self}
      me.list = append(me.list, mul)

      pair[OpMul] = self
   }
   return self
}

func (me *Mul) NumOps() int {
   return me.db.list[me.left].NumOps() + me.db.list[me.right].NumOps() + 1
}

func (me *Mul) Label() int {
   return me.self
}

// Represents an addition
type Add struct {
   db                *ExprDB
   left, right, self int
}

func (me *ExprDB) NewAdd(left, right int) int {

   if left > right {
      left, right = right, left
   }

   mp := me.lut[left]
   if mp == nil {
      mp = make(map[int][]int)
      me.lut[left] = mp
   }

   pair := mp[right]
   if pair == nil {
      pair = make([]int, 2)
      mp[right] = pair
   }

   self := pair[OpAdd]
   if self == 0 {
      self = len(me.list)
      add := &Add{db: me, left: left, right: right, self: self}
      me.list = append(me.list, add)

      pair[OpAdd] = self
   }
   return self
}

func (me *Add) NumOps() int {
   return me.db.list[me.left].NumOps() + me.db.list[me.right].NumOps() + 1
}

func (me *Add) Label() int {
   return me.self
}
