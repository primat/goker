package goker

import (
	"testing"
	"strconv"
)

func TestMakeCard(t *testing.T) {

	// Test two of clubs
	result := MakeCard(0)
	var expected uint32 = 0x41
	if uint32(result) != expected {
		t.Fatalf("Expected 0x%08v, got 0x%08v", strconv.FormatUint(uint64(expected), 2), strconv.FormatUint(uint64(result), 2))
	}

	// Test ace of clubs
	result = MakeCard(12)
	expected = 0x40000001
	if uint32(result) != expected {
		t.Fatalf("Expected 0x%08v, got 0x%08v", strconv.FormatUint(uint64(expected), 2), strconv.FormatUint(uint64(result), 2))
	}

	// Test three of diamonds
	result = MakeCard(14)
	expected = 0x00000102
	if uint32(result) != expected {
		t.Fatalf("Expected 0x%08v, got 0x%08v", strconv.FormatUint(uint64(expected), 2), strconv.FormatUint(uint64(result), 2))
	}

	// Test king of diamonds
	result = MakeCard(24)
	expected = 0x10000002
	if uint32(result) != expected {
		t.Fatalf("Expected 0x%08v, got 0x%08v", strconv.FormatUint(uint64(expected), 2), strconv.FormatUint(uint64(result), 2))
	}

	// Test four of hearts
	result = MakeCard(28)
	expected = 0x00000404
	if uint32(result) != expected {
		t.Fatalf("Expected 0x%08v, got 0x%08v", strconv.FormatUint(uint64(expected), 2), strconv.FormatUint(uint64(result), 2))
	}

	// Test queen of hearts
	result = MakeCard(36)
	expected = 0x04000004
	if uint32(result) != expected {
		t.Fatalf("Expected 0x%08v, got 0x%08v", strconv.FormatUint(uint64(expected), 2), strconv.FormatUint(uint64(result), 2))
	}

	// Test five of spades
	result = MakeCard(43)
	expected = 0x00004008
	if uint32(result) != expected {
		t.Fatalf("Expected 0x%08v, got 0x%08v", strconv.FormatUint(uint64(expected), 2), strconv.FormatUint(uint64(result), 2))
	}

	// Test jack of spades
	result = MakeCard(48)
	expected = 0x01000008
	if uint32(result) != expected {
		t.Fatalf("Expected 0x%08v, got 0x%08v", strconv.FormatUint(uint64(expected), 2), strconv.FormatUint(uint64(result), 2))
	}
}
