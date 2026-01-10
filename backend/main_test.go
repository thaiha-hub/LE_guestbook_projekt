package main

import "testing"

// This is the Unit Test
func TestValidateEntry(t *testing.T) {
    // Scenario 1: Happy Path (Valid Input)
    err := ValidateEntry("Thai", "Hello World")
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    // Scenario 2: Bad Path (Empty Name)
    err = ValidateEntry("", "Hello World")
    if err == nil {
        t.Errorf("Expected error for empty name, but got nil")
    }
}