package array

import "testing"

func TestChunk(t *testing.T) {
	var testData = []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
	t.Run("Chunk test 1", func(t *testing.T) {
		var result, err = Chunk(testData, 1)
		if len(result) != 8 || err != nil {
			t.Errorf("Chunk() = %v, want %v", len(result), 8)
		}
	})
	t.Run("Chunk test 2", func(t *testing.T) {
		var result, err = Chunk(testData, 2)
		if len(result) != 4 || err != nil {
			t.Errorf("Chunk() = %v, want %v", len(result), 4)
		}
	})
	t.Run("Chunk test 3", func(t *testing.T) {
		var result, err = Chunk(testData, 3)
		if len(result) != 3 || err != nil {
			t.Errorf("Chunk() = %v, want %v", len(result), 3)
		}
	})
}
