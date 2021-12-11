package goretry

import (
	"testing"
	"time"
)

// Wait time to retry
var WaitTimes = []int{1, 1, 2, 2, 3, 4, 5}

// Retry the test.
// Write your test in a function that returns true on success.
//
// Example:
//	Retry(t, func () bool {
//		result := Example()
//		if result == "ok" {
//			return true
//		}
//		return false
//	}, "This test status message.")
func Retry(t *testing.T, f func() bool, message string) {
	if f() {
		return
	}

	isSuccess := false

	for i, wait := range WaitTimes {
		t.Logf("Try! count: %v, sleep: %vs", i+1, wait)
		time.Sleep(time.Duration(wait) * time.Second)
		if f() {
			isSuccess = true
			break
		}
	}

	if !isSuccess {
		t.Fatalf("Retry failed: %v", message)
	}
}
