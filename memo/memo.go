package memo

import (
	"context"
	"grpc-lab/db"
	pb "grpc-lab/proto/gen/go"
)

type MemoServiceServer struct {
	pb.UnimplementedMemoServiceServer
}

func (s *MemoServiceServer) CreateMemo(ctx context.Context, req *pb.CreateMemoRequest) (*pb.MemoResponse, error) {
	// 메모 생성
	result, err := db.DB.NamedExec(`INSERT INTO memos (title, content, priority) VALUES (:title, :content, :priority)`,
		map[string]interface{}{
			"title":    req.Title,
			"content":  req.Content,
			"priority": req.Priority,
		})
	if err != nil {
		return nil, handleError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, handleError(err)
	}

	// 생성된 메모 정보 조회
	var memo db.Memo
	err = db.DB.Get(&memo, "SELECT id, title, content, priority, created_at FROM memos WHERE id = ?", id)
	if err != nil {
		return nil, handleError(err)
	}

	return memo.ToProto(), nil
}

func (s *MemoServiceServer) DeleteMemo(ctx context.Context, req *pb.DeleteMemoRequest) (*pb.DeleteMemoResponse, error) {
	// 데이터베이스에서 메모를 삭제하는 로직
	_, err := db.DB.Exec("DELETE FROM memos WHERE id = ?", req.Id)
	if err != nil {
		return nil, handleError(err)
	}
	return &pb.DeleteMemoResponse{Success: true}, nil
}

func (s *MemoServiceServer) GetMemo(ctx context.Context, req *pb.GetMemoRequest) (*pb.MemoResponse, error) {
	var memo db.Memo
	err := db.DB.Get(&memo, "SELECT id, title, content, priority, created_at FROM memos WHERE id = ?", req.Id)
	if err != nil {
		return nil, handleError(err)
	}

	return memo.ToProto(), nil
}

func (s *MemoServiceServer) ListMemos(ctx context.Context, req *pb.ListMemosRequest) (*pb.ListMemosResponse, error) {
	var memos []db.Memo
	err := db.DB.Select(&memos, "SELECT id, title, content, priority, created_at FROM memos")
	if err != nil {
		return nil, handleError(err)
	}

	var pbMemos []*pb.MemoResponse
	for _, memo := range memos {
		pbMemos = append(pbMemos, memo.ToProto())
	}

	return &pb.ListMemosResponse{Memos: pbMemos}, nil
}

func (s *MemoServiceServer) UpdateMemo(ctx context.Context, req *pb.UpdateMemoRequest) (*pb.MemoResponse, error) {
	// 메모 업데이트
	_, err := db.DB.Exec("UPDATE memos SET title = ?, content = ?, priority = ? WHERE id = ?",
		req.Title, req.Content, req.Priority, req.Id)
	if err != nil {
		return nil, handleError(err)
	}

	var memo db.Memo
	err = db.DB.Get(&memo, "SELECT id, title, content, priority, created_at FROM memos WHERE id = ?", req.Id)
	if err != nil {
		return nil, handleError(err)
	}

	return memo.ToProto(), nil
}
