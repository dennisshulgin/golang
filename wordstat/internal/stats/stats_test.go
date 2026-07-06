package stats

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetWords(t *testing.T) {
	s := "Denis denis! Hello my friend! I live in New-York... Ready to  work. "
	actualWords := GetWords(s)
	expectedWords := []string{"denis", "denis", "hello", "my", "friend", "i", "live", "in", "new", "york", "ready", "to", "work"}
	fmt.Println(actualWords)
	if !reflect.DeepEqual(actualWords, expectedWords) {
		t.Error("Slices are not equal")
	}
}
