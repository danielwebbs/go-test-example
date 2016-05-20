package pack

import (
  "bytes"
  "testing"
  "text/template"
)

func BenchmarkExample(b *testing.B)  {
  temp, _ := template.New("").Parse("Hello, Go")

  // We reset the timer to prevent go from including the setup code in the benchmark as this would skewer results
  b.ResetTimer()

  var buf bytes.Buffer
  //N is used to ensure the test takes the alloted time ( 1 second by default, provided by b)
  for i := 0; i < b.N; i++ {
    temp.Execute(&buf, nil)
    buf.Reset()
  }
}
