package main

import (
	"./proto/calls.pb.go"
	"fmt"
)

type CallService struct{
	calls map[string] *calls.Call
}

func (s *CallService) Dial (ctx context.Context, req *calls.DialReq) (*calls.Call,error){
	sid:=id.New()
	call := &calls.Call{Sid : sid, From : req.From, To : req.To, CreatedAt : time.Now().Unix()}
	s.calls[sid] = call
	// make call
	return call,nil
}

func (s *CallService) Get (ctx context.Context, req *calls.GetReq) (*calls.Call,error){
	sid:=id.New()
	call, ok := s.calls[req.Sid]
	if(!ok){
		return nil, fmt.Errorf("No call found with SID",req.sid)
	}
	return call,nil
}