# 엔진 아키텍처

## Concern
jsondb의 핵심 엔진 모듈이 어떤 구조로 설계되어야 하는가?

## Status
approved

## Context
- 파일 기반 데이터 모델 + Storage Backend 추상화 채택
- 기본 경로 + JSONPath 혼합 지원
- 하이브리드 인터페이스 (CLI + Server)
- Go 언어로 구현

## Decision
**엔진 아키텍처 (모듈 기반 설계)** 사용자 승인: Engine 인터페이스 + Query Parser/Executor + Storage Backend 구조로 설계.

## 아키텍처 개요

### 전체 구조

```
┌─────────────────────────────────────────────────────────────────┐
│                         jsondb                                   │
├─────────────────────────────────────────────────────────────────┤
│  CLI 모드                    │  Server 모드                      │
│  ┌─────────────────────┐     │  ┌─────────────────────────────┐ │
│  │   CLI (cmd/)        │     │  │  HTTP Server (server/)      │ │
│  │   - 명령어 파싱      │     │  │  - 라우팅                    │ │
│  │   - HTTP 요청       │ ──▶ │  │  - 요청/응답 처리           │ │
│  └─────────────────────┘     │  └──────────┬──────────────────┘ │
│                              │             │                     │
│                              │             ▼                     │
│                              │  ┌─────────────────────────────┐ │
│                              │  │     Engine (engine/)        │ │
│                              │  │  ┌─────────────────────┐    │ │
│                              │  │  │  Query Parser        │    │ │
│                              │  │  │  - 경로 파싱         │    │ │
│                              │  │  │  - JSONPath 파싱     │    │ │
│                              │  │  └─────────────────────┘    │ │
│                              │  │  ┌─────────────────────┐    │ │
│                              │  │  │  Query Executor      │    │ │
│                              │  │  │  - 조회 실행         │    │ │
│                              │  │  │  - 변경 실행         │    │ │
│                              │  │  └─────────────────────┘    │ │
│                              │  └──────────┬──────────────────┘ │
│                              │             │                     │
│                              │             ▼                     │
│                              │  ┌─────────────────────────────┐ │
│                              │  │  Storage Backend (storage/) │ │
│                              │  │  ┌─────────────────────┐    │ │
│                              │  │  │  StorageBackend      │    │ │
│                              │  │  │  (인터페이스)        │    │ │
│                              │  │  └─────────────────────┘    │ │
│                              │  │  ┌─────────────────────┐    │ │
│                              │  │  │  DirectIOBackend     │    │ │
│                              │  │  │  (Phase 1 구현체)    │    │ │
│                              │  │  └─────────────────────┘    │ │
│                              │  └─────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
```

### 모듈 구성

```
jsondb/
├── cmd/                    ← CLI 진입점
│   └── main.go
├── server/                 ← HTTP 서버 모듈
│   ├── server.go
│   └── handler.go
├── engine/                 ← 핵심 엔진 모듈
│   ├── engine.go           ← 엔진 인터페이스
│   ├── query/
│   │   ├── parser.go       ← 경로 파싱 (기본 + JSONPath)
│   │   └── executor.go     ← 쿼리 실행
│   └── types.go            ← 공통 타입 정의
├── storage/                ← Storage Backend 모듈
│   ├── backend.go          ← StorageBackend 인터페이스
│   ├── direct.go           ← DirectIO 구현체
│   └── cache.go            ← 캐싱 (선택적)
├── go.mod
└── go.sum
```

## 상세 설계

### 1. Engine 인터페이스

```go
// engine/engine.go

type Engine interface {
    // 파일 열기
    Open(path string) error
    
    // 조회
    Get(fieldPath string) (interface{}, error)
    
    // 변경
    Set(fieldPath string, value interface{}) error
    
    // 삭제
    Delete(fieldPath string) error
    
    // 파일 저장
    Save() error
    
    // 파일 닫기
    Close() error
}

// engine 구현체
type jsonEngine struct {
    storage  storage.StorageBackend
    query    *query.Executor
    filePath string
    data     map[string]interface{}
}
```

### 2. Query Parser

