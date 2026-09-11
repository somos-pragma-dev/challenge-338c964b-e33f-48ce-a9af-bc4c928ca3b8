package audit

import (
	"context"
	"log"

	pb "fintech/audit/api"
	"google.golang.org/grpc"
)

type server struct {}

func (s *server) LogPayment(ctx context.Context, req *pb.LogPaymentRequest) (*pb.LogPaymentResponse, error) {
	log.Println("Logging payment", req)
	// Simular registro de auditoría
	time.Sleep(100 * time.Millisecond)
	return &pb.LogPaymentResponse{Status: "LOGGED"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err!= nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuditServiceServer(s, &server{})
	log.Println("Server started at ":50052")
	if err := s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}