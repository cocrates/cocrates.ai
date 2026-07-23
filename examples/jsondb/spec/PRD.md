# jsondb — Product Requirement Definition

## Product Overview

jsondb는 에이전트와의 협업에서 파일 I/O 비효율성을 해결하기 위한 경량 JSON 데이터베이스이다. 기존 Markdown/JSON 파일의 전체 읽기-파싱-수정-쓰기 과정 대신, 필드 레벨 접근으로 효율적인 데이터 조작을 제공한다.

## Target Audience

- AI 에이전트 (subprocess 또는 HTTP API를 통해 jsondb에 접근)
- 개발자 (CLI를 통해 직접 조작)

## Core Problem

현재 에이전트가 JSON 파일의 단일 필드를 변경하려고 해도 전체 파일을 읽고 파싱해야 하는 비효율이 존재한다. jsondb는 이 문제를 해결하여 에이전트의 데이터 조작 효율성을 향상시킨다.

## Architecture Summary

```
CLI (cmd/) ──HTTP──▶ Server (server/) ──▶ Engine (engine/) ──▶ Storage (storage/)
```

- **CLI**: Server에 HTTP 요청을 전달하는 클라이언트
- **Server**: HTTP 인터페이스 + Engine 모듈
- **Engine**: 핵심 데이터베이스 로직 (Query Parser + Executor)
- **Storage**: 파일 기반 저장 + Backend 추상화

## Technology Stack

| 항목 | 결정 |
|------|------|
| 언어 | Go (Golang) |
| JSONPath | `github.com/ohler55/ojg` |
| JSON Patch | `github.com/evanphx/json-patch` |
| 동시성 | 파일 잠금 (flock) + RWLock 혼합 |
| 저장 방식 | 파일 기반 + Storage Backend 추상화 |

## Key Features

### 1. 필드 레벨 접근

- 기본 경로 (점 구분자): `title`, `metadata.author`, `items[0].name`
- JSONPath (RFC 9535): `$.items[?(@.active==true)].name`
- 문법에 따라 자동 감지

### 2. 하이브리드 인터페이스

- **CLI 모드**: `jsondb get overview.json title`
- **Server 모드**: `jsondb serve --port 8080`
- CLI는 내부적으로 Server에 HTTP 요청 전달

### 3. RESTful HTTP API

- 파일 관리: `GET/POST/DELETE /api/v1/files`
- 데이터 조회: `GET /api/v1/data/{path}?field={field}`
- 데이터 변경: `PUT /api/v1/data/{path}?field={field}`
- 부분 업데이트: `PATCH /api/v1/data/{path}` (JSON Patch + Merge Patch)

### 4. 동시성 제어

- 프로세스 간: 파일 잠금 (flock)
- 스레드 간: RWLock (Server 모드)
- CLI 모드: 파일 잠금만 사용

## Data Model

- 파일 시스템 유사 구조
- JSON 파일 단위 저장
- 디렉토리로 논리적 그룹화
- 메타데이터 없음 (Git 친화적)

## Dependencies

| 라이브러리 | 용도 |
|-----------|------|
| `github.com/ohler55/ojg` | JSONPath 파싱 및 실행 |
| `github.com/evanphx/json-patch` | JSON Patch/Merge Patch 처리 |

## Out of Scope

- 관계형 데이터베이스 기능 (조인, 트랜잭션)
- 네트워크 기반 분산 시스템
- 그래픽 사용자 인터페이스
- 실시간 동기화 (WebSocket)
- 인덱싱 시스템 (Phase 1)

## Quality Goals

- **단순성**: 추가 설정 없이 즉시 사용 가능
- **성능**: 필드 레벨 접근으로 전체 파일 I/O 대비 향상
- **안정성**: 동시 접근 시 데이터 무결성 보장
- **유연성**: CLI + Server 모드 모두 지원
