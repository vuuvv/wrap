package wrap

import (
	"errors"
	"testing"
)

func testResultError() (res Result[int]) {
	return Error[int](errors.New("test"))
}

func testResultOK() (res Result[int]) {
	return Ok[int](1)
}

func testRecover() (res Result[int]) {
	defer res.Recover()
	res = testResultOK()
	testResultError().Unwrap()
	return testResultOK()
}

func testRecoverOr() (res Result[int]) {
	defer RecoverOr(&res, Ok(100))
	testResultError().Unwrap()
	return testResultOK()
}

func testRecoverHandle() (res Result[int]) {
	defer RecoverHandle(&res, func(res *Result[int], err error) Result[int] {
		return Ok(10000)
	})
	testResultError().Unwrap()
	return testResultOK()
}

func TestUnwrapError(t *testing.T) {
	t.Log(testResultOK().Unwrap())
	testResultError().Unwrap()
}

func TestRecover(t *testing.T) {
	r := testRecover()
	t.Log("error:", r.err)
}

func TestRecoverOr(t *testing.T) {
	r := testRecoverOr()
	t.Log("val:", r.value)
	t.Log("error:", r.err)
}

func TestRecoverHandle(t *testing.T) {
	r := testRecoverHandle()
	t.Log("val:", r.value)
	t.Log("error:", r.err)
}
