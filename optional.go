package optional

type Args []interface{}

func Return(args ...interface{}) Args {
	return args
}

type Optional struct {
	args    Args
	asserts []assert
}

func Of(args ...interface{}) *Optional {
	return &Optional{args: args, asserts: []assert{NilAssert()}}
}

// Map 断言成功后将params经mapper规则变换后生成新的Optional返回,mapper处理的结果以[]interface{}的形式返回
func (optional *Optional) Map(mapper func(args Args) Args) *Optional {
	if optional.Assert(optional.args) {
		optional.args = mapper(optional.args)
	}
	return optional
}

// ErrorMap 断言失败后将params经mapper规则变换后生成新的Optional返回,mapper处理的结果以[]interface{}的形式返回
func (optional *Optional) ErrorMap(mapper func(args Args) Args) *Optional {
	if !optional.Assert(optional.args) {
		optional.args = mapper(optional.args)
	}
	return optional
}

func (optional *Optional) Assert(args Args) bool {
	if optional.asserts == nil || len(optional.asserts) <= 0 {
		return true
	}
	for i := 0; i < len(optional.asserts); i++ {
		if !optional.asserts[i](args) {
			return false
		}
	}
	return true
}

func (optional *Optional) AddAssert(assert assert) *Optional {
	optional.asserts = append(optional.asserts, assert)
	return optional
}

func (optional *Optional) AssertChan(asserts ...assert) *Optional {
	optional.asserts = asserts
	return optional
}

func (optional *Optional) DefaultAssertChan() *Optional {
	optional.asserts = []assert{NilAssert()}
	return optional
}

func (optional *Optional) OrElse(other ...interface{}) []interface{} {
	if optional.Assert(optional.args) {
		return optional.args
	}
	return other
}

func (optional *Optional) OrElseGet(other func(args Args) Args) Args {
	return optional.OrElse(other(optional.args))
}

func (optional *Optional) OrElseIndex(i int, other interface{}) interface{} {
	if i >= 0 && i < len(optional.args) && optional.Assert(optional.args) {
		return optional.args[i]
	}
	return other
}

func (optional *Optional) OrElseGetIndex(i int, other func(args Args) interface{}) interface{} {
	return optional.OrElseIndex(i, other(optional.args))
}

func (optional *Optional) Get() Args {
	return optional.args
}
