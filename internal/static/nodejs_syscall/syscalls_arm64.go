//go:build linux && arm64

package nodejs_syscall

import "syscall"

var ALLOW_SYSCALLS = []int{
	// file
	syscall.SYS_CLOSE, syscall.SYS_WRITE,
	syscall.SYS_FSTAT, syscall.SYS_FCNTL,
	syscall.SYS_READLINKAT,

	// process
	syscall.SYS_GETPID, syscall.SYS_TGKILL, syscall.SYS_FUTEX, syscall.SYS_EXIT_GROUP,

	// memory
	syscall.SYS_RT_SIGPROCMASK, syscall.SYS_SIGALTSTACK, syscall.SYS_RT_SIGACTION,
	syscall.SYS_MMAP, syscall.SYS_MUNMAP, syscall.SYS_MADVISE, syscall.SYS_MPROTECT,

	//user/group
	syscall.SYS_SETUID, syscall.SYS_SETGID,
	syscall.SYS_GETUID, syscall.SYS_GETGID,

	// epoll
	syscall.SYS_EPOLL_CTL, syscall.SYS_EPOLL_PWAIT,
}

var ERROR_CODE_SYSCALLS = []int{
	// file io
	syscall.SYS_OPENAT, syscall.SYS_FSTATAT, syscall.SYS_IOCTL, syscall.SYS_READ, syscall.SYS_LSEEK,
}

var ALLOW_NETWORK_SYSCALLS = []int{
	syscall.SYS_SOCKET, syscall.SYS_CONNECT, syscall.SYS_BIND, syscall.SYS_LISTEN, syscall.SYS_ACCEPT, syscall.SYS_SENDTO, syscall.SYS_RECVFROM,
	syscall.SYS_GETSOCKNAME, syscall.SYS_GETPEERNAME,
}