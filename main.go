package main

import (
	"fmt"
	"time"

	"github.com/iovisor/gobpf/elf"
	"golang.org/x/sys/unix"
)

func main() {

		err := unix.Setrlimit(unix.RLIMIT_MEMLOCK, &unix.Rlimit{
		Cur: unix.RLIM_INFINITY,
		Max: unix.RLIM_INFINITY,
		})
		if err != nil {
			panic(err)
		}


        fmt.Println("NewModule")
        mod := elf.NewModule("ebpf/kprobe_example")

        fmt.Println("mod.Load")
        err = mod.Load(nil)
        if err != nil {
                panic(err)
        }

        fmt.Println("mod.EnableKprobes")
        err = mod.EnableKprobes(0)
        if err != nil {
				fmt.Println(err)
                panic(err)
        }

        fmt.Println("Loop")
        for {
                fmt.Println("Waiting...")
                time.Sleep(10 * time.Second)
        }}


