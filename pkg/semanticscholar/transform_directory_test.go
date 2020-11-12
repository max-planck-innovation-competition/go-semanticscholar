package semanticscholar

import (
	"log"
	"testing"
)

// Tests the complete directory
func TestTransformDirectory(t *testing.T) {
	err := TransformDirectory("../../test", "../../test", true)
	if err != nil {
		log.Fatal(err)
	}
}
