/** Package: core */
package core

type Input interface {
	Initialize(args ...interface{}) error
	Run()
}
