package main

import (
	"io"
	"os"
	"os/exec"
)

type Logger interface {
	Logf(string, ...interface{})
}

func runCommandStringPipes(cmd string, t Logger) (io.WriteCloser, io.Reader, io.Reader, chan error, error) {
	sess := exec.Command("/bin/sh", "-c", cmd)

	stdin_pipe, err := sess.StdinPipe()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	stdout_pipe, err := sess.StdoutPipe()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	stderr_pipe, err := sess.StderrPipe()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	t.Logf("file descriptors: stdout %v %v stderr %v %v",
		stdout_pipe.(*os.File).Fd(),
		sess.Stdout.(*os.File).Fd(),
		stderr_pipe.(*os.File).Fd(),
		sess.Stderr.(*os.File).Fd(),
	)

	err = sess.Start()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	finished := make(chan error)
	go func() {
		defer close(finished)
		finished <- sess.Wait()
	}()
	return stdin_pipe, stdout_pipe, stderr_pipe, finished, nil
}
