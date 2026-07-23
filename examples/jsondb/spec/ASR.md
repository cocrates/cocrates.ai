# Architecturally Significant Requirements

Living registry for jsondb. Status of each ASR must stay current.

## Summary

| ID | Title | Category | Status | Related ADRs | Spec |
|----|-------|----------|--------|--------------|------|
| ASR-001 | 언어/플랫폼 선택 | Constraints | approved | adr/language-platform.md | spec/engine.md |
| ASR-002 | 데이터 구조 | Structure & organization | approved | adr/data-model.md, adr/file-io-method.md | spec/engine.md |
| ASR-003 | 인터페이스形式 | Deliverable form | approved | adr/interface-type.md | spec/server.md, spec/cli.md |
| ASR-004 | 키/필드 레벨 접근 | Quality bar | approved | adr/json-operations.md | spec/engine.md |
| ASR-005 | 에이전트 통합 | Integration & dependencies | designed | — | spec/server.md |
| ASR-006 | 동시성 제어 | Quality bar | designed | adr/concurrency-control.md, adr/server-engine-lifecycle.md, adr/concurrent-write-serialization.md | spec/engine.md |
| ASR-007 | 엔진 아키텍처 | Structure & organization | designed | adr/engine-architecture.md, adr/server-engine-lifecycle.md, adr/server-shutdown-cleanup.md | spec/engine.md |
| ASR-008 | HTTP API 설계 | Deliverable form | approved | adr/http-api.md, adr/json-patch.md | spec/server.md |
| ASR-009 | CLI 설계 | Deliverable form | approved | adr/cli-design.md | spec/cli.md |

## Dependency Order (recommended review path)

1. ASR-001 (언어/플랫폼) → ASR-003 (인터페이스) → ASR-007 (엔진) → ASR-006 (동시성) → ASR-002 (데이터 구조) → ASR-004 (접근 방식) → ASR-005 (통합) → ASR-008 (HTTP API) → ASR-009 (CLI)

## ASR Detail

### ASR-001 — 언어/플랫폼 선택

- **Category:** Constraints
- **Status:** approved
- **Statement:** jsondb를 구현할 프로그래밍 언어와 플랫폼을 결정해야 한다.
- **Why it matters:** 언어 선택은 에이전트와의 통합 용이성, 개발 생산성, 성능, 배포 방식에 직접적인 영향을 미친다.
- **Depends on:** 없음
- **Related ADRs:**
  - `jsondb/adr/language-platform.md` — approved — 언어/플랫폼 비교 분석
- **Resolution path:** adr
- **Resolution:** Go (Golang) 선택 — 성능 및 배포 용이성 근거
- **Spec:** `jsondb/spec/engine.md`
- **Notes:** 에이전트와의 협업이 주요 사용 사례이므로 통합 용이성이 중요했으나, 성능과 배포 용이성이 더 중요하다고 판단됨

### ASR-002 — 데이터 구조

- **Category:** Structure & organization
- **Status:** approved
- **Statement:** 복합 JSON 구조를 어떤 형태로 저장하고 관리할지 결정해야 한다.
- **Why it matters:** 데이터 구조는 쿼리 효율성, 확장성, 직렬화 방식에 영향을 미친다.
- **Depends on:** ASR-001
- **Related ADRs:**
  - `jsondb/adr/data-model.md` — approved — 데이터 모델 비교 분석
  - `jsondb/adr/file-io-method.md` — approved — 파일 I/O 방식 비교 분석
- **Resolution path:** adr
- **Resolution:** 파일 기반 접근 + Storage Backend 추상화 — 직접 I/O로 시작, 향후 mmap/WAL 변경 용이하도록 인터페이스 추상화
- **Spec:** —
- **Notes:** 중첩된 객체와 배열을 지원, 디렉토리 구조로 그룹화

### ASR-003 — 인터페이스形式

- **Category:** Deliverable form
- **Status:** approved
- **Statement:** jsondb를 어떤 형태로 제공할지 결정해야 한다 (라이브러리, CLI, 서버).
- **Why it matters:** 인터페이스 형태는 사용 방식과 에이전트와의 통합 방식을 결정한다.
- **Depends on:** ASR-001
- **Related ADRs:**
  - `jsondb/adr/interface-type.md` — approved — CLI vs 서버 비교 분석
- **Resolution path:** adr
- **Resolution:** 하이브리드 방식 (CLI + Server) — CLI는 Server에 HTTP 요청, Server는 엔진 모듈과 HTTP 인터페이스 서버 모듈로 구분
- **Spec:** `jsondb/spec/server.md`, `jsondb/spec/cli.md`
- **Notes:** CLI로 빠른 테스트, Server로 프로덕션 사용 가능하도록 설계. 에이전트는 HTTP API 직접 접근.

### ASR-004 — 키/필드 레벨 접근

- **Category:** Quality bar
- **Status:** approved
- **Statement:** 단일 키/필드를 효율적으로 읽고 수정할 수 있어야 한다.
- **Why it matters:** 현재 파일 I/O의 주요 비효율성을 해결하는 핵심 기능이다.
- **Depends on:** ASR-002
- **Related ADRs:**
  - `jsondb/adr/json-operations.md` — approved — JSON 작업 방식 비교 분석
- **Resolution path:** adr
- **Resolution:** 기본 경로 + JSONPath 혼합 지원 — 단순 경로(기본)와 JSONPath(고급)를 모두 지원하며, 문법에 따라 자동 감지
- **Spec:** `jsondb/spec/engine.md`
- **Notes:** Go 언어 JSONPath 라이브러리 (`ojg`) 사용, 에이전트와의 호환성 확보

### ASR-005 — 에이전트 통합

