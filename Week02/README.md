

## Go Errors

### Errors 简介

```go
type error interface{
	Error() string
}

```

在go语言的定义中， error 就是个简单的接口。只要实现了`Error() string` 的方法的值就能被当做error 使用。


### GO 1.13 之前的Errors

`Go errors are values` 

在go中， 错误也是值类型。所以它能够和其他类型进行比较。所以我们经常的做法都是用`error` 和 `nil`进行比较来看看之前的操作是不是正常的。
```go
if err != nil{
  // do somthing here
}
// continue the logic
```

但是有时候 我们需要和其他已经预定义的 error 进行比较，来获取是不是某一个具体的error 发生了，从而进行一些处理， 比如错误的降级处理。

```go
var ErrNotFound = errors.New("not found")
if err == ErrNotFound {
    // something wasn't found
}

```
但是这样 `ErrorNotFound` 就会成为这个package的一部分。这样必然会导致这个package的公有导出部分变大。

```go
package bufio

var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)
```
比如我们 `go 1.14.2` 中的bufio包中，导出的这4个Err... 会被外部引用者强依赖来判断是不是某个错误， 这样意味着以后版本升级都很难对其进行改动。

有没有更好的办法呢？ 当然是有的。

```go

// https://github.com/juju/errors/blob/3fe23663418f/errortypes.go#L279

// badRequest represents an error when a request has bad parameters.
type badRequest struct {
	 Err
}

// IsBadRequest reports whether err was created with BadRequestf() or
// NewBadRequest().
func IsBadRequest(err error) bool {
	err = Cause(err)
	_, ok := err.(*badRequest)
	return ok
}
```
在这个包中， 他定义了 `badRequest`这种类型的error。然后通过定义`IsBadRequest`函数来进行断言，判断是不是一个`badRequest`。

但是只是package导出变面积的增大并不是最大的问题。最大的问题是： 我们只能够获取，error 的文本信息， 而不能够获取导致错误的上下文。导致调试的难度增加。

虽然我们能通过定义自定义的错误类型来为我们获取更多的信息。比如:

```go
// PathError records an error and the operation and file path that caused it.
type PathError struct {
	Op   string
	Path string
	Err  error
}

```
但是这样也只是提供了部分信息。但是与调用者产生了强耦合。不利于package的改造。


### GO 1.13 Errors

其实 go.1.13 实现的功能和 `github.com/pkg/errors` 这个包中的功能差不多。
就用 `github.com/pkg/errors`这个包来举例子。

#### Wrap() 为错误添加上下文

```go
r, _ := os.Open("/dev/a.txt")
_, err := ioutil.ReadAll(r)

if err != nil {
  err = errors.Wrap(err, "read failed")

  log.Printf("%+v", err)
}
```

能够返回：
```
2020/12/01 16:46:33 invalid argument
read failed
main.main
        /Users/xxx/hey/test.go:16
runtime.main
        /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/proc.go:203
runtime.goexit
        /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/asm_amd64.s:1373

```

#### Unwrap 返回错误的根因

```go
var ErrSomething = errors.New("something")

func main() {
	err := ErrSomething
	if err != nil {
		err = errors.Wrap(err, "more information")
		cause := errors.Cause(err)
		fmt.Println("cause is ErrSomething", errors.Is(cause, ErrSomething))
	}
}
```

我们也能通过 `errors.Is(err, target error)` 来比较 `sentinel error`

#### WithMessage  为错误添加更多的信息

```go
package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	cause := errors.New("whoops")
	err := errors.WithMessage(cause, "oh noes")
	fmt.Println(err)
	// oh noes: whoops

}

```

### 何时使用 Wrap

> 在发生错误的地方使用Wrap.

虽然Wrap 能够保存堆栈信息对调试带来了便利。
但是也不能够滥用Wrap。比如Wrap 两次， 这样会带来过多的堆栈信息反而不利于调试。

我们需要在产生错误的地方调用Wrap 保存错误堆栈， 然后一层层往上传。
也可以在应用的最上层 通过 调用  `errors.Is(err, SentinelError)` 来进行一些对比处理。在最上层或者中间件处统一记录日志。

