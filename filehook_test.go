package main

import (
	"fmt"
	"testing"

	logger2 "github.com/nori-io/logger"
	logger "github.com/nori-io/logger/examples/hooks/filehook"
)

func TestPlugin_Instance(t *testing.T) {
	hook, err := logger.NewFileHookTest("","file_test")
	if err != nil {
		t.Errorf("Can't create hook")
	}
	testPlugin := plugin{log: logger2.New(), instance: hook}
	testPlugin.log.AddHook(hook)
	fmt.Println(testPlugin.instance.Levels())
	fmt.Println("testPlugin", testPlugin.Instance())
}
