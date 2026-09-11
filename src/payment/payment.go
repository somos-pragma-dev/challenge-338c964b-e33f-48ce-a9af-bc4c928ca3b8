package payment

import (
	"context"
	"log"
	"time"

	pb "fintech/payment/api"
	auditpb "fintech/audit/api"
	"google.golang.org/grpc"
)

type server struct {}

func (s *server) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Println("Processing payment", req)
	// Validación de la solicitud de pago
	if req.Amount <= 0 {
		return nil, &pb.PaymentError{Message: "Invalid amount"}
	}
	// Simular procesamiento de pago
	time.Sleep(100 * time.Millisecond)
	// Comunicación con el servicio de auditoría
	conn, err := grpc.DialContext(ctx, "audit:50051", grpc.WithInsecure())
	if err!= nil {
		return nil, err
	}
	defer conn.Close()
	auditClient := auditpb.NewAuditServiceClient(conn)
	_, err = auditClient.LogPayment(ctx, &auditpb.LogPaymentRequest{PaymentId: req.Id})
	if err!= nil {
		return nil, err
	}
	return &pb.PaymentResponse{Status: "SUCCESS"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err!= nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &server{})
	log.Println("Server started at ":50051")
	if err := s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}