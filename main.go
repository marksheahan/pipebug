package main

import (
	"io"
	"os/exec"
)

func newPipesSession(cmd string) (*exec.Cmd, io.WriteCloser, io.Reader, io.Reader, error) {
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

	return sess, stdin_pipe, stdout_pipe, stderr_pipe, nil
}

func runCommandStringPipes(cmd string) (io.WriteCloser, io.Reader, io.Reader, chan error, error) {
	c, stdin, stdout, stderr, err := newPipesSession(cmd)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	err = c.Start()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	finished := make(chan error)
	go func() {
		defer close(finished)
		finished <- c.Wait()
	}()
	return stdin, stdout, stderr, finished, nil
}
