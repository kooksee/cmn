package cmn

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"os"
)

var OS = myOS{}

type myOS struct{}

// TrapSignal catches the SIGTERM and executes cb function. After that it exits
// with code 1.
func (myOS) TrapSignal(cb func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			fmt.Printf("captured %v, exiting...\n", sig)
			if cb != nil {
				cb()
			}
			os.Exit(1)
		}
	}()
	select {}
}

// Kill the running process by sending itself SIGTERM.
func (myOS) Kill() error {
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		return err
	}
	return p.Signal(syscall.SIGTERM)
}

func (myOS) Exit(s string) {
	fmt.Printf(s + "\n")
	os.Exit(1)
}

func (myOS) EnsureDir(dir string, mode os.FileMode) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, mode)
		if err != nil {
			return fmt.Errorf("Could not create directory %v. %v", dir, err)
		}
	}
	return nil
}

func (myOS) IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		if os.IsNotExist(err) {
			return true, err
		}
		// Otherwise perhaps a permission
		// error or some other error.
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

func (myOS) FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func (myOS) ReadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func (o myOS) MustReadFile(filePath string) []byte {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		o.Exit(F("MustReadFile failed: %v", err))
		return nil
	}
	return fileBytes
}

func (myOS) WriteFile(filePath string, contents []byte, mode os.FileMode) error {
	return ioutil.WriteFile(filePath, contents, mode)
}

func (o myOS) MustWriteFile(filePath string, contents []byte, mode os.FileMode) {
	err := o.WriteFile(filePath, contents, mode)
	if err != nil {
		o.Exit(F("MustWriteFile failed: %v", err))
	}
}

// WriteFileAtomic writes newBytes to temp and atomically moves to filePath
// when everything else succeeds.
func (myOS) WriteFileAtomic(filePath string, newBytes []byte, mode os.FileMode) error {
	dir := filepath.Dir(filePath)
	f, err := ioutil.TempFile(dir, "")
	if err != nil {
		return err
	}
	_, err = f.Write(newBytes)
	if err == nil {
		err = f.Sync()
	}
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	if permErr := os.Chmod(f.Name(), mode); err == nil {
		err = permErr
	}
	if err == nil {
		err = os.Rename(f.Name(), filePath)
	}
	// any err should result in full cleanup
	if err != nil {
		os.Remove(f.Name())
	}
	return err
}

//--------------------------------------------------------------------------------

func (myOS) Tempfile(prefix string) (*os.File, string) {
	file, err := ioutil.TempFile("", prefix)
	if err != nil {
		Err.MustNotErr(err)
	}
	return file, file.Name()
}

func (o myOS) Tempdir(prefix string) (*os.File, string) {
	tempDir := os.TempDir() + "/" + prefix + Rand.RandStr(12)
	err := o.EnsureDir(tempDir, 0700)
	if err != nil {
		panic(F("Error creating temp dir: %v", err))
	}
	dir, err := os.Open(tempDir)
	if err != nil {
		panic(F("Error opening temp dir: %v", err))
	}
	return dir, tempDir
}

//--------------------------------------------------------------------------------

func (myOS) Prompt(prompt string, defaultValue string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return defaultValue, err
	} else {
		line = strings.TrimSpace(line)
		if line == "" {
			return defaultValue, nil
		}
		return line, nil
	}
}
