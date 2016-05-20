package pack

func Add(numbers ... int) int  {

    var result int

    for _, i := range numbers {
        result += i
    }

    return result
}

func Subtract(initial int,numbers ... int) int {
  for _, i := range numbers {
    initial -= i
  }

  return initial
}