- **Category:** Integration & dependencies
- **Status:** designed
- **Statement:** 에이전트가 jsondb와 쉽게 통합할 수 있어야 한다.
- **Why it matters:** 프로젝트의 주요 목적이 에이전트와의 협업 효율성 향상이기 때문이다.
- **Depends on:** ASR-001, ASR-003
- **Related ADRs:** —
- **Resolution path:** direct-input
- **Resolution:** 에이전트는 HTTP API를 통해 서버에 직접 접근. CLI는 사람용 인터페이스.
- **Spec:** `jsondb/spec/server.md`
- **Notes:** 에이전트가 사용하는 언어/프레임워크와의 호환성 중요. HTTP API가 에이전트 통합의 핵심 인터페이스.

### ASR-006 — 동시성 제어

- **Category:** Quality bar
- **Status:** designed
- **Statement:** 여러 에이전트가 동시에 jsondb에 접근할 때 데이터 무결성과 일관성을 보장해야 한다.
- **Why it matters:** 동시 쓰기/읽기 시 경합 상태가 발생하면 데이터 손상이나 불일치가 발생할 수 있다. 이는 시스템의 신뢰성을 직접적으로 훼손한다.
- **Depends on:** ASR-001, ASR-002
- **Related ADRs:**
  - `jsondb/adr/concurrency-control.md` — approved — 동시성 제어 방식 비교 분석
  - `jsondb/adr/server-engine-lifecycle.md` — approved — 서버 엔진 생명주기 관리 (요청 레벨 Race Condition 해결)
  - `jsondb/adr/concurrent-write-serialization.md` — approved — 서버 동시 쓰기 직렬화 (파일 수준 잠금)
- **Resolution path:** adr
- **Resolution:** 파일 잠금 + 메모리 잠금 혼합 — 프로세스 간 동기화는 파일 잠금, 서버 내부 스레드 간 동기화는 RWLock. 요청 레벨에서는 파일별 독립 Engine 인스턴스로 Race Condition 해결. 동시 쓰기 직렬화는 파일 경로 기반 RWLock으로 해결
- **Spec:** `jsondb/spec/engine.md`
- **Notes:** CLI 모드는 파일 잠금만, Server 모드는 파일 잠금 + RWLock 혼합. 요청 레벨 Race Condition은 server-engine-lifecycle ADR에서 파일별 Long-lived Engine으로 해결. 동시 쓰기 직렬화는 concurrent-write-serialization ADR에서 파일 경로 기반 RWLock으로 결정

### ASR-007 — 엔진 아키텍처

- **Category:** Structure & organization
- **Status:** designed
- **Statement:** jsondb의 핵심 엔진 모듈이 어떤 구조로 설계되어야 하는지 결정해야 한다.
- **Why it matters:** 엔진은 데이터 저장, 조회, 수정의 핵심 로직을 담당하며, 시스템의 성능과 확장성을 결정한다.
- **Depends on:** ASR-001, ASR-003
- **Related ADRs:**
  - `jsondb/adr/engine-architecture.md` — approved — 엔진 아키텍처 설계
  - `jsondb/adr/server-engine-lifecycle.md` — approved — 서버 엔진 생명주기 관리 (Engine 수명 주기 설계)
  - `jsondb/adr/server-shutdown-cleanup.md` — approved — 서버 종료 시 정리 (Engine 정상 종료 보장)
- **Resolution path:** adr
- **Resolution:** 모듈 기반 설계 — Engine 인터페이스 + Query Parser/Executor + Storage Backend 구조. 서버 모드에서는 파일별 Long-lived Engine으로 생명주기 관리. 종료 시 graceful shutdown으로 정리
- **Spec:** `jsondb/spec/engine.md`
- **Notes:** CLI/Server에서 공통으로 사용하는 핵심 모듈. 서버 모드에서의 Engine 생명주기는 server-engine-lifecycle ADR에서 파일별 Long-lived Engine + Write-through로 설계 완료. 서버 종료 시 정리는 server-shutdown-cleanup ADR에서 signal 핸들러 + Graceful Shutdown으로 결정

### ASR-008 — HTTP API 설계

- **Category:** Deliverable form
- **Status:** approved
- **Statement:** Server 모듈이 제공할 HTTP API의 엔드포인트와 요청/응답 형식을 결정해야 한다.
- **Why it matters:** HTTP API는 CLI와 에이전트가 jsondb에 접근하는 주요 인터페이스이다.
- **Depends on:** ASR-003, ASR-007
- **Related ADRs:**
  - `jsondb/adr/http-api.md` — approved — HTTP API 설계
  - `jsondb/adr/json-patch.md` — approved — JSON Patch/Merge Patch 지원
- **Resolution path:** adr
- **Resolution:** RESTful API 설계 — 파일 관리 + 데이터 조회/변경 엔드포인트, JSON Patch/Merge Patch 지원
- **Spec:** `jsondb/spec/server.md`
- **Notes:** `github.com/evanphx/json-patch` 라이브러리 사용. 에이전트가 HTTP API를 직접 접근.

### ASR-009 — CLI 설계

- **Category:** Deliverable form
- **Status:** approved
- **Statement:** CLI 도구의 명령어 구조와 사용법을 결정해야 한다.
- **Why it matters:** CLI는 사람이 jsondb를 직접 조작하는 인터페이스이다.
- **Depends on:** ASR-003, ASR-008
- **Related ADRs:**
  - `jsondb/adr/cli-design.md` — approved — CLI 명령어 체계 설계
- **Resolution path:** adr
- **Resolution:** CLI 명령어 체계 설계 — 파일 관리 + 데이터 조회/변경 + 서버 관리 명령어
- **Spec:** `jsondb/spec/cli.md`
- **Notes:** get, set, delete, patch, ls, serve 등 주요 명령어. CLI는 사람용, 에이전트는 HTTP API 사용.