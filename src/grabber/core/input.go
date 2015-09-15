/** Package: core */
package core

type Input interface {
	Initialize(args ...interface{}) error
	Run()
}

func Test() string {
	return "This is a test!"
}
