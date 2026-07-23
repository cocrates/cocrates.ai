# TODO: jsondb

> **Workspace:** `jsondb/`
> **Updated:** 2026-07-21

## Snapshot

| Done | In progress | Pending | Blocked | Skipped |
|------|-------------|---------|---------|---------|
| 15   | 0           | 0       | 0       | 0       |

**Current focus:** —
**Recommended next:** —

## Active

(현재 진행 중인 작업 없음)

## Backlog

### Phase 1: 아키텍처 설계

- [x] **T-001** `done` — 언어/플랫폼 선택 (ASR-001)
  - Phase: 아키텍처 설계
  - Artifact: `jsondb/adr/language-platform.md`
  - Completed: 2026-07-21
  - Notes: Go 언어 선택

- [x] **T-002** `done` — 인터페이스形式 결정 (ASR-003)
  - Phase: 아키텍처 설계
  - Artifact: `jsondb/adr/interface-type.md`
  - Completed: 2026-07-21
  - Notes: 하이브리드 방식 (CLI + Server)

- [x] **T-003** `done` — 엔진 아키텍처 설계 (ASR-007)
  - Phase: 아키텍처 설계
  - Artifact: `jsondb/adr/engine-architecture.md`
  - Completed: 2026-07-21
  - Notes: 모듈 기반 설계 (Engine + Query + Storage)

- [x] **T-004** `done` — 저장 방식 결정 (ASR-002)
  - Phase: 아키텍처 설계
  - Artifact: `jsondb/adr/data-model.md`, `jsondb/adr/file-io-method.md`
  - Completed: 2026-07-21
  - Notes: 파일 기반 접근 + Storage Backend 추상화

- [x] **T-005** `done` — 동시성 제어 방식 결정 (ASR-006)
  - Phase: 아키텍처 설계
  - Artifact: `jsondb/adr/concurrency-control.md`
  - Completed: 2026-07-21
  - Notes: 파일 잠금 + 메모리 잠금 혼합 (프로세스 간 + 스레드 간)

- [x] **T-006** `done` — HTTP API 설계 (ASR-008)
  - Phase: 아키텍처 설계
  - Artifact: `jsondb/adr/http-api.md`, `jsondb/adr/json-patch.md`
  - Completed: 2026-07-21
  - Notes: RESTful API + JSON Patch/Merge Patch 지원 (json-patch 라이브러리)

- [x] **T-007** `done` — CLI 설계 (ASR-009)
  - Phase: 아키텍처 설계
  - Artifact: `jsondb/adr/cli-design.md`
  - Completed: 2026-07-21
  - Notes: 파일 관리 + 데이터 조회/변경 + 서버 관리 명령어

- [x] **T-011** `done` — JSON 작업 방식 결정 (ASR-004)
  - Phase: 아키텍처 설계
  - Artifact: `jsondb/adr/json-operations.md`
  - Completed: 2026-07-21
  - Notes: 기본 경로 + JSONPath 혼합 지원 (ojg 라이브러리 사용)

### Phase 1.5: 스펙 작성

- [x] **T-012** `done` — PRD.md 작성
  - Phase: 스펙 작성
  - Artifact: `jsondb/spec/PRD.md`
  - Completed: 2026-07-21
  - Notes: 제품 요구사항 정의서 작성 완료

- [x] **T-013** `done` — ASR-005 에이전트 통합 해결
  - Phase: 스펙 작성
  - Artifact: `jsondb/spec/ASR.md`
  - Completed: 2026-07-21
  - Notes: 에이전트는 HTTP API 직접 접근, CLI는 사람용

- [x] **T-014** `done` — 엔진 모듈 스펙 작성
  - Phase: 스펙 작성
  - Artifact: `jsondb/spec/engine.md`
  - Completed: 2026-07-21
  - Notes: 엔진 모듈 상세 스펙

- [x] **T-015** `done` — 서버 모듈 스펙 작성
  - Phase: 스펙 작성
  - Artifact: `jsondb/spec/server.md`
  - Completed: 2026-07-21
  - Notes: 서버 모듈 상세 스펙 (에이전트 HTTP API 접근 명시)

- [x] **T-016** `done` — CLI 모듈 스펙 작성
  - Phase: 스펙 작성
  - Artifact: `jsondb/spec/cli.md`
  - Completed: 2026-07-21
  - Notes: CLI 모듈 상세 스펙 (에이전트는 HTTP API 사용 명시)

### Phase 2: 구현

- [x] **T-008** `done` — 엔진 모듈 구현
  - Phase: 구현
  - Artifact: `jsondb/src/engine/`
  - Completed: 2026-07-21
  - Notes: 핵심 데이터베이스 로직, 테스트 통과

- [x] **T-009** `done` — 서버 모듈 구현
  - Phase: 구현
  - Artifact: `jsondb/src/server/`
  - Completed: 2026-07-21
  - Notes: HTTP 인터페이스, 기본 핸들러 구현

- [x] **T-010** `done` — CLI 구현
  - Phase: 구현
  - Artifact: `jsondb/src/cmd/`
  - Completed: 2026-07-21
  - Notes: CLI 명령어 구현, get/set/delete/create 동작 확인

## Completed

(이전 단계 완료 항목은 Phase 1, 1.5 섹션 참조)

## Notes

- **아키텍처 구조**: CLI → HTTP 요청 → Server (HTTP 인터페이스 + 엔진 모듈)
- **언어**: Go (성능 및 배포 용이성)
- **동시성**: 파일 잠금 + 메모리 잠금 혼합 (프로세스 간 + 스레드 간)
- **저장 방식**: 파일 기반 접근 + Storage Backend 추상화 (직접 I/O → mmap/WAL 변경 용이)
- **JSON 작업**: 기본 경로 + JSONPath 혼합 지원 (ojg 라이브러리 사용)
- **엔진 구조**: 모듈 기반 설계 (Engine + Query Parser/Executor + Storage Backend)
- **HTTP API**: RESTful API + JSON Patch/Merge Patch 지원 (json-patch 라이브러리)
- **CLI**: 파일 관리 + 데이터 조회/변경 + 서버 관리 명령어
- **에이전트 통합**: 에이전트는 HTTP API를 통해 서버에 직접 접근 (CLI 사용하지 않음)
- **구현 위치**: `jsondb/src/` (Go 모듈: github.com/drajin/jsondb)
- **테스트**: 엔진 모듈 테스트 통과, CLI 기본 기능 동작 확인
