package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cilium/ebpf/link"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc $BPF_CLANG -cflags $BPF_CFLAGS bpf probe.c -- -I../headers

func main() {
	fmt.Println(hello)

	objs := bpfObjects{}
	if err := loadBpfObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()

	kp, err := link.Kprobe("sys_mount", objs.KprobeMount, nil)
	if err != nil {
		log.Fatalf("opening kprobe: %v", err)
	}
	defer kp.Close()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("still waiting...")
	}

}
