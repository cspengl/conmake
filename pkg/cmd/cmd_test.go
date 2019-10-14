package cmd

import "testing"

func TestRootCmdExecution(t *testing.T) {
	err := ConmakeCmd.Execute()

	if err != nil {
		t.Fatal(err)
	}
}
