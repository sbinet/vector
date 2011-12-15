package codegen

import "testing"

func TestExprDB(t *testing.T) {
   db := NewExprDB()

   a, b, c := db.NewLiteral(), db.NewLiteral(), db.NewLiteral()

   aPbXc := db.NewAdd(a, db.NewMul(b, c))

   bXc := db.NewMul(b, c)

   if aPbXc != db.NewAdd(a, bXc) {
      t.Errorf("Did not generate same name for 'a + b * c' twice in a row")
   }

   if aPbXc == db.NewAdd(a, db.NewMul(a, b)) {
      t.Errorf("Generated same name for different expressions")
   }
}
