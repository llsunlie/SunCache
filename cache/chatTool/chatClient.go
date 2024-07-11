package chatTool

import (
	context "context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	serverSocket string
}

func NewClient(serverSocket string) (client *Client) {
	client = &Client{
		serverSocket: serverSocket,
	}
	return
}

func (c *Client) GetServerSocket() (serverSocket string) {
	return c.serverSocket
}

func (c *Client) Get(ctx context.Context, req *Request) (res *Response, err error) {
	var conn *grpc.ClientConn
	conn, err = grpc.NewClient(c.serverSocket, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()

	client := NewChatToolClient(conn)

	res, err = client.Get(ctx, req)
	return
}
