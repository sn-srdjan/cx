package cxcore

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/amherag/skycoin/src/cipher/encoder"
)

func opF64F64(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	out1Offset := GetFinalOffset(fp, out1)

	switch out1.Type {
	case TYPE_STR:
		WriteObject(out1Offset, encoder.Serialize(strconv.FormatFloat(ReadF64(fp, inp1), 'f', -1, 64)))
	case TYPE_BYTE:
		WriteMemory(out1Offset, FromByte(byte(ReadF64(fp, inp1))))
	case TYPE_I32:
		WriteMemory(out1Offset, FromI32(int32(ReadF64(fp, inp1))))
	case TYPE_I64:
		WriteMemory(out1Offset, FromI64(int64(ReadF64(fp, inp1))))
	case TYPE_F32:
		WriteMemory(out1Offset, FromF32(float32(ReadF64(fp, inp1))))
	case TYPE_F64:
		WriteMemory(out1Offset, FromF64(ReadF64(fp, inp1)))
	}
}

// The built-in isnan function returns true if operand is nan value.
func opF64Isnan(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	outB0 := FromBool(math.IsNaN(ReadF64(fp, expr.Inputs[0])))
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), outB0)
}

// The print built-in function formats its arguments and prints them.
//
func opF64Print(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1 := expr.Inputs[0]
	fmt.Println(ReadF64(fp, inp1))
}

// The built-in add function returns the sum of the two operands.
//
func opF64Add(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromF64(ReadF64(fp, inp1) + ReadF64(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in sub function returns the difference between the two operands.
//
func opF64Sub(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	var outB1 []byte
	if len(expr.Inputs) == 2 {
		inp2 := expr.Inputs[1]
		outB1 = FromF64(ReadF64(fp, inp1) - ReadF64(fp, inp2))
	} else {
		outB1 = FromF64(-ReadF64(fp, inp1))
	}
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in mul function returns the product of the two operands.
//
func opF64Mul(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromF64(ReadF64(fp, inp1) * ReadF64(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in div function returns the quotient between the two operands.
//
func opF64Div(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromF64(ReadF64(fp, inp1) / ReadF64(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in abs function returns the absolute value of the operand.
//
func opF64Abs(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromF64(math.Abs(ReadF64(fp, inp1)))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in pow function returns x**n for n>0 otherwise 1
//
func opF64Pow(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromF64(math.Pow(ReadF64(fp, inp1), ReadF64(fp, inp2)))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in gt function returns true if operand 1 is larger than operand 2.
//
func opF64Gt(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromBool(ReadF64(fp, inp1) > ReadF64(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in gteq function returns true if operand 1 is greater than or
// equal to operand 2.
//
func opF64Gteq(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromBool(ReadF64(fp, inp1) >= ReadF64(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in lt function returns true if operand 1 is less than operand 2.
//
func opF64Lt(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromBool(ReadF64(fp, inp1) < ReadF64(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in lteq function returns true if operand 1 is less than or equal
// to operand 2.
//
func opF64Lteq(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromBool(ReadF64(fp, inp1) <= ReadF64(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in eq function returns true if operand 1 is equal to operand 2.
//
func opF64Eq(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromBool(ReadF64(fp, inp1) == ReadF64(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in uneq function returns true if operand 1 is different from operand 2.
//
func opF64Uneq(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromBool(ReadF64(fp, inp1) != ReadF64(fp, inp2))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in rand function returns a pseudo-random number in [0.0,1.0) from the default Source
//
func opF64Rand(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), FromF64(rand.Float64()))
}

// The built-in acos function returns the arc cosine of the operand.
//
func opF64Acos(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromF64(math.Acos(ReadF64(fp, inp1)))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in cos function returns the cosine of the operand.
//
func opF64Cos(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromF64(math.Cos(ReadF64(fp, inp1)))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in asin function returns the arc sine of the operand.
//
func opF64Asin(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromF64(math.Asin(ReadF64(fp, inp1)))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in sin function returns the sine of the operand.
//
func opF64Sin(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromF64(math.Sin(ReadF64(fp, inp1)))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in sqrt function returns the square root of the operand.
//
func opF64Sqrt(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromF64(math.Sqrt(ReadF64(fp, inp1)))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in log function returns the natural logarithm of the operand.
//
func opF64Log(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromF64(math.Log(ReadF64(fp, inp1)))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in log2 function returns the 2-logarithm of the operand.
//
func opF64Log2(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromF64(math.Log2(ReadF64(fp, inp1)))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in log10 function returns the 10-logarithm of the operand.
//
func opF64Log10(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	outB1 := FromF64(math.Log10(ReadF64(fp, inp1)))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in max function returns the largest value of the two operands.
//
func opF64Max(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromF64(math.Max(ReadF64(fp, inp1), ReadF64(fp, inp2)))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

// The built-in min function returns the smallest value of the two operands.
//
func opF64Min(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	outB1 := FromF64(math.Min(ReadF64(fp, inp1), ReadF64(fp, inp2)))
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}