```go
// engine/query/parser.go

type QueryType int

const (
    QueryTypeBasic QueryType = iota  // 단순 경로 (점 구분자)
    QueryTypeJSONPath                // JSONPath 문법
)

type ParsedQuery struct {
    Type     QueryType
    Path     string           // 원본 경로
    Segments []PathSegment    // 파싱된 세그먼트
}

type PathSegment struct {
    Field    string           // 필드명
    Index    *int             // 배열 인덱스 (선택적)
    Filter   string           // JSONPath 필터 (선택적)
}

// 파싱 함수
func Parse(path string) (*ParsedQuery, error) {
    if strings.HasPrefix(path, "$.") {
        return parseJSONPath(path)
    }
    return parseBasicPath(path)
}
```

### 3. Query Executor

```go
// engine/query/executor.go

type Executor struct {
    jsonPathParser jp.Parser  // ojg 라이브러리 사용
}

func (e *Executor) Get(data interface{}, query *ParsedQuery) (interface{}, error) {
    switch query.Type {
    case QueryTypeBasic:
        return e.getBasic(data, query)
    case QueryTypeJSONPath:
        return e.getJSONPath(data, query)
    }
}

func (e *Executor) Set(data interface{}, query *ParsedQuery, value interface{}) error {
    switch query.Type {
    case QueryTypeBasic:
        return e.setBasic(data, query, value)
    case QueryTypeJSONPath:
        return e.setJSONPath(data, query, value)
    }
}
```

### 4. Storage Backend

```go
// storage/backend.go

type StorageBackend interface {
    Open(path string) error
    Read() (map[string]interface{}, error)
    Write(data map[string]interface{}) error
    Close() error
}

// storage/direct.go
type DirectIOBackend struct {
    file    *os.File
    data    map[string]interface{}
    path    string
    dirty   bool  // 변경 여부
}

func (b *DirectIOBackend) Read() (map[string]interface{}, error) {
    // 파일 읽기 → JSON 파싱 → map 반환
}

func (b *DirectIOBackend) Write(data map[string]interface{}) error {
    // map → JSON 직렬화 → 파일 쓰기
}
```

### 5. 데이터 흐름

**조회 흐름**:
```
CLI/Server
    ↓
Engine.Get(fieldPath)
    ↓
Query Parser (경로 타입 감지)
    ↓
Query Executor (조회 실행)
    ↓
Storage Backend (데이터 읽기)
    ↓
결과 반환
```

**변경 흐름**:
```
CLI/Server
    ↓
Engine.Set(fieldPath, value)
    ↓
Query Parser (경로 타입 감지)
    ↓
Storage Backend (데이터 읽기)
    ↓
Query Executor (변경 실행)
    ↓
Storage Backend (데이터 쓰기)
    ↓
완료
```

## 에러 처리

### 에러 타입

```go
type ErrorType int

const (
    ErrFileNotFound ErrorType = iota
    ErrInvalidPath
    ErrFieldNotFound
    ErrTypeMismatch
    ErrPermissionDenied
    ErrStorageFailure
)

type JsonDBError struct {
    Type    ErrorType
    Message string
    Path    string
    Cause   error
}
```

### 에러 처리 방식

| 상황 | 에러 타입 | 처리 |
|------|----------|------|
| 파일 없음 | ErrFileNotFound | 명확한 에러 메시지 |
| 잘못된 경로 | ErrInvalidPath | 사용자 친화적 안내 |
| 필드 없음 | ErrFieldNotFound | nil 반환 또는 에러 |
| 타입 불일치 | ErrTypeMismatch | 기대 타입 안내 |
| 접근 거부 | ErrPermissionDenied | 권한 안내 |
| 저장 실패 | ErrStorageFailure | 재시도 안내 |

## 로깅/디버깅

### 로그 레벨

| 레벨 | 용도 | 예시 |
|------|------|------|
| DEBUG | 개발 시 상세 정보 | 쿼리 실행 상세 |
| INFO | 일반적인 동작 | 파일 열기/닫기 |
| WARN | 경고 상황 | 캐시 미스, 재시도 |
| ERROR | 오류 상황 | 파일 읽기 실패 |

### 디버그 모드

```bash
# 디버그 모드로 실행
$ jsondb --debug get overview.json title

# 로그 파일 저장
$ jsondb --log-file=jsondb.log get overview.json title
```

## Downstream Concerns
- [ ] **동시성 제어**: 엔진 수준에서의 잠금 전략
- [ ] **캐싱 전략**: 메모리 캐시 정책
- [ ] **배치 처리**: 여러 작업을 하나로 묶기
- [ ] **성능 모니터링**: 쿼리 성능 측정

## Tags
`engine`, `architecture`, `module-design`

## Approved
- 2026-07-21: 엔진 아키텍처 (모듈 기반 설계), 사용자 승인