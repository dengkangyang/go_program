package ipc

import(
	"testing"
)

type EchoServer struct{

}

func (server *EchoServer) Handle(method, params string) *Response{
	return &Response{method, params}
}

func (server *EchoServer) Name() string{
	return "EchoServer"
}

func TestIpc(t *testing.T){
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClinet(server)
	client2 := NewIpcClinet(server)

	resp1, err1 := client1.Call("ECHO", "From Client1")
	resp2, err2 := client2.Call("ECHO", "From Client2")
	if err1 != nil {
		t.Error(err1)
	}
	if err2 != nil {
		t.Error(err2)
	}


	if resp1.Body != "From Client1" || resp2.Body != "From Client2" {
		t.Error("IpcClient.call failed.", resp1.Body, "resp2:", resp2.Body)
	}
	

	client1.Close()
	client2.Close() 
}