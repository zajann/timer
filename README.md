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
BenchmarkTimer-8        	511586547	       2.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkOriginTime-8   	 8739404	       134 ns/op	       0 B/op	       0 allocs/op
```
