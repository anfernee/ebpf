//go:build ignore

#include "common.h"

char __license[] SEC("license") = "Dual MIT/GPL";

SEC("kprobe/sys_execve")
int kprobe_execve() {
    return 0;
}

SEC("kprobe/sys_mount")
int kprobe_mount(struct pt_regs *ctx) {
    const char fmt[] = "call mount(dev_name=%s, dir_name=%s)";
    bpf_trace_printk(fmt, sizeof(fmt), "", "");

    return 0;
}
