package pack

import (
  "testing"
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
        t.Fail()
      }
}
