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
	// Prepare the SQL INSERT statement
	sqlStatement := `INSERT INTO memos (title, content, priority) VALUES (?, ?, ?)`

	// Execute the SQL statement
	_, err := db.DB.Exec(sqlStatement, req.Title, req.Content, req.Priority)
	if err != nil {
		// If there is an error executing the SQL statement, return nil and the error
		return nil, err
	}

	// Create a MemoResponse object with the details of the new memo
	response := &pb.MemoResponse{
		Title:    req.Title,
		Content:  req.Content,
		Priority: req.Priority,
	}

	// Return the MemoResponse object
	return response, nil
}

func (s *MemoServiceServer) DeleteMemo(ctx context.Context, req *pb.DeleteMemoRequest) (*pb.DeleteMemoResponse, error) {
	// 데이터베이스에서 메모를 삭제하는 로직
	_, err := db.DB.Exec("DELETE FROM memos WHERE id = ?", req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteMemoResponse{Success: true}, nil
}

func (s *MemoServiceServer) GetMemo(ctx context.Context, req *pb.GetMemoRequest) (*pb.MemoResponse, error) {
	// Fetch the memo from the database using the provided ID
	row := db.DB.QueryRow("SELECT id, title, content, priority FROM memos WHERE id = ?", req.Id)

	var response pb.MemoResponse
	err := row.Scan(&response.Id, &response.Title, &response.Content, &response.Priority)
	if err != nil {
		return nil, err
	}

	// Return the fetched memo
	return &response, nil
}

func (s *MemoServiceServer) ListMemos(ctx context.Context, req *pb.ListMemosRequest) (*pb.ListMemosResponse, error) {
	// Implement logic to fetch a list of memos based on the request parameters
	// This is a simplified example. You would typically add filtering based on the request.
	rows, err := db.DB.Query("SELECT id, title, content, priority FROM memos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memos []*pb.MemoResponse
	for rows.Next() {
		var memo pb.MemoResponse
		if err := rows.Scan(&memo.Id, &memo.Title, &memo.Content, &memo.Priority); err != nil {
			return nil, err
		}
		memos = append(memos, &memo)
	}

	return &pb.ListMemosResponse{Memos: memos}, nil
}

func (s *MemoServiceServer) UpdateMemo(ctx context.Context, req *pb.UpdateMemoRequest) (*pb.MemoResponse, error) {
	// Update the memo in the database
	_, err := db.DB.Exec("UPDATE memos SET title = ?, content = ?, priority = ? WHERE id = ?", req.Title, req.Content, req.Priority, req.Id)
	if err != nil {
		return nil, err
	}

	// Return the updated memo
	return &pb.MemoResponse{
		Id:       req.Id,
		Title:    req.Title,
		Content:  req.Content,
		Priority: req.Priority,
	}, nil
}
