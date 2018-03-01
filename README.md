# pipebug
possible golang bug / api race condition

Cannot reproduce on Darwin:
```
Darwin mark-sheahan 17.4.0 Darwin Kernel Version 17.4.0: Sun Dec 17 09:19:54 PST 2017; root:xnu-4570.41.2~1/RELEASE_X86_64 x86_64
```

Can reproduce on Linux with several go releases (1.8.3, 1.10), and several Linux kernels:
```
Linux centos-vm 3.10.0-514.26.2.el7.x86_64 #1 SMP Tue Jul 4 15:04:05 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux
Linux ubuntu-vm 3.13.0-133-generic #182-Ubuntu SMP Tue Sep 19 15:49:21 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux
Linux vagrant 4.13.0-19-generic #22-Ubuntu SMP Mon Dec 4 11:58:07 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux
```

Just checkout and type `make` a few times to reproduce:
```bash
vagrant@vagrant:~/go/src/github.com/marksheahan/pipebug$ make
go test -count 100
--- FAIL: TestRunCommandStringPipes (0.00s)
	main.go:29: file descriptors: stdout 6 7 stderr 8 9
	main_test.go:18: stdout *os.File &{0xc4200d56d0} 6 stderr *os.File &{0xc4200d5770} 8
	main_test.go:51: stderr io.Copy err: read |0: file already closed
	main_test.go:69: read |0: file already closed
--- FAIL: TestRunCommandStringPipes (0.01s)
	main.go:29: file descriptors: stdout 6 7 stderr 8 9
	main_test.go:18: stdout *os.File &{0xc4201d8eb0} 6 stderr *os.File &{0xc4201d8f50} 8
	main_test.go:44: stdout io.Copy err: read |0: file already closed
	main_test.go:51: stderr io.Copy err: read |0: file already closed
	main_test.go:69: read |0: file already closed
	main_test.go:69: read |0: file already closed
--- FAIL: TestRunCommandStringPipes (0.01s)
	main.go:29: file descriptors: stdout 6 7 stderr 8 9
	main_test.go:18: stdout *os.File &{0xc4201d9310} 6 stderr *os.File &{0xc4201d93b0} 8
	main_test.go:51: stderr io.Copy err: read |0: file already closed
	main_test.go:69: read |0: file already closed
--- FAIL: TestRunCommandStringPipes (0.00s)
	main.go:29: file descriptors: stdout 6 7 stderr 8 9
	main_test.go:18: stdout *os.File &{0xc42021b0e0} 6 stderr *os.File &{0xc42021b180} 8
	main_test.go:51: stderr io.Copy err: read |0: file already closed
	main_test.go:69: read |0: file already closed
FAIL
exit status 1
FAIL	github.com/marksheahan/pipebug	0.226s
Makefile:4: recipe for target 'test' failed
make: *** [test] Error 1
vagrant@vagrant:~/go/src/github.com/marksheahan/pipebug$ go version
go version go1.10 linux/amd64
vagrant@vagrant:~/go/src/github.com/marksheahan/pipebug$ uname -a
Linux vagrant 4.13.0-19-generic #22-Ubuntu SMP Mon Dec 4 11:58:07 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux
vagrant@vagrant:~/go/src/github.com/marksheahan/pipebug$ cat /etc/os-release 
NAME="Ubuntu"
VERSION="17.10 (Artful Aardvark)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 17.10"
VERSION_ID="17.10"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
VERSION_CODENAME=artful
UBUNTU_CODENAME=artful
```

