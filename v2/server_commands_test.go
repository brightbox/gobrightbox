// Code generated by generate_server_commands; DO NOT EDIT.

package brightbox

import "path"
import "testing"

func TestStartServer(t *testing.T) {
	testCommand(
		t,
		(*Client).StartServer,
		"srv-lv426",
		"POST",
		path.Join("servers", "srv-lv426", "start"),
	)
}

func TestStopServer(t *testing.T) {
	testCommand(
		t,
		(*Client).StopServer,
		"srv-lv426",
		"POST",
		path.Join("servers", "srv-lv426", "stop"),
	)
}

func TestRebootServer(t *testing.T) {
	testCommand(
		t,
		(*Client).RebootServer,
		"srv-lv426",
		"POST",
		path.Join("servers", "srv-lv426", "reboot"),
	)
}

func TestResetServer(t *testing.T) {
	testCommand(
		t,
		(*Client).ResetServer,
		"srv-lv426",
		"POST",
		path.Join("servers", "srv-lv426", "reset"),
	)
}

func TestShutdownServer(t *testing.T) {
	testCommand(
		t,
		(*Client).ShutdownServer,
		"srv-lv426",
		"POST",
		path.Join("servers", "srv-lv426", "shutdown"),
	)
}