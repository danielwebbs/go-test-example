package pack

import (
  "testing"
  "time"
)

//  AAA testing
func TestCanAddNumbers(t *testing.T) {
      // Arrange
      one := 1
      two := 2

      // Act
      result := Add(one, two)

      // Assert
      if result != 3 {
        t.Log("Could not add numbers")
        // FailNow will stop any further testing, this makes senes if there is no point in doing the other tests if a main one fails
        t.FailNow()
      }
}

func TestCanSubtractNumbers(t *testing.T) {
  // Arrange
  ten := 10
  five := 5

  // Act
  result := Subtract(ten, five)
    // Simple timeout
    time.Sleep(3 * time.Second)

  // Assert
    if result != 5 {
      // Error is the same as t.Log and t.Fail
      // Fatal is the same as t.log and t.FailNow
      t.Error("Failed to subtract 5 from 10")
    }
}
