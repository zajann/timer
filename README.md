# timer

`timer` is a simple Go package to replace `time` package.

When you want to deal with time, calling `time.Now()` every time can be a little inefficient.

Let's suppose you're trying to write logs with timestamp. If 10k logs that should be written per sec are, you'll make 10k new `time.Time` object per sec. (In simple way)

So, I propose a way for derivering value from periodially updated internal `time.Time`.

## Usage

```go
func main() {
  if err := timer.Init(0); err != nil {		// required
    // error
  }
  year := timer.GetYear()
  month := timer.GetMonth()
  monthStr := timer.GetMonthString()
  day := timer.GetDay()
  hour := timer.GetHour()
  min := timer.GetMinute()
  sec := timer.GetSecond()
  nano := timer.GetNanoSecnd()
  unix := timer.GetUnix()
  unixnano := timer.GetUnixNano()
  timeStamp1 := timer.GetTimeFormat(time.RFC1123)
  timeStamp2 := timer.GetTimeFormat("2006-01-02 15:04:05")
  t := timer.GetTime()
}
```

You can set update cycle by `Init(d time.Duration)` , `0` means default: 100ms.

Once you've init timer by `Init()` , you can use globally. It has no worries about thread-safety because it has only read action.

## Benchmark

You can becnchmark with `timer_test.go` in package.

```bash
$ go test -bench=. -benchmem
```

```
goos: linux
goarch: amd64
pkg: github.com/zajann/timer
BenchmarkGetFormatStringWithTimer-16             2350820               463 ns/op              32 B/op          1 allocs/op
BenchmarkGetFormatStringWithoutTimer-16          1804837               643 ns/op              32 B/op          1 allocs/op
BenchmarkGetMonthStringWithTimer-16             395061864                3.08 ns/op            0 B/op          0 allocs/op
BenchmarkGetMonthStringWithoutTimer-16          12656670                88.1 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/zajann/timer 6.956s
```
