package cxcore

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/amherag/skycoin/src/cipher/encoder"
)

func opLt(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_BYTE:
		outB1 = FromBool(ReadByte(fp, inp1) < ReadByte(fp, inp2))
	case TYPE_STR:
		outB1 = FromBool(ReadStr(fp, inp1) < ReadStr(fp, inp2))
	case TYPE_I32:
		outB1 = FromBool(ReadI32(fp, inp1) < ReadI32(fp, inp2))
	case TYPE_I64:
		outB1 = FromBool(ReadI64(fp, inp1) < ReadI64(fp, inp2))
	case TYPE_F32:
		outB1 = FromBool(ReadF32(fp, inp1) < ReadF32(fp, inp2))
	case TYPE_F64:
		outB1 = FromBool(ReadF64(fp, inp1) < ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opGt(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_BYTE:
		outB1 = FromBool(ReadByte(fp, inp1) > ReadByte(fp, inp2))
	case TYPE_STR:
		outB1 = FromBool(ReadStr(fp, inp1) > ReadStr(fp, inp2))
	case TYPE_I32:
		outB1 = FromBool(ReadI32(fp, inp1) > ReadI32(fp, inp2))
	case TYPE_I64:
		outB1 = FromBool(ReadI64(fp, inp1) > ReadI64(fp, inp2))
	case TYPE_F32:
		outB1 = FromBool(ReadF32(fp, inp1) > ReadF32(fp, inp2))
	case TYPE_F64:
		outB1 = FromBool(ReadF64(fp, inp1) > ReadF64(fp, inp2))
	}
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opLteq(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_BYTE:
		outB1 = FromBool(ReadByte(fp, inp1) <= ReadByte(fp, inp2))
	case TYPE_STR:
		outB1 = FromBool(ReadStr(fp, inp1) <= ReadStr(fp, inp2))
	case TYPE_I32:
		outB1 = FromBool(ReadI32(fp, inp1) <= ReadI32(fp, inp2))
	case TYPE_I64:
		outB1 = FromBool(ReadI64(fp, inp1) <= ReadI64(fp, inp2))
	case TYPE_F32:
		outB1 = FromBool(ReadF32(fp, inp1) <= ReadF32(fp, inp2))
	case TYPE_F64:
		outB1 = FromBool(ReadF64(fp, inp1) <= ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opGteq(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_BYTE:
		outB1 = FromBool(ReadByte(fp, inp1) >= ReadByte(fp, inp2))
	case TYPE_STR:
		outB1 = FromBool(ReadStr(fp, inp1) >= ReadStr(fp, inp2))
	case TYPE_I32:
		outB1 = FromBool(ReadI32(fp, inp1) >= ReadI32(fp, inp2))
	case TYPE_I64:
		outB1 = FromBool(ReadI64(fp, inp1) >= ReadI64(fp, inp2))
	case TYPE_F32:
		outB1 = FromBool(ReadF32(fp, inp1) >= ReadF32(fp, inp2))
	case TYPE_F64:
		outB1 = FromBool(ReadF64(fp, inp1) >= ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opEqual(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_BYTE:
		outB1 = FromBool(ReadByte(fp, inp1) == ReadByte(fp, inp2))
	case TYPE_BOOL:
		outB1 = FromBool(ReadBool(fp, inp1) == ReadBool(fp, inp2))
	case TYPE_STR:
		outB1 = FromBool(ReadStr(fp, inp1) == ReadStr(fp, inp2))
	case TYPE_I32:
		outB1 = FromBool(ReadI32(fp, inp1) == ReadI32(fp, inp2))
	case TYPE_I64:
		outB1 = FromBool(ReadI64(fp, inp1) == ReadI64(fp, inp2))
	case TYPE_F32:
		outB1 = FromBool(ReadF32(fp, inp1) == ReadF32(fp, inp2))
	case TYPE_F64:
		outB1 = FromBool(ReadF64(fp, inp1) == ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opUnequal(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_BYTE:
		outB1 = FromBool(ReadByte(fp, inp1) != ReadByte(fp, inp2))
	case TYPE_BOOL:
		outB1 = FromBool(ReadBool(fp, inp1) != ReadBool(fp, inp2))
	case TYPE_STR:
		outB1 = FromBool(ReadStr(fp, inp1) != ReadStr(fp, inp2))
	case TYPE_I32:
		outB1 = FromBool(ReadI32(fp, inp1) != ReadI32(fp, inp2))
	case TYPE_I64:
		outB1 = FromBool(ReadI64(fp, inp1) != ReadI64(fp, inp2))
	case TYPE_F32:
		outB1 = FromBool(ReadF32(fp, inp1) != ReadF32(fp, inp2))
	case TYPE_F64:
		outB1 = FromBool(ReadF64(fp, inp1) != ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opBitand(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_I32:
		outB1 = FromI32(ReadI32(fp, inp1) & ReadI32(fp, inp2))
	case TYPE_I64:
		outB1 = FromI64(ReadI64(fp, inp1) & ReadI64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opBitor(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_I32:
		outB1 = FromI32(ReadI32(fp, inp1) | ReadI32(fp, inp2))
	case TYPE_I64:
		outB1 = FromI64(ReadI64(fp, inp1) | ReadI64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opBitxor(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_I32:
		outB1 = FromI32(ReadI32(fp, inp1) ^ ReadI32(fp, inp2))
	case TYPE_I64:
		outB1 = FromI64(ReadI64(fp, inp1) ^ ReadI64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opMul(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_BYTE:
		outB1 = FromByte(ReadByte(fp, inp1) * ReadByte(fp, inp2))
	case TYPE_I32:
		outB1 = FromI32(ReadI32(fp, inp1) * ReadI32(fp, inp2))
	case TYPE_I64:
		outB1 = FromI64(ReadI64(fp, inp1) * ReadI64(fp, inp2))
	case TYPE_F32:
		outB1 = FromF32(ReadF32(fp, inp1) * ReadF32(fp, inp2))
	case TYPE_F64:
		outB1 = FromF64(ReadF64(fp, inp1) * ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opDiv(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_BYTE:
		outB1 = FromByte(ReadByte(fp, inp1) / ReadByte(fp, inp2))
	case TYPE_I32:
		outB1 = FromI32(ReadI32(fp, inp1) / ReadI32(fp, inp2))
	case TYPE_I64:
		outB1 = FromI64(ReadI64(fp, inp1) / ReadI64(fp, inp2))
	case TYPE_F32:
		outB1 = FromF32(ReadF32(fp, inp1) / ReadF32(fp, inp2))
	case TYPE_F64:
		outB1 = FromF64(ReadF64(fp, inp1) / ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opMod(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_BYTE:
		outB1 = FromByte(ReadByte(fp, inp1) % ReadByte(fp, inp2))
	case TYPE_I32:
		outB1 = FromI32(ReadI32(fp, inp1) % ReadI32(fp, inp2))
	case TYPE_I64:
		outB1 = FromI64(ReadI64(fp, inp1) % ReadI64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opAdd(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_BYTE:
		outB1 = FromByte(ReadByte(fp, inp1) + ReadByte(fp, inp2))
	case TYPE_I32:
		outB1 = FromI32(ReadI32(fp, inp1) + ReadI32(fp, inp2))
	case TYPE_I64:
		outB1 = FromI64(ReadI64(fp, inp1) + ReadI64(fp, inp2))
	case TYPE_F32:
		outB1 = FromF32(ReadF32(fp, inp1) + ReadF32(fp, inp2))
	case TYPE_F64:
		outB1 = FromF64(ReadF64(fp, inp1) + ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opSub(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	var outB1 []byte
	if len(expr.Inputs) == 2 {
		inp2 := expr.Inputs[1]
		switch inp1.Type {
		case TYPE_BYTE:
			outB1 = FromByte(ReadByte(fp, inp1) - ReadByte(fp, inp2))
		case TYPE_I32:
			outB1 = FromI32(ReadI32(fp, inp1) - ReadI32(fp, inp2))
		case TYPE_I64:
			outB1 = FromI64(ReadI64(fp, inp1) - ReadI64(fp, inp2))
		case TYPE_F32:
			outB1 = FromF32(ReadF32(fp, inp1) - ReadF32(fp, inp2))
		case TYPE_F64:
			outB1 = FromF64(ReadF64(fp, inp1) - ReadF64(fp, inp2))
		}
	} else {
		switch inp1.Type {
		case TYPE_BYTE:
			outB1 = FromByte(-ReadByte(fp, inp1))
		case TYPE_I32:
			outB1 = FromI32(-ReadI32(fp, inp1))
		case TYPE_I64:
			outB1 = FromI64(-ReadI64(fp, inp1))
		case TYPE_F32:
			outB1 = FromF32(-ReadF32(fp, inp1))
		case TYPE_F64:
			outB1 = FromF64(-ReadF64(fp, inp1))
		}
	}
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opBitshl(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_I32:
		outB1 = FromI32(int32(uint32(ReadI32(fp, inp1)) << uint32(ReadI32(fp, inp2))))
	case TYPE_I64:
		outB1 = FromI64(int64(uint64(ReadI64(fp, inp1)) << uint64(ReadI64(fp, inp2))))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opBitshr(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_I32:
		outB1 = FromI32(int32(uint32(ReadI32(fp, inp1)) >> uint32(ReadI32(fp, inp2))))
	case TYPE_I64:
		outB1 = FromI64(int64(uint32(ReadI64(fp, inp1)) >> uint32(ReadI64(fp, inp2))))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opBitclear(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TYPE_I32:
		outB1 = FromI32(int32(uint32(ReadI32(fp, inp1)) &^ uint32(ReadI32(fp, inp2))))
	case TYPE_I64:
		outB1 = FromI64(int64(uint32(ReadI64(fp, inp1)) &^ uint32(ReadI64(fp, inp2))))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opLen(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	elt := GetAssignmentElement(inp1)

	if elt.IsSlice || elt.Type == TYPE_AFF {
		var sliceOffset = GetSliceOffset(fp, inp1)
		var sliceLen []byte
		if sliceOffset > 0 {
			sliceLen = GetSliceHeader(sliceOffset)[4:8]
		} else if sliceOffset == 0 {
			sliceLen = FromI32(0)
		} else {
			panic(CX_RUNTIME_ERROR)
		}

		WriteMemory(GetFinalOffset(fp, out1), sliceLen)
		// TODO: Had to add elt.Lengths to avoid doing this for arrays, but not entirely sure why
	} else if elt.Type == TYPE_STR && elt.Lengths == nil {
		var strOffset = GetStrOffset(fp, inp1)
		// Checking if the string lives on the heap.
		if strOffset > PROGRAM.HeapStartsAt {
			// Then it's on the heap and we need to consider
			// the object's header.
			strOffset += OBJECT_HEADER_SIZE
		}

		WriteMemory(GetFinalOffset(fp, out1), PROGRAM.Memory[strOffset:strOffset+STR_HEADER_SIZE])
	} else {
		outB1 := FromI32(int32(elt.Lengths[0]))
		WriteMemory(GetFinalOffset(fp, out1), outB1)
	}
}

func opAppend(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]

	eltInp1 := GetAssignmentElement(inp1)
	eltOut1 := GetAssignmentElement(out1)
	if inp1.Type != inp2.Type || inp1.Type != out1.Type || !eltInp1.IsSlice || !eltOut1.IsSlice {
		panic(CX_RUNTIME_INVALID_ARGUMENT)
	}

	// outputSlicePointer := GetFinalOffset(fp, out1)
	// outputSliceOffset := GetPointerOffset(int32(outputSlicePointer))
	inputSliceOffset := GetSliceOffset(fp, inp1)

	var inputSliceLen int32
	if inputSliceOffset != 0 {
		inputSliceLen = GetSliceLen(inputSliceOffset)
	}

	// Preparing slice in case more memory is needed for the new element.
	// outputSliceOffset = SliceAppendResize(outputSliceOffset, inputSliceOffset, inp2.Size)
	outputSliceOffset := SliceAppendResize(fp, out1, inp1, inp2.Size)

	// We need to update the address of the output and input, as the final offsets
	// could be on the heap and they could have been moved by the GC.
	outputSlicePointer := GetFinalOffset(fp, out1)
	inputSliceOffset = GetSliceOffset(fp, inp1)

	var obj []byte
	if inp2.Type == TYPE_STR || inp2.Type == TYPE_AFF {
		obj = encoder.SerializeAtomic(int32(GetStrOffset(fp, inp2)))
	} else {
		obj = ReadMemory(GetFinalOffset(fp, inp2), inp2)
	}

	SliceAppendWrite(outputSliceOffset, inputSliceOffset, obj, inputSliceLen)

	copy(PROGRAM.Memory[outputSlicePointer:], encoder.SerializeAtomic(outputSliceOffset))
}

func opResize(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]

	if inp1.Type != out1.Type || !GetAssignmentElement(inp1).IsSlice || !GetAssignmentElement(out1).IsSlice {
		panic(CX_RUNTIME_INVALID_ARGUMENT)
	}

	outputSlicePointer := GetFinalOffset(fp, out1)
	// outputSliceOffset := GetPointerOffset(int32(outputSlicePointer))

	outputSliceOffset := int32(SliceResize(fp, out1, inp1, ReadI32(fp, inp2), GetAssignmentElement(inp1).TotalSize))
	copy(PROGRAM.Memory[outputSlicePointer:], encoder.SerializeAtomic(outputSliceOffset))
}

func opInsert(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, inp3, out1 := expr.Inputs[0], expr.Inputs[1], expr.Inputs[2], expr.Outputs[0]

	if inp1.Type != inp3.Type || inp1.Type != out1.Type || !GetAssignmentElement(inp1).IsSlice || !GetAssignmentElement(out1).IsSlice {
		panic(CX_RUNTIME_INVALID_ARGUMENT)
	}

	outputSlicePointer := GetFinalOffset(fp, out1)
	// outputSliceOffset := GetPointerOffset(int32(outputSlicePointer))

	var obj []byte
	if inp3.Type == TYPE_STR || inp3.Type == TYPE_AFF {
		obj = encoder.SerializeAtomic(int32(GetStrOffset(fp, inp3)))
	} else {
		obj = ReadMemory(GetFinalOffset(fp, inp3), inp3)
	}

	outputSliceOffset := int32(SliceInsert(fp, out1, inp1, ReadI32(fp, inp2), obj))
	copy(PROGRAM.Memory[outputSlicePointer:], encoder.SerializeAtomic(outputSliceOffset))
}

func opRemove(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]

	if inp1.Type != out1.Type || !GetAssignmentElement(inp1).IsSlice || !GetAssignmentElement(out1).IsSlice {
		panic(CX_RUNTIME_INVALID_ARGUMENT)
	}

	outputSlicePointer := GetFinalOffset(fp, out1)
	// outputSliceOffset := GetPointerOffset(int32(outputSlicePointer))

	outputSliceOffset := int32(SliceRemove(fp, out1, inp1, ReadI32(fp, inp2), int32(GetAssignmentElement(inp1).TotalSize)))
	copy(PROGRAM.Memory[outputSlicePointer:], encoder.SerializeAtomic(outputSliceOffset))
}

func opCopy(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	dstInput := expr.Inputs[0]
	srcInput := expr.Inputs[1]
	dstOffset := GetSliceOffset(fp, dstInput)
	srcOffset := GetSliceOffset(fp, srcInput)

	dstElem := GetAssignmentElement(dstInput)
	srcElem := GetAssignmentElement(srcInput)

	if dstInput.Type != srcInput.Type || !dstElem.IsSlice || !srcElem.IsSlice || dstElem.TotalSize != srcElem.TotalSize {
		panic(CX_RUNTIME_INVALID_ARGUMENT)
	}

	var count int
	if dstInput.Type == srcInput.Type && dstOffset >= 0 && srcOffset >= 0 {
		count = copy(GetSliceData(dstOffset, dstElem.TotalSize), GetSliceData(srcOffset, srcElem.TotalSize))
		if count%dstElem.TotalSize != 0 {
			panic(CX_RUNTIME_ERROR)
		}
	} else {
		panic(CX_RUNTIME_INVALID_ARGUMENT)
	}
	WriteMemory(GetFinalOffset(fp, expr.Outputs[0]), FromI32(int32(count/dstElem.TotalSize)))
}

func buildString(expr *CXExpression, fp int) []byte {
	inp1 := expr.Inputs[0]

	fmtStr := ReadStr(fp, inp1)

	var res []byte
	var specifiersCounter int
	var lenStr = int(len(fmtStr))

	for c := 0; c < len(fmtStr); c++ {
		var nextCh byte
		ch := fmtStr[c]
		if c < lenStr-1 {
			nextCh = fmtStr[c+1]
		}
		if ch == '\\' {
			switch nextCh {
			case '%':
				c++
				res = append(res, nextCh)
				continue
			case 'n':
				c++
				res = append(res, '\n')
				continue
			default:
				res = append(res, ch)
				continue
			}
		}
		if ch == '%' {
			if specifiersCounter+1 == len(expr.Inputs) {
				res = append(res, []byte(fmt.Sprintf("%%!%c(MISSING)", nextCh))...)
				c++
				continue
			}

			inp := expr.Inputs[specifiersCounter+1]
			switch nextCh {
			case 's':
				res = append(res, []byte(checkForEscapedChars(ReadStr(fp, inp)))...)
			case 'd':
				switch inp.Type {
				case TYPE_I32:
					res = append(res, []byte(strconv.FormatInt(int64(ReadI32(fp, inp)), 10))...)
				case TYPE_I64:
					res = append(res, []byte(strconv.FormatInt(ReadI64(fp, inp), 10))...)
				}
			case 'f':
				switch inp.Type {
				case TYPE_F32:
					res = append(res, []byte(strconv.FormatFloat(float64(ReadF32(fp, inp)), 'f', 7, 32))...)
				case TYPE_F64:
					res = append(res, []byte(strconv.FormatFloat(ReadF64(fp, inp), 'f', 16, 64))...)
				}
			case 'v':
				res = append(res, []byte(GetPrintableValue(fp, inp))...)
			}
			c++
			specifiersCounter++
		} else {
			res = append(res, ch)
		}
	}

	if specifiersCounter != len(expr.Inputs)-1 {
		extra := "%!(EXTRA "
		// for _, inp := range expr.Inputs[:specifiersCounter] {
		lInps := len(expr.Inputs[specifiersCounter+1:])
		for c := 0; c < lInps; c++ {
			inp := expr.Inputs[specifiersCounter+1+c]
			elt := GetAssignmentElement(inp)
			typ := ""
			_ = typ
			if elt.CustomType != nil {
				// then it's custom type
				typ = elt.CustomType.Name
			} else {
				// then it's native type
				typ = TypeNames[elt.Type]
			}

			if c == lInps-1 {
				extra += fmt.Sprintf("%s=%s", typ, GetPrintableValue(fp, elt))
			} else {
				extra += fmt.Sprintf("%s=%s, ", typ, GetPrintableValue(fp, elt))
			}

		}

		extra += ")"

		res = append(res, []byte(extra)...)
	}

	return res
}

func opSprintf(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	out1 := expr.Outputs[0]
	out1Offset := GetFinalOffset(fp, out1)

	byts := encoder.Serialize(string(buildString(expr, fp)))
	WriteObject(out1Offset, byts)
}

func opPrintf(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	fmt.Print(string(buildString(expr, fp)))
}

func opRead(prgrm *CXProgram) {
	expr := prgrm.GetExpr()
	fp := prgrm.GetFramePointer()

	out1 := expr.Outputs[0]
	out1Offset := GetFinalOffset(fp, out1)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	// text = strings.Trim(text, " \n")
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	if err != nil {
		panic("")
	}
	byts := encoder.Serialize(text)
	size := encoder.Serialize(int32(len(byts)) + OBJECT_HEADER_SIZE)
	heapOffset := AllocateSeq(len(byts) + OBJECT_HEADER_SIZE)

	var header = make([]byte, OBJECT_HEADER_SIZE)
	for c := 5; c < OBJECT_HEADER_SIZE; c++ {
		header[c] = size[c-5]
	}

	obj := append(header, byts...)

	WriteMemory(heapOffset, obj)

	off := encoder.SerializeAtomic(int32(heapOffset + OBJECT_HEADER_SIZE))

	WriteMemory(out1Offset, off)
}
