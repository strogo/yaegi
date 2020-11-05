// Code generated by 'yaegi extract syscall'. DO NOT EDIT.

// +build go1.15

package syscall

import (
	"go/constant"
	"go/token"
	"reflect"
	"syscall"
)

func init() {
	Symbols["syscall"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AF_INET":             reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"AF_INET6":            reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"AF_UNIX":             reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"AF_UNSPEC":           reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"Accept":              reflect.ValueOf(syscall.Accept),
		"Bind":                reflect.ValueOf(syscall.Bind),
		"BytePtrFromString":   reflect.ValueOf(syscall.BytePtrFromString),
		"ByteSliceFromString": reflect.ValueOf(syscall.ByteSliceFromString),
		"Chdir":               reflect.ValueOf(syscall.Chdir),
		"Chmod":               reflect.ValueOf(syscall.Chmod),
		"Chown":               reflect.ValueOf(syscall.Chown),
		"Clearenv":            reflect.ValueOf(syscall.Clearenv),
		"Close":               reflect.ValueOf(syscall.Close),
		"CloseOnExec":         reflect.ValueOf(syscall.CloseOnExec),
		"Connect":             reflect.ValueOf(syscall.Connect),
		"Dup":                 reflect.ValueOf(syscall.Dup),
		"Dup2":                reflect.ValueOf(syscall.Dup2),
		"E2BIG":               reflect.ValueOf(syscall.E2BIG),
		"EACCES":              reflect.ValueOf(syscall.EACCES),
		"EADDRINUSE":          reflect.ValueOf(syscall.EADDRINUSE),
		"EADDRNOTAVAIL":       reflect.ValueOf(syscall.EADDRNOTAVAIL),
		"EADV":                reflect.ValueOf(syscall.EADV),
		"EAFNOSUPPORT":        reflect.ValueOf(syscall.EAFNOSUPPORT),
		"EAGAIN":              reflect.ValueOf(syscall.EAGAIN),
		"EALREADY":            reflect.ValueOf(syscall.EALREADY),
		"EBADE":               reflect.ValueOf(syscall.EBADE),
		"EBADF":               reflect.ValueOf(syscall.EBADF),
		"EBADFD":              reflect.ValueOf(syscall.EBADFD),
		"EBADMSG":             reflect.ValueOf(syscall.EBADMSG),
		"EBADR":               reflect.ValueOf(syscall.EBADR),
		"EBADRQC":             reflect.ValueOf(syscall.EBADRQC),
		"EBADSLT":             reflect.ValueOf(syscall.EBADSLT),
		"EBFONT":              reflect.ValueOf(syscall.EBFONT),
		"EBUSY":               reflect.ValueOf(syscall.EBUSY),
		"ECANCELED":           reflect.ValueOf(syscall.ECANCELED),
		"ECASECLASH":          reflect.ValueOf(syscall.ECASECLASH),
		"ECHILD":              reflect.ValueOf(syscall.ECHILD),
		"ECHRNG":              reflect.ValueOf(syscall.ECHRNG),
		"ECOMM":               reflect.ValueOf(syscall.ECOMM),
		"ECONNABORTED":        reflect.ValueOf(syscall.ECONNABORTED),
		"ECONNREFUSED":        reflect.ValueOf(syscall.ECONNREFUSED),
		"ECONNRESET":          reflect.ValueOf(syscall.ECONNRESET),
		"EDEADLK":             reflect.ValueOf(syscall.EDEADLK),
		"EDEADLOCK":           reflect.ValueOf(syscall.EDEADLOCK),
		"EDESTADDRREQ":        reflect.ValueOf(syscall.EDESTADDRREQ),
		"EDOM":                reflect.ValueOf(syscall.EDOM),
		"EDOTDOT":             reflect.ValueOf(syscall.EDOTDOT),
		"EDQUOT":              reflect.ValueOf(syscall.EDQUOT),
		"EEXIST":              reflect.ValueOf(syscall.EEXIST),
		"EFAULT":              reflect.ValueOf(syscall.EFAULT),
		"EFBIG":               reflect.ValueOf(syscall.EFBIG),
		"EFTYPE":              reflect.ValueOf(syscall.EFTYPE),
		"EHOSTDOWN":           reflect.ValueOf(syscall.EHOSTDOWN),
		"EHOSTUNREACH":        reflect.ValueOf(syscall.EHOSTUNREACH),
		"EIDRM":               reflect.ValueOf(syscall.EIDRM),
		"EILSEQ":              reflect.ValueOf(syscall.EILSEQ),
		"EINPROGRESS":         reflect.ValueOf(syscall.EINPROGRESS),
		"EINTR":               reflect.ValueOf(syscall.EINTR),
		"EINVAL":              reflect.ValueOf(syscall.EINVAL),
		"EIO":                 reflect.ValueOf(syscall.EIO),
		"EISCONN":             reflect.ValueOf(syscall.EISCONN),
		"EISDIR":              reflect.ValueOf(syscall.EISDIR),
		"EL2HLT":              reflect.ValueOf(syscall.EL2HLT),
		"EL2NSYNC":            reflect.ValueOf(syscall.EL2NSYNC),
		"EL3HLT":              reflect.ValueOf(syscall.EL3HLT),
		"EL3RST":              reflect.ValueOf(syscall.EL3RST),
		"ELBIN":               reflect.ValueOf(syscall.ELBIN),
		"ELIBACC":             reflect.ValueOf(syscall.ELIBACC),
		"ELIBBAD":             reflect.ValueOf(syscall.ELIBBAD),
		"ELIBEXEC":            reflect.ValueOf(syscall.ELIBEXEC),
		"ELIBMAX":             reflect.ValueOf(syscall.ELIBMAX),
		"ELIBSCN":             reflect.ValueOf(syscall.ELIBSCN),
		"ELNRNG":              reflect.ValueOf(syscall.ELNRNG),
		"ELOOP":               reflect.ValueOf(syscall.ELOOP),
		"EMFILE":              reflect.ValueOf(syscall.EMFILE),
		"EMLINK":              reflect.ValueOf(syscall.EMLINK),
		"EMSGSIZE":            reflect.ValueOf(syscall.EMSGSIZE),
		"EMULTIHOP":           reflect.ValueOf(syscall.EMULTIHOP),
		"ENAMETOOLONG":        reflect.ValueOf(syscall.ENAMETOOLONG),
		"ENETDOWN":            reflect.ValueOf(syscall.ENETDOWN),
		"ENETRESET":           reflect.ValueOf(syscall.ENETRESET),
		"ENETUNREACH":         reflect.ValueOf(syscall.ENETUNREACH),
		"ENFILE":              reflect.ValueOf(syscall.ENFILE),
		"ENMFILE":             reflect.ValueOf(syscall.ENMFILE),
		"ENOANO":              reflect.ValueOf(syscall.ENOANO),
		"ENOBUFS":             reflect.ValueOf(syscall.ENOBUFS),
		"ENOCSI":              reflect.ValueOf(syscall.ENOCSI),
		"ENODATA":             reflect.ValueOf(syscall.ENODATA),
		"ENODEV":              reflect.ValueOf(syscall.ENODEV),
		"ENOENT":              reflect.ValueOf(syscall.ENOENT),
		"ENOEXEC":             reflect.ValueOf(syscall.ENOEXEC),
		"ENOLCK":              reflect.ValueOf(syscall.ENOLCK),
		"ENOLINK":             reflect.ValueOf(syscall.ENOLINK),
		"ENOMEDIUM":           reflect.ValueOf(syscall.ENOMEDIUM),
		"ENOMEM":              reflect.ValueOf(syscall.ENOMEM),
		"ENOMSG":              reflect.ValueOf(syscall.ENOMSG),
		"ENONET":              reflect.ValueOf(syscall.ENONET),
		"ENOPKG":              reflect.ValueOf(syscall.ENOPKG),
		"ENOPROTOOPT":         reflect.ValueOf(syscall.ENOPROTOOPT),
		"ENOSHARE":            reflect.ValueOf(syscall.ENOSHARE),
		"ENOSPC":              reflect.ValueOf(syscall.ENOSPC),
		"ENOSR":               reflect.ValueOf(syscall.ENOSR),
		"ENOSTR":              reflect.ValueOf(syscall.ENOSTR),
		"ENOSYS":              reflect.ValueOf(syscall.ENOSYS),
		"ENOTCONN":            reflect.ValueOf(syscall.ENOTCONN),
		"ENOTDIR":             reflect.ValueOf(syscall.ENOTDIR),
		"ENOTEMPTY":           reflect.ValueOf(syscall.ENOTEMPTY),
		"ENOTSOCK":            reflect.ValueOf(syscall.ENOTSOCK),
		"ENOTSUP":             reflect.ValueOf(syscall.ENOTSUP),
		"ENOTTY":              reflect.ValueOf(syscall.ENOTTY),
		"ENOTUNIQ":            reflect.ValueOf(syscall.ENOTUNIQ),
		"ENXIO":               reflect.ValueOf(syscall.ENXIO),
		"EOPNOTSUPP":          reflect.ValueOf(syscall.EOPNOTSUPP),
		"EOVERFLOW":           reflect.ValueOf(syscall.EOVERFLOW),
		"EPERM":               reflect.ValueOf(syscall.EPERM),
		"EPFNOSUPPORT":        reflect.ValueOf(syscall.EPFNOSUPPORT),
		"EPIPE":               reflect.ValueOf(syscall.EPIPE),
		"EPROCLIM":            reflect.ValueOf(syscall.EPROCLIM),
		"EPROTO":              reflect.ValueOf(syscall.EPROTO),
		"EPROTONOSUPPORT":     reflect.ValueOf(syscall.EPROTONOSUPPORT),
		"EPROTOTYPE":          reflect.ValueOf(syscall.EPROTOTYPE),
		"ERANGE":              reflect.ValueOf(syscall.ERANGE),
		"EREMCHG":             reflect.ValueOf(syscall.EREMCHG),
		"EREMOTE":             reflect.ValueOf(syscall.EREMOTE),
		"EROFS":               reflect.ValueOf(syscall.EROFS),
		"ESHUTDOWN":           reflect.ValueOf(syscall.ESHUTDOWN),
		"ESOCKTNOSUPPORT":     reflect.ValueOf(syscall.ESOCKTNOSUPPORT),
		"ESPIPE":              reflect.ValueOf(syscall.ESPIPE),
		"ESRCH":               reflect.ValueOf(syscall.ESRCH),
		"ESRMNT":              reflect.ValueOf(syscall.ESRMNT),
		"ESTALE":              reflect.ValueOf(syscall.ESTALE),
		"ETIME":               reflect.ValueOf(syscall.ETIME),
		"ETIMEDOUT":           reflect.ValueOf(syscall.ETIMEDOUT),
		"ETOOMANYREFS":        reflect.ValueOf(syscall.ETOOMANYREFS),
		"EUNATCH":             reflect.ValueOf(syscall.EUNATCH),
		"EUSERS":              reflect.ValueOf(syscall.EUSERS),
		"EWOULDBLOCK":         reflect.ValueOf(syscall.EWOULDBLOCK),
		"EXDEV":               reflect.ValueOf(syscall.EXDEV),
		"EXFULL":              reflect.ValueOf(syscall.EXFULL),
		"Environ":             reflect.ValueOf(syscall.Environ),
		"F_CNVT":              reflect.ValueOf(constant.MakeFromLiteral("12", token.INT, 0)),
		"F_DUPFD":             reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"F_DUPFD_CLOEXEC":     reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"F_GETFD":             reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"F_GETFL":             reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"F_GETLK":             reflect.ValueOf(constant.MakeFromLiteral("7", token.INT, 0)),
		"F_GETOWN":            reflect.ValueOf(constant.MakeFromLiteral("5", token.INT, 0)),
		"F_RDLCK":             reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"F_RGETLK":            reflect.ValueOf(constant.MakeFromLiteral("10", token.INT, 0)),
		"F_RSETLK":            reflect.ValueOf(constant.MakeFromLiteral("11", token.INT, 0)),
		"F_RSETLKW":           reflect.ValueOf(constant.MakeFromLiteral("13", token.INT, 0)),
		"F_SETFD":             reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"F_SETFL":             reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"F_SETLK":             reflect.ValueOf(constant.MakeFromLiteral("8", token.INT, 0)),
		"F_SETLKW":            reflect.ValueOf(constant.MakeFromLiteral("9", token.INT, 0)),
		"F_SETOWN":            reflect.ValueOf(constant.MakeFromLiteral("6", token.INT, 0)),
		"F_UNLCK":             reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"F_UNLKSYS":           reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"F_WRLCK":             reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"Fchdir":              reflect.ValueOf(syscall.Fchdir),
		"Fchmod":              reflect.ValueOf(syscall.Fchmod),
		"Fchown":              reflect.ValueOf(syscall.Fchown),
		"ForkLock":            reflect.ValueOf(&syscall.ForkLock).Elem(),
		"Fstat":               reflect.ValueOf(syscall.Fstat),
		"Fsync":               reflect.ValueOf(syscall.Fsync),
		"Ftruncate":           reflect.ValueOf(syscall.Ftruncate),
		"Getcwd":              reflect.ValueOf(syscall.Getcwd),
		"Getegid":             reflect.ValueOf(syscall.Getegid),
		"Getenv":              reflect.ValueOf(syscall.Getenv),
		"Geteuid":             reflect.ValueOf(syscall.Geteuid),
		"Getgid":              reflect.ValueOf(syscall.Getgid),
		"Getgroups":           reflect.ValueOf(syscall.Getgroups),
		"Getpagesize":         reflect.ValueOf(syscall.Getpagesize),
		"Getpid":              reflect.ValueOf(syscall.Getpid),
		"Getppid":             reflect.ValueOf(syscall.Getppid),
		"GetsockoptInt":       reflect.ValueOf(syscall.GetsockoptInt),
		"Gettimeofday":        reflect.ValueOf(syscall.Gettimeofday),
		"Getuid":              reflect.ValueOf(syscall.Getuid),
		"Getwd":               reflect.ValueOf(syscall.Getwd),
		"IPPROTO_IP":          reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"IPPROTO_IPV4":        reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"IPPROTO_IPV6":        reflect.ValueOf(constant.MakeFromLiteral("41", token.INT, 0)),
		"IPPROTO_TCP":         reflect.ValueOf(constant.MakeFromLiteral("6", token.INT, 0)),
		"IPPROTO_UDP":         reflect.ValueOf(constant.MakeFromLiteral("17", token.INT, 0)),
		"IPV6_V6ONLY":         reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"ImplementsGetwd":     reflect.ValueOf(syscall.ImplementsGetwd),
		"Lchown":              reflect.ValueOf(syscall.Lchown),
		"Link":                reflect.ValueOf(syscall.Link),
		"Listen":              reflect.ValueOf(syscall.Listen),
		"Lstat":               reflect.ValueOf(syscall.Lstat),
		"Mkdir":               reflect.ValueOf(syscall.Mkdir),
		"NsecToTimespec":      reflect.ValueOf(syscall.NsecToTimespec),
		"NsecToTimeval":       reflect.ValueOf(syscall.NsecToTimeval),
		"O_APPEND":            reflect.ValueOf(constant.MakeFromLiteral("1024", token.INT, 0)),
		"O_CLOEXEC":           reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"O_CREAT":             reflect.ValueOf(constant.MakeFromLiteral("64", token.INT, 0)),
		"O_CREATE":            reflect.ValueOf(constant.MakeFromLiteral("64", token.INT, 0)),
		"O_EXCL":              reflect.ValueOf(constant.MakeFromLiteral("128", token.INT, 0)),
		"O_RDONLY":            reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"O_RDWR":              reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"O_SYNC":              reflect.ValueOf(constant.MakeFromLiteral("4096", token.INT, 0)),
		"O_TRUNC":             reflect.ValueOf(constant.MakeFromLiteral("512", token.INT, 0)),
		"O_WRONLY":            reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"Open":                reflect.ValueOf(syscall.Open),
		"ParseDirent":         reflect.ValueOf(syscall.ParseDirent),
		"PathMax":             reflect.ValueOf(constant.MakeFromLiteral("256", token.INT, 0)),
		"Pipe":                reflect.ValueOf(syscall.Pipe),
		"Pread":               reflect.ValueOf(syscall.Pread),
		"Pwrite":              reflect.ValueOf(syscall.Pwrite),
		"Read":                reflect.ValueOf(syscall.Read),
		"ReadDirent":          reflect.ValueOf(syscall.ReadDirent),
		"Readlink":            reflect.ValueOf(syscall.Readlink),
		"Recvfrom":            reflect.ValueOf(syscall.Recvfrom),
		"Recvmsg":             reflect.ValueOf(syscall.Recvmsg),
		"Rename":              reflect.ValueOf(syscall.Rename),
		"Rmdir":               reflect.ValueOf(syscall.Rmdir),
		"SIGCHLD":             reflect.ValueOf(syscall.SIGCHLD),
		"SIGINT":              reflect.ValueOf(syscall.SIGINT),
		"SIGKILL":             reflect.ValueOf(syscall.SIGKILL),
		"SIGQUIT":             reflect.ValueOf(syscall.SIGQUIT),
		"SIGTERM":             reflect.ValueOf(syscall.SIGTERM),
		"SIGTRAP":             reflect.ValueOf(syscall.SIGTRAP),
		"SOCK_DGRAM":          reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"SOCK_RAW":            reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"SOCK_SEQPACKET":      reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"SOCK_STREAM":         reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"SOMAXCONN":           reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"SO_ERROR":            reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"SYS_FCNTL":           reflect.ValueOf(constant.MakeFromLiteral("500", token.INT, 0)),
		"S_IEXEC":             reflect.ValueOf(constant.MakeFromLiteral("64", token.INT, 0)),
		"S_IFBLK":             reflect.ValueOf(constant.MakeFromLiteral("24576", token.INT, 0)),
		"S_IFBOUNDSOCK":       reflect.ValueOf(constant.MakeFromLiteral("77824", token.INT, 0)),
		"S_IFCHR":             reflect.ValueOf(constant.MakeFromLiteral("8192", token.INT, 0)),
		"S_IFCOND":            reflect.ValueOf(constant.MakeFromLiteral("90112", token.INT, 0)),
		"S_IFDIR":             reflect.ValueOf(constant.MakeFromLiteral("16384", token.INT, 0)),
		"S_IFDSOCK":           reflect.ValueOf(constant.MakeFromLiteral("69632", token.INT, 0)),
		"S_IFIFO":             reflect.ValueOf(constant.MakeFromLiteral("4096", token.INT, 0)),
		"S_IFLNK":             reflect.ValueOf(constant.MakeFromLiteral("40960", token.INT, 0)),
		"S_IFMT":              reflect.ValueOf(constant.MakeFromLiteral("126976", token.INT, 0)),
		"S_IFMUTEX":           reflect.ValueOf(constant.MakeFromLiteral("86016", token.INT, 0)),
		"S_IFREG":             reflect.ValueOf(constant.MakeFromLiteral("32768", token.INT, 0)),
		"S_IFSEMA":            reflect.ValueOf(constant.MakeFromLiteral("94208", token.INT, 0)),
		"S_IFSHM":             reflect.ValueOf(constant.MakeFromLiteral("81920", token.INT, 0)),
		"S_IFSHM_SYSV":        reflect.ValueOf(constant.MakeFromLiteral("98304", token.INT, 0)),
		"S_IFSOCK":            reflect.ValueOf(constant.MakeFromLiteral("49152", token.INT, 0)),
		"S_IFSOCKADDR":        reflect.ValueOf(constant.MakeFromLiteral("73728", token.INT, 0)),
		"S_IREAD":             reflect.ValueOf(constant.MakeFromLiteral("256", token.INT, 0)),
		"S_IRGRP":             reflect.ValueOf(constant.MakeFromLiteral("32", token.INT, 0)),
		"S_IROTH":             reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"S_IRUSR":             reflect.ValueOf(constant.MakeFromLiteral("256", token.INT, 0)),
		"S_IRWXG":             reflect.ValueOf(constant.MakeFromLiteral("56", token.INT, 0)),
		"S_IRWXO":             reflect.ValueOf(constant.MakeFromLiteral("7", token.INT, 0)),
		"S_IRWXU":             reflect.ValueOf(constant.MakeFromLiteral("448", token.INT, 0)),
		"S_ISGID":             reflect.ValueOf(constant.MakeFromLiteral("1024", token.INT, 0)),
		"S_ISUID":             reflect.ValueOf(constant.MakeFromLiteral("2048", token.INT, 0)),
		"S_ISVTX":             reflect.ValueOf(constant.MakeFromLiteral("512", token.INT, 0)),
		"S_IWGRP":             reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"S_IWOTH":             reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"S_IWRITE":            reflect.ValueOf(constant.MakeFromLiteral("128", token.INT, 0)),
		"S_IWUSR":             reflect.ValueOf(constant.MakeFromLiteral("128", token.INT, 0)),
		"S_IXGRP":             reflect.ValueOf(constant.MakeFromLiteral("8", token.INT, 0)),
		"S_IXOTH":             reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"S_IXUSR":             reflect.ValueOf(constant.MakeFromLiteral("64", token.INT, 0)),
		"S_UNSUP":             reflect.ValueOf(constant.MakeFromLiteral("126976", token.INT, 0)),
		"Seek":                reflect.ValueOf(syscall.Seek),
		"Sendfile":            reflect.ValueOf(syscall.Sendfile),
		"SendmsgN":            reflect.ValueOf(syscall.SendmsgN),
		"Sendto":              reflect.ValueOf(syscall.Sendto),
		"SetNonblock":         reflect.ValueOf(syscall.SetNonblock),
		"SetReadDeadline":     reflect.ValueOf(syscall.SetReadDeadline),
		"SetWriteDeadline":    reflect.ValueOf(syscall.SetWriteDeadline),
		"Setenv":              reflect.ValueOf(syscall.Setenv),
		"SetsockoptInt":       reflect.ValueOf(syscall.SetsockoptInt),
		"Socket":              reflect.ValueOf(syscall.Socket),
		"Stat":                reflect.ValueOf(syscall.Stat),
		"Stderr":              reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"Stdin":               reflect.ValueOf(constant.MakeFromLiteral("0", token.INT, 0)),
		"Stdout":              reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"StopIO":              reflect.ValueOf(syscall.StopIO),
		"StringBytePtr":       reflect.ValueOf(syscall.StringBytePtr),
		"StringByteSlice":     reflect.ValueOf(syscall.StringByteSlice),
		"Symlink":             reflect.ValueOf(syscall.Symlink),
		"Sysctl":              reflect.ValueOf(syscall.Sysctl),
		"TimespecToNsec":      reflect.ValueOf(syscall.TimespecToNsec),
		"TimevalToNsec":       reflect.ValueOf(syscall.TimevalToNsec),
		"Truncate":            reflect.ValueOf(syscall.Truncate),
		"Umask":               reflect.ValueOf(syscall.Umask),
		"Unlink":              reflect.ValueOf(syscall.Unlink),
		"Unsetenv":            reflect.ValueOf(syscall.Unsetenv),
		"UtimesNano":          reflect.ValueOf(syscall.UtimesNano),
		"Wait4":               reflect.ValueOf(syscall.Wait4),
		"Write":               reflect.ValueOf(syscall.Write),

		// type definitions
		"Conn":          reflect.ValueOf((*syscall.Conn)(nil)),
		"Dirent":        reflect.ValueOf((*syscall.Dirent)(nil)),
		"Errno":         reflect.ValueOf((*syscall.Errno)(nil)),
		"Iovec":         reflect.ValueOf((*syscall.Iovec)(nil)),
		"ProcAttr":      reflect.ValueOf((*syscall.ProcAttr)(nil)),
		"RawConn":       reflect.ValueOf((*syscall.RawConn)(nil)),
		"Rusage":        reflect.ValueOf((*syscall.Rusage)(nil)),
		"Signal":        reflect.ValueOf((*syscall.Signal)(nil)),
		"Sockaddr":      reflect.ValueOf((*syscall.Sockaddr)(nil)),
		"SockaddrInet4": reflect.ValueOf((*syscall.SockaddrInet4)(nil)),
		"SockaddrInet6": reflect.ValueOf((*syscall.SockaddrInet6)(nil)),
		"SockaddrUnix":  reflect.ValueOf((*syscall.SockaddrUnix)(nil)),
		"Stat_t":        reflect.ValueOf((*syscall.Stat_t)(nil)),
		"SysProcAttr":   reflect.ValueOf((*syscall.SysProcAttr)(nil)),
		"Timespec":      reflect.ValueOf((*syscall.Timespec)(nil)),
		"Timeval":       reflect.ValueOf((*syscall.Timeval)(nil)),
		"WaitStatus":    reflect.ValueOf((*syscall.WaitStatus)(nil)),

		// interface wrapper definitions
		"_Conn":     reflect.ValueOf((*_syscall_Conn)(nil)),
		"_RawConn":  reflect.ValueOf((*_syscall_RawConn)(nil)),
		"_Sockaddr": reflect.ValueOf((*_syscall_Sockaddr)(nil)),
	}
}

// _syscall_Conn is an interface wrapper for Conn type
type _syscall_Conn struct {
	WSyscallConn func() (syscall.RawConn, error)
}

func (W _syscall_Conn) SyscallConn() (syscall.RawConn, error) { return W.WSyscallConn() }

// _syscall_RawConn is an interface wrapper for RawConn type
type _syscall_RawConn struct {
	WControl func(f func(fd uintptr)) error
	WRead    func(f func(fd uintptr) (done bool)) error
	WWrite   func(f func(fd uintptr) (done bool)) error
}

func (W _syscall_RawConn) Control(f func(fd uintptr)) error           { return W.WControl(f) }
func (W _syscall_RawConn) Read(f func(fd uintptr) (done bool)) error  { return W.WRead(f) }
func (W _syscall_RawConn) Write(f func(fd uintptr) (done bool)) error { return W.WWrite(f) }

// _syscall_Sockaddr is an interface wrapper for Sockaddr type
type _syscall_Sockaddr struct {
}
