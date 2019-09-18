// +build base opengl

package cxcore

// op codes
const (
	// time
	OP_TIME_SLEEP = iota + END_OF_CORE_OPS
	OP_TIME_UNIX
	OP_TIME_UNIX_MILLI
	OP_TIME_UNIX_NANO

	// serialize
	OP_SERIAL_PROGRAM

	// os
	OP_OS_GET_WORKING_DIRECTORY
	OP_OS_OPEN
	OP_OS_CLOSE
	OP_OS_RUN
	OP_OS_EXIT

	// http
	// OP_HTTP_GET

	// object explorer
	OP_OBJ_QUERY

	END_OF_BASE_OPS
)

func init() {
	// time
	AddOpCode(OP_TIME_SLEEP, "time.Sleep",
		[]*CXArgument{newOpPar(TYPE_I32, false)},
		[]*CXArgument{})
	AddOpCode(OP_TIME_UNIX_MILLI, "time.UnixMilli",
		[]*CXArgument{},
		[]*CXArgument{newOpPar(TYPE_I64, false)})
	AddOpCode(OP_TIME_UNIX_NANO, "time.UnixNano",
		[]*CXArgument{},
		[]*CXArgument{newOpPar(TYPE_I64, false)})

	// http
	// AddOpCode(OP_HTTP_GET, "http.Get",
	// 	[]*CXArgument{newOpPar(TYPE_STR, false)},
	// 	[]*CXArgument{newOpPar(TYPE_STR, false)})

	// os
	AddOpCode(OP_OS_GET_WORKING_DIRECTORY, "os.GetWorkingDirectory",
		[]*CXArgument{},
		[]*CXArgument{newOpPar(TYPE_STR, false)})
	AddOpCode(OP_OS_OPEN, "os.Open",
		[]*CXArgument{newOpPar(TYPE_STR, false)},
		[]*CXArgument{})
	AddOpCode(OP_OS_CLOSE, "os.Close",
		[]*CXArgument{newOpPar(TYPE_STR, false)},
		[]*CXArgument{})
	AddOpCode(OP_OS_RUN, "os.Run",
		[]*CXArgument{newOpPar(TYPE_STR, false), newOpPar(TYPE_I32, false), newOpPar(TYPE_I32, false), newOpPar(TYPE_STR, false)},
		[]*CXArgument{newOpPar(TYPE_I32, false), newOpPar(TYPE_I32, false), newOpPar(TYPE_STR, false)})
	AddOpCode(OP_OS_EXIT, "os.Exit",
		[]*CXArgument{newOpPar(TYPE_I32, false)},
		[]*CXArgument{})

	// exec
	handleOpcode := func(opCode int) opcodeHandler {
		switch opCode {
		// time
		case OP_TIME_SLEEP:
			return op_time_Sleep
		case OP_TIME_UNIX:
		case OP_TIME_UNIX_MILLI:
			return op_time_UnixMilli
		case OP_TIME_UNIX_NANO:
			return op_time_UnixNano

		// http
		// case OP_HTTP_GET:
		// 	return op_http_get

		// os
		case OP_OS_GET_WORKING_DIRECTORY:
			return op_os_GetWorkingDirectory
		case OP_OS_OPEN:
			return op_os_Open
		case OP_OS_CLOSE:
			return op_os_Close
		case OP_OS_RUN:
			return op_os_Run
		case OP_OS_EXIT:
			return op_os_Exit
		}

		return nil
	}

	opcodeHandlerFinders = append(opcodeHandlerFinders, handleOpcode)
}
