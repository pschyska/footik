package tests

import (
	"github.com/robfig/revel"
	"net/url"
	"regexp"
)

type ApplicationTest struct {
	rev.TestSuite
}

func (t ApplicationTest) Before() {
	println("Set up")
}

func (t ApplicationTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html")
}

func (t ApplicationTest) TestThatHelloActionWorks() {
	v := url.Values{}
	v.Set("myName", "Ava")
	t.PostForm("/Application/Hello", v)
	t.AssertOk()
	t.AssertContentType("text/html")
	r := regexp.MustCompile("<h1>Hello Ava</h1>")
	res := r.Find(t.ResponseBody)
	t.Assert(len(res) > 0)
}

func (t ApplicationTest) After() {
	println("Tear down")
}
