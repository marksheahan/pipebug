package main

import (
	"bytes"
	"io"
	"sync"
	"testing"
)

func TestRunCommandStringPipes(t *testing.T) {

	stdin, stdout, stderr, finished, err := runCommandStringPipes("cat ; echo tostderr >&2")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("stdout %T %v %#v stderr %T %v %#v", stdout, stdout, stdout, stderr, stderr, stderr)

	stdoutExpected := "hello world"
	stdinBuf := bytes.NewBufferString(stdoutExpected)
	stderrExpected := "tostderr\n"
	stdoutBuf := &bytes.Buffer{}
	stderrBuf := &bytes.Buffer{}
	errors := make(chan error, 5)

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		if _, err := io.Copy(stdin, stdinBuf); err != nil {
			t.Logf("err: %v", err)
			errors <- err
		}
		if err := stdin.Close(); err != nil {
			t.Logf("err: %v", err)
			errors <- err
		}
	}()

	go func() {
		defer wg.Done()
		if _, err := io.Copy(stdoutBuf, stdout); err != nil {
			t.Logf("stdout io.Copy err: %v", err)
			errors <- err
		}
	}()
	go func() {
		defer wg.Done()
		if _, err := io.Copy(stderrBuf, stderr); err != nil {
			t.Logf("stderr io.Copy err: %v", err)
			errors <- err
		}
	}()

	if err := <-finished; err != nil {
		t.Fatal(err)
	}
	wg.Wait()
	close(errors)

	for err := range errors {
		t.Fatal(err)
	}

	if stdoutBuf.String() != stdoutExpected {
		t.Fatalf("stdout data died %v vs %v", stdoutBuf.String(), stdoutExpected)
	}
	if stderrBuf.String() != stderrExpected {
		t.Fatalf("stderr data died %v vs %v", stderrBuf.String(), stderrExpected)
	}
}
