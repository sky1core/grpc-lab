package db

import (
	"grpc-lab/legacy"
	pb "grpc-lab/proto/gen/go"
)

type Memo struct {
	ID        int64      `db:"id"`
	Title     string     `db:"title"`
	Content   string     `db:"content"`
	Priority  int32      `db:"priority"`
	CreatedAt UnixMillis `db:"created_at"`
}

func (m *Memo) ToProto() *pb.MemoResponse {
	return &pb.MemoResponse{
		Id:        m.ID,
		Title:     m.Title,
		Content:   m.Content,
		Priority:  m.Priority,
		CreatedAt: int64(m.CreatedAt),
	}
}

func (m *Memo) ToProtoV0() *legacy.MemoResponseV0 {
	return &legacy.MemoResponseV0{
		Id:        m.ID,
		Title:     m.Title,
		Content:   m.Content,
		Priority:  m.Priority,
		CreatedAt: int64(m.CreatedAt),
	}
}
