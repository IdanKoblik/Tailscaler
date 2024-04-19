package main

import (
	"io"
	"os"
	"testing"
)

func TestTailscaler(t *testing.T) {
	// Redirect standard output to capture printed output
	old := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	main()

	// Restore standard output
	w.Close()
	os.Stdout = old

	// Read captured output
	out, _ := io.ReadAll(r)

	expected := "$$$$$$$$\\  $$$$$$\\  $$$$$$\\ $$\\       $$$$$$\\   $$$$$$\\   $$$$$$\\  $$\\       $$$$$$$$\\ $$$$$$$\\\n\\__$$  __|$$  __$$\\ \\_$$  _|$$ |     $$  __$$\\ $$  __$$\\ $$  __$$\\ $$ |      $$  _____|$$  __$$\\\n   $$ |   $$ /  $$ |  $$ |  $$ |     $$ /  \\__|$$ /  \\__|$$ /  $$ |$$ |      $$ |      $$ |  $$ |\n   $$ |   $$$$$$$$ |  $$ |  $$ |     \\$$$$$$\\  $$ |      $$$$$$$$ |$$ |      $$$$$\\    $$$$$$$  |\n   $$ |   $$  __$$ |  $$ |  $$ |      \\____$$\\ $$ |      $$  __$$ |$$ |      $$  __|   $$  __$$<\n   $$ |   $$ |  $$ |  $$ |  $$ |     $$\\   $$ |$$ |  $$\\ $$ |  $$ |$$ |      $$ |      $$ |  $$ |\n   $$ |   $$ |  $$ |$$$$$$\\ $$$$$$$$\\\\$$$$$$  |\\$$$$$$  |$$ |  $$ |$$$$$$$$\\ $$$$$$$$\\ $$ |  $$ |\n   \\__|   \\__|  \\__|\\______|\\________|\\______/  \\______/ \\__|  \\__|\\________|\\________|\\__|  \\__|\nWelcome to Tailscaler please select an option: \n1) Get nodes\n2) Lookup node\n3) EXIT\n"
	if string(out) != expected {
		t.Fatalf("Expected %q\n got %q", expected, out)
	}
}
