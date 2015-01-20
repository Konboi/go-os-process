# os process

## type Process

* pid と プロセスに完成するフィールドが含まれてることが多い
* StartProcess メソッドで作ったプロセスを格納する

```go-lang
type Process struct {
    Pid int
    // contains filtered or unexported fields
}
```

## func FindProcess(pid int) (p *Process, err error)

* process id を渡すとプロセスを返す

![img](http://i.gyazo.com/350d201e5edb5a93a33274f895bdcff7.png)

```go-lang
package main

import (
	"fmt"
	"os"
)

func main() {
	ps, err := os.FindProcess(99360)
	if err != nil {
		fmt.Printf("error is %s :", err.Error())
	}
	fmt.Print(ps.Pid)}

```


## func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)

* program経由でプロセスを作る
* `os/exec` 使ったほうが簡単

## func (p *Process) Kill() error

* process殺す
* `type Process`を渡す

![kill](http://i.gyazo.com/2753e92349e687fa0bbc892bb5c3a46c.gif)


```go-lang
package main

import (
	"fmt"
	"os"
)

func main() {
	ps, err := os.FindProcess(54327)

	if err != nil {
		fmt.Printf("error is %s :", err.Error())
	}
	fmt.Print(ps.Pid)

	err = ps.Kill()
	if err != nil {
		fmt.Printf("kill process error is %s :", err.Error())
	}
}
```

## func (p *Process) Release() error

* processをリリースする
* リリースって？


##func (p *Process) Signal(sig Signal) error

* processに信号を送ります
* windowsで中断のシグナルは送れない

## func (p *Process) Wait() (*ProcessState, error)

* processが終了するのを待つメソッド
* ProcessState経由でステータスがわかる？


## type ProcessState

## func (p *ProcessState) Exited() bool

* processが消えたかどうか

## func (p *ProcessState) Pid() int

* processのpid

## func (p *ProcessState) String() string

* process名

## func (p *ProcessState) Success() bool

* processがきちんと死んだかどうか

## func (p *ProcessState) Sys() interface{}

* 終了するprocessの依存している情報を返す

例えば unix だと syscallのwaitstatusとか

## func (p *ProcessState) SysUsage() interface{}

* 終了するprocessの依存しているリソースを返す

Unix の syscall.Rusage とか

## func (p *ProcessState) SystemTime() time.Duration

* 終了したプロセスや子プロセスのシステムのcpu時間を返す


## func (p *ProcessState) UserTime() time.Duration

* 終了したプロセスや子プロセスのユーザーのcpu時間を返す
