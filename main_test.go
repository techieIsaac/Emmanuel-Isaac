package main


import "testing"

func TestgetData(t *testing.T){
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  resp := ExecuteDataRequest(ctx,cancel,600)
  if resp !="data received"{
    t.Error("receive data test failed")
  }
}

func TestsetInitData(t *testing.T){
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  resp := ExecuteDataRequest(ctx,cancel,500)
  if resp !="timeout"{
    t.Error("receive data test failed")
  }
}



