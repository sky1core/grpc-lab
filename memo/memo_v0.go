package memo

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"grpc-lab/db"
	"grpc-lab/legacy"
	pb "grpc-lab/proto/gen/go"
	"log"
	"net/http"
	"strconv"
)

var decoder = schema.NewDecoder()

type ListMemosRequestV0 struct {
	Ids           []int64 `schema:"ids,comma"`
	Keyword       string  `schema:"keyword"`
	Priority      *int32  `schema:"priority"`
	FromCreatedAt *int64  `schema:"from_created_at"`
	ToCreatedAt   *int64  `schema:"to_created_at"`
}

type CreateMemoRequestV0 struct {
	Title    string `schema:"title"`
	Content  string `schema:"content"`
	Priority int32  `schema:"priority"`
}

type UpdateMemoRequestV0 struct {
	Title    string `schema:"title"`
	Content  string `schema:"content"`
	Priority int32  `schema:"priority"`
}

// ApiHandler 커스텀 핸들러 함수 타입
type ApiHandler func(http.ResponseWriter, *http.Request) (interface{}, error)

// wrapHandler 래퍼 함수
func wrapHandler(fn ApiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := fn(w, r)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError) // 오류 상태 코드 설정
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(data)
	}
}

func SetupLegacyHandlers() *http.ServeMux {
	r := mux.NewRouter()

	r.HandleFunc("/v0/memo", wrapHandler(handleLegacyCreateMemo)).Methods("POST")
	r.HandleFunc("/v0/memo/{id}", wrapHandler(handleLegacyGetMemo)).Methods("GET")
	r.HandleFunc("/v0/memo/{id}", wrapHandler(handleLegacyUpdateMemo)).Methods("PATCH")
	r.HandleFunc("/v0/memo/{id}", wrapHandler(handleLegacyDeleteMemo)).Methods("DELETE")
	r.HandleFunc("/v0/memos", wrapHandler(handleLegacyListMemos)).Methods("GET")

	legacyHandler := http.NewServeMux()
	legacyHandler.Handle("/", r)

	return legacyHandler
}

func handleLegacyCreateMemo(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if err := r.ParseForm(); err != nil {
		return nil, handleError(err)
	}
	var req CreateMemoRequestV0
	if err := decoder.Decode(&req, r.PostForm); err != nil {
		return nil, handleError(err)
	}

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

	var memo db.Memo
	err = db.DB.Get(&memo, "SELECT id, title, content, priority, created_at FROM memos WHERE id = ?", id)
	if err != nil {
		return nil, handleError(err)
	}

	return memo.ToProtoV0(), nil
}

func handleLegacyUpdateMemo(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return nil, handleError(err)
	}

	if err := r.ParseForm(); err != nil {
		return nil, handleError(err)
	}
	var updateReq UpdateMemoRequestV0
	if err := decoder.Decode(&updateReq, r.PostForm); err != nil {
		return nil, handleError(err)
	}

	_, err = db.DB.Exec("UPDATE memos SET title = ?, content = ?, priority = ? WHERE id = ?",
		updateReq.Title, updateReq.Content, updateReq.Priority, id)
	if err != nil {
		return nil, handleError(err)
	}

	var memo db.Memo
	err = db.DB.Get(&memo, "SELECT id, title, content, priority, created_at FROM memos WHERE id = ?", id)
	if err != nil {
		return nil, handleError(err)
	}

	return memo.ToProtoV0(), nil
}

func handleLegacyGetMemo(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return nil, handleError(err)
	}

	var memo db.Memo
	err = db.DB.Get(&memo, "SELECT id, title, content, priority, created_at FROM memos WHERE id = ?", id)
	if err != nil {
		return nil, handleError(err)
	}

	return memo.ToProtoV0(), nil
}

func handleLegacyDeleteMemo(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return nil, handleError(err)
	}

	_, err = db.DB.Exec("DELETE FROM memos WHERE id = ?", id)
	if err != nil {
		return nil, handleError(err)
	}

	return &pb.DeleteMemoResponse{Success: true}, nil
}

// handleLegacyListMemos는 메모 목록을 조회하는 핸들러 함수입니다.
func handleLegacyListMemos(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var request ListMemosRequestV0
	if err := decoder.Decode(&request, r.URL.Query()); err != nil {
		log.Printf("Error decoding query params: %v", err)
		return nil, err
	}

	var memos []db.Memo
	err := db.DB.Select(&memos, "SELECT id, title, content, priority, created_at FROM memos") // 예시 쿼리
	if err != nil {
		return nil, err
	}

	var pbMemos []*legacy.MemoResponseV0
	for _, memo := range memos {
		pbMemos = append(pbMemos, memo.ToProtoV0())
	}

	return pbMemos, nil
}
