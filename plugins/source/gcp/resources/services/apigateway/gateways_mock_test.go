// Code generated by codegen; DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"testing"

	pb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
)

func createGateways(gsrv *grpc.Server) error {
	fakeServer := &fakeGatewaysServer{}
	pb.RegisterApiGatewayServiceServer(gsrv, fakeServer)
	return nil
}

type fakeGatewaysServer struct {
	pb.UnimplementedApiGatewayServiceServer
}

func (f *fakeGatewaysServer) ListGateways(context.Context, *pb.ListGatewaysRequest) (*pb.ListGatewaysResponse, error) {
	resp := pb.ListGatewaysResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestGateways(t *testing.T) {
	client.MockTestGrpcHelper(t, Gateways(), createGateways, client.TestOptions{})
}