package main

import "testing"

func TestFormatInput(t *testing.T) {
	testInput := "asd     "
	expectedOutput := "ASD"

	formattedInput := formatInput(testInput)
	if formattedInput != expectedOutput {
		t.Errorf("Error in Format Input. Expected: %v. Got: %v.", expectedOutput, formattedInput)
	}
}

func TestValidInputValidation(t *testing.T) {
	testInput := "BMC"
	expectedOutput := true

	validatedInput, _ := validateInput(testInput)
	if validatedInput != expectedOutput {
		t.Errorf("Error in Validate Input. Expected: %v. Got: %v.", expectedOutput, validatedInput)
	}
}

func TestLongInputInputValidation(t *testing.T) {
	testInput := "GHFD"
	expectedOutput := false

	validatedInput, _ := validateInput(testInput)
	if validatedInput != expectedOutput {
		t.Errorf("Error in Validate Input. Expected: %v. Got: %v.", expectedOutput, validatedInput)
	}
}

func TestShortInputInputValidation(t *testing.T) {
	testInput := "GH"
	expectedOutput := false

	validatedInput, _ := validateInput(testInput)
	if validatedInput != expectedOutput {
		t.Errorf("Error in Validate Input. Expected: %v. Got: %v.", expectedOutput, validatedInput)
	}
}

func TestNoInputInputValidation(t *testing.T) {
	testInput := ""
	expectedOutput := false

	validatedInput, _ := validateInput(testInput)
	if validatedInput != expectedOutput {
		t.Errorf("Error in Validate Input. Expected: %v. Got: %v.", expectedOutput, validatedInput)
	}
}
