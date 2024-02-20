package db

import "strings"

// QueryBuilder - 쿼리와 인자를 관리하는 래퍼 구조체
type QueryBuilder struct {
	query strings.Builder
	args  []interface{}
}

// NewQueryBuilder - QueryBuilder의 새 인스턴스를 생성
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

// AppendCondition - 조건과 인자를 쿼리에 추가
func (qb *QueryBuilder) Append(condition string, args ...interface{}) {
	qb.query.WriteString(condition)
	if len(args) > 0 {
		qb.args = append(qb.args, args...)
	}
}

// Build - 최종 쿼리 문자열과 인자 슬라이스를 반환
func (qb *QueryBuilder) Build() (string, []interface{}) {
	return qb.query.String(), qb.args
}
