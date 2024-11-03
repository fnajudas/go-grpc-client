package usecase

import (
	"base-project/constructs"
	student "base-project/proto/students"
	"context"
	"fmt"
	"log/slog"
)

type studentInt struct {
	client student.StudentServiceClient
}

func NewStudentSvc(client student.StudentServiceClient) *studentInt {
	return &studentInt{
		client: client,
	}
}

func (s *studentInt) GetStudentByNik(ctx context.Context, nik string) (*constructs.StudentResponse, error) {
	req := &student.StudentReq{Nik: nik}
	resp, err := s.client.Student(ctx, req)
	if err != nil {
		slog.Any("error", err)
		return nil, fmt.Errorf("failed to get student by nik: %v", err)
	}

	return &constructs.StudentResponse{
		Name: resp.Name,
	}, nil
}
