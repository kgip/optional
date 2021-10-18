package test

import (
	"errors"
	"fmt"
	"optional"
	"testing"
)

func F(params ...interface{}) (int, bool, error) {
	fmt.Println(params...)
	return len(params), true, nil
}

func TestParams(t *testing.T) {
	//optional.Of("1",errors.New("2")).Map(func(params []interface{}) interface{} {
	//	return nil
	//})
	//optional.Of(F())
	//f, err := F()
	//list := []interface{}{f,err}
	//t.Log(list)
	i := 1
	t.Log(fmt.Sprintf("%v", &i))
	t.Log(optional.Of(F(1, 2, 3, "aa")).AssertChan(optional.ErrorAssert(), optional.NilAssert(1, 0), func(args optional.Args) bool {
		return args[1].(bool)
	}).Map(func(params optional.Args) optional.Args {
		i := params[0].(int)
		i++
		t.Log(i)
		return optional.Return([]int{i, i + 1})
	}).AssertChan(optional.ErrorAssert()).Map(func(params optional.Args) optional.Args {
		t.Log(params)
		return optional.Return("hello optional", errors.New("aaa"))
	}).OrElseIndex(0, "failed optional"))
}
