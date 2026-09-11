package payment

import (
	"context"
	"testing"

	pb "fintech/payment/api"
	"google.golang.org/grpc"
)

func TestProcessPayment(t *testing.T) {
	conn, err := grpc.DialContext(context.Background(), "localhost:50051", grpc.WithInsecure())
	if err!= nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewPaymentServiceClient(conn)
	req := &pb.PaymentRequest{Id: "1", Amount: 100}
	res, err := client.ProcessPayment(context.Background(), req)
	if err!= nil {
		t.Fatalf("ProcessPayment failed: %v", err)
	}
	if res.Status!= "SUCCESS" {
		t.Fatalf("expected SUCCESS, got %v", res.Status)
	}
}