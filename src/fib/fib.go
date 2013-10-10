package fib

func Fibonacci(n int) int {

  if n == 0 || n == 1 {
    return 1
  }
  return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciGenerator() func() int {
  var f0, f1 int = 1, 1
  return func() int {
    fib := f1 + f0
    f0 = f1
    f1 = fib
    return fib
  }
}

