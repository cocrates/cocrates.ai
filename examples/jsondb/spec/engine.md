# Engine Module Specification

## Requirement

jsondb의 핵심 엔진 모듈은 데이터 저장, 조회, 수정의 핵심 로직을 담당한다. CLI/Server에서 공통으로 사용되며, 파일 기반 저장과 JSONPath 지원을 제공한다.

## Context

- jsondb 프로젝트의 핵심 모듈
- CLI와 Server에서 공통으로 사용
- 파일 기반 데이터 모델 + Storage Backend 추상화
- 기본 경로 + JSONPath 혼합 지원

## Decisions

- **언어**: Go (Golang)
- **저장 방식**: 파일 기반 + Storage Backend 추상화 (Phase 1: 직접 I/O)
- **쿼리 방식**: 기본 경로 (점 구분자) + JSONPath (RFC 9535) 혼합 지원
- **JSONPath 라이브러리**: `github.com/ohler55/ojg`
- **동시성**: 파일 잠금 (flock) + RWLock 혼합
- **에러 처리**: JsonDBError 타입으로 구조화된 에러 반환

## Requirements

### Engine 인터페이스

- `Open(path string) error` — JSON 파일 열기
- `Get(fieldPath string) (interface{}, error)` — 필드 조회
- `Set(fieldPath string, value interface{}) error` — 필드 값 설정
- `Delete(fieldPath string) error` — 필드 삭제
- `Save() error` — 파일 저장
- `Close() error` — 파일 닫기

### Query Parser

- 점(.) 구분자 기반 경로 파싱 지원
- JSONPath (`$` 접두사) 자동 감지 및 파싱
- 배열 인덱스 (`[0]`) 지원
- 중첩 객체 경로 지원

### Query Executor

- 기본 경로 기반 조회/설정/삭제 실행
- JSONPath 기반 조회/설정/삭제 실행
- `ojg` 라이브러리 사용한 JSONPath 처리

### Storage Backend

- `StorageBackend` 인터페이스 정의
- `DirectIOBackend` 구현체 (Phase 1)
- `Open`, `Read`, `Write`, `Close` 메서드
- 파일 잠금 (flock) 지원

### 데이터 모델

- 파일 시스템 유사 구조
- JSON 파일 단위 저장
- 디렉토리로 논리적 그룹화
- 중첩 객체와 배열 지원

### 동시성 제어

- 프로세스 간: 파일 잠금 (flock)
- 스레드 간: RWLock (Server 모드)
- CLI 모드: 파일 잠금만 사용

### 에러 처리

- `ErrFileNotFound` — 파일 없음
- `ErrInvalidPath` — 잘못된 경로
- `ErrFieldNotFound` — 필드 없음
- `ErrTypeMismatch` — 타입 불일치
- `ErrPermissionDenied` — 접근 거부
- `ErrStorageFailure` — 저장소 오류

## Constraints

- Go 언어로 구현
- `github.com/ohler55/ojg` 라이브러리 사용
- 파일 기반 저장 (DB 백엔드 사용 불가)
- Phase 1은 직접 I/O로 구현

## Out of Scope

- mmap/WAL 등 고성능 저장 방식 (Phase 2+)
- 인덱싱 시스템
- 관계형 데이터베이스 기능
- 네트워크 기반 분산 시스템

## Open Questions

- 없음

## Related

- `jsondb/adr/engine-architecture.md`
- `jsondb/adr/data-model.md`
- `jsondb/adr/file-io-method.md`
- `jsondb/adr/json-operations.md`
- `jsondb/adr/concurrency-control.md`
- `jsondb/spec/PRD.md`

## Tags

`engine`, `core`, `storage`, `query`, `jsonpath`
