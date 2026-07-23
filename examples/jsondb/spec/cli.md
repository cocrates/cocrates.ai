# CLI Module Specification

## Requirement

jsondb의 CLI 모듈은 사람이 jsondb를 직접 조작하는 인터페이스를 제공한다. CLI는 Server에 HTTP 요청을 전달하며, 단일 바이너리로 배포된다.

## Context

- 사람이 jsondb를 직접 조작하는 인터페이스
- CLI는 Server에 HTTP 요청 전달
- 단일 바이너리로 배포
- 에이전트는 HTTP API를 직접 사용 (CLI 사용하지 않음)

## Decisions

- **이름**: `jsondb`
- **배포**: 단일 바이너리
- **모드**: CLI 모드 (직접 실행) + Server 모드 (백그라운드)
- **출력**: JSON 형식 (기본), 원시 값 (선택)

## Requirements

### 파일 관리 명령어

- `ls` — 파일 목록 조회
- `ls {dir}` — 디렉토리 목록 조회
- `create {file}` — 새 파일 생성
- `rm {file}` — 파일 삭제
- `cp {src} {dst}` — 파일 복사
- `mv {src} {dst}` — 파일 이동/이름 변경

### 데이터 조회 명령어

- `get {file}` — 전체 데이터 조회
- `get {file} {field}` — 특정 필드 조회
- `get {file} {jsonpath}` — JSONPath 조회

### 데이터 변경 명령어

- `set {file} {field} {value}` — 필드 값 설정
- `delete {file} {field}` — 필드 삭제
- `patch {file} {patch}` — 부분 업데이트 (JSON Merge Patch)
- `patch {file} --json-patch {file}` — JSON Patch 적용

### 서버 관리 명령어

- `serve` — 서버 시작
- `serve --port {port}` — 포트 지정 서버 시작
- `status` — 서버 상태 확인
- `stop` — 서버 중지

### 옵션

- `--pretty` — 보기 좋게 출력
- `--raw` — 원시 값 출력
- `--output {file}` — 파일로 저장
- `--type {type}` — 타입 지정 (set 명령어)
- `--create` — 필드가 없으면 생성 (set 명령어)
- `--json-patch` — JSON Patch 형식으로 처리 (patch 명령어)
- `--dry-run` — 변경 사항 미리보기 (patch 명령어)
- `--verbose` — 상세 출력
- `--debug` — 디버그 모드
- `--log-file {file}` — 로그 파일 경로

### 환경 변수

- `JSONDB_SERVER` — Server URL
- `JSONDB_ROOT` — 데이터 루트 디렉토리
- `JSONDB_LOG_LEVEL` — 로그 레벨

### 설정 파일

- 위치: `~/.jsondb/config.json` 또는 `./.jsondb/config.json`
- 설정 항목: 서버 URL, 타임아웃, 기본 옵션

## Constraints

- Go 언어로 구현
- 단일 바이너리로 배포
- CLI는 Server에 HTTP 요청 전달
- 에이전트는 HTTP API를 직접 사용

## Out of Scope

- Shell 자동 완성 (bash/zsh/fish)
- 플러그인 시스템
- 다국어 지원

## Open Questions

- 없음

## Related

- `jsondb/adr/cli-design.md`
- `jsondb/adr/interface-type.md`
- `jsondb/spec/PRD.md`

## Tags

`cli`, `command-line`, `interface`
