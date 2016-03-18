// Copyright 2012-2016 Apcera Inc. All rights reserved.

package logray

import (
	"net/url"
	"testing"
)

const (
	formatString = "%color:class% %classfixed% " +
		"%year%-%month%-%day% %hour%:%minute%:%second%.%nanosecond% " +
		"%tzoffset% %tz% pid=%pid%"
)

func TestLoggerOutputs(t *testing.T) {

	// Create new URI for default stdout output.
	defaultUrlString := "stdout://?format=" + url.QueryEscape(formatString)
	// Append one extra field 'tid' to the url.
	newDefaultUrlString := defaultUrlString + url.QueryEscape(" tid='%field:tid%'")

	// Add default stdout output.
	if err := AddDefaultOutput(defaultUrlString, ALL); err != nil {
		t.Fatal(err)
	}

	logger := New()

	// Update for DEBUG should not update the output.
	logger.UpdateOutput(newDefaultUrlString, DEBUG)
	if len(logger.outputs) > 1 {
		t.Fatalf("More than 1 logger outputs defined: %d", len(logger.outputs))
	}
	output := logger.outputs[0]

	defaultUrl, err := url.Parse(defaultUrlString)
	if err != nil {
		t.Fatal(err)
	}

	newUrl, err := url.Parse(newDefaultUrlString)
	if err != nil {
		t.Fatal(err)
	}

	if output.OutputWrapper.URL.Scheme != defaultUrl.Scheme ||
		output.OutputWrapper.URL.RawQuery != defaultUrl.RawQuery {
		t.Fatalf("Output: %v", output.OutputWrapper)
	}

	// Now update for ALL.
	logger.UpdateOutput(newDefaultUrlString, ALL)

	if len(logger.outputs) > 1 {
		t.Fatalf("More than 1 logger outputs defined: %d", len(logger.outputs))
	}

	output = logger.outputs[0]
	// Must have new URL in the output.
	if output.OutputWrapper.URL.Scheme != newUrl.Scheme ||
		output.OutputWrapper.URL.RawQuery != newUrl.RawQuery {
		t.Fatalf("Output: %v", output.OutputWrapper.URL)
	}
}
