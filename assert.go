package optional

type assert func(args Args) bool

var (
	//所有参数均不为nil则返回true
	nilAssert assert = func(args Args) bool {
		if args == nil || len(args) <= 0 {
			return false
		}
		for i := 0; i < len(args); i++ {
			if args[i] == nil {
				return false
			}
		}
		return true
	}
	//参数中的error类型均为nil则返回true
	errorAssert assert = func(args Args) bool {
		if args == nil || len(args) <= 0 {
			return true
		}
		for i := 0; i < len(args); i++ {
			if arg, ok := args[i].(error); ok && arg != nil {
				return false
			}
		}
		return true
	}
)

func NilAssert(index ...int) assert {
	if len(index) == 0 {
		return nilAssert
	}
	//指定下标处的参数均不为nil则返回true
	return func(args Args) bool {
		for i := 0; i < len(index); i++ {
			if index[i] < 0 || index[i] >= len(args) || args[index[i]] == nil {
				return false
			}
		}
		return true
	}
}

func ErrorAssert() assert {
	return errorAssert
}
