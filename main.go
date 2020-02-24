/*
A function `getData` makes an API request. 

If it takes more than 500ms, another function called `setInitData` should run at the 500th ms mark.

Else, `getData` should run only.

Write unit tests for, and implement this system.
*/
package main

import (
  "fmt"
  "time"
  "golang.org/x/net/context"
  "errors"
)

func getData(ctx context.Context, tm time.Duration) error {
  //this simulates the time taken to call API and get response
	time.Sleep(tm * time.Millisecond)
	return errors.New("failed to get data")
}


func setInitData()string{
 return "timeout"
}

func handler(ctx context.Context) string{
  var resp string
  select {
	case <-time.After(500 * time.Millisecond):
		resp = setInitData()
	case <-ctx.Done():
		resp = "data received"
	}
  return resp
}

func ExecuteDataRequest(ctx context.Context, cancel context.CancelFunc, tm time.Duration) string{

  //running on different goroutines
   go func() {
		err := getData(ctx,tm)
    //if getData() returns err
    //cancel the execution
		if err != nil {
			cancel()
		}
	}()
  return handler(ctx)
}
func main() {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
    
  fmt.Println(ExecuteDataRequest(ctx,cancel,700))

}




//test with time 500 Milliseconds for setInitdata() to be called
//test with less than the threshold e.g 499 Milliseconds to
// have the getData func executed