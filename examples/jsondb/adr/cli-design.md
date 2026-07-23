# CLI 설계

## Concern
CLI 도구의 명령어 구조와 사용법은 어떤 형태로 설계해야 하는가?

## Status
approved

## Context
- 하이브리드 인터페이스 (CLI + Server)
- CLI는 Server에 HTTP 요청을 전달
- 에이전트가 subprocess로 호출
- 단일 바이너리로 배포

## Decision
**CLI 명령어 체계 설계** 사용자 승인: 파일 관리 + 데이터 조회/변경 + 서버 관리 명령어 제공.

## CLI 개요

### 기본 정보
- **이름**: `jsondb`
- **배포**: 단일 바이너리
- **모드**: CLI 모드 (직접 실행) + Server 모드 (백그라운드)

### 실행 방식

```bash
# CLI 모드 (직접 실행)
$ jsondb get overview.json title

# Server 모드 (백그라운드 실행)
$ jsondb serve --port 8080

# CLI가 Server에 요청
$ jsondb --server http://localhost:8080 get overview.json title
```

## 명령어 설계

### 1. 파일 관리 명령어

| 명령어 | 설명 | 예시 |
|--------|------|------|
| `ls` | 파일 목록 조회 | `jsondb ls` |
| `ls {dir}` | 디렉토리 목록 조회 | `jsondb ls projects/` |
| `create {file}` | 새 파일 생성 | `jsondb create overview.json` |
| `rm {file}` | 파일 삭제 | `jsondb rm overview.json` |
| `cp {src} {dst}` | 파일 복사 | `jsondb cp a.json b.json` |
| `mv {src} {dst}` | 파일 이동/이름 변경 | `jsondb mv old.json new.json` |

### 2. 데이터 조회 명령어

| 명령어 | 설명 | 예시 |
|--------|------|------|
| `get {file}` | 전체 데이터 조회 | `jsondb get overview.json` |
| `get {file} {field}` | 특정 필드 조회 | `jsondb get overview.json title` |
| `get {file} {jsonpath}` | JSONPath 조회 | `jsondb get overview.json "$.items[0].name"` |

### 3. 데이터 변경 명령어

| 명령어 | 설명 | 예시 |
|--------|------|------|
| `set {file} {field} {value}` | 필드 값 설정 | `jsondb set overview.json title "새 제목"` |
| `delete {file} {field}` | 필드 삭제 | `jsondb delete overview.json description` |
| `patch {file} {patch}` | 부분 업데이트 | `jsondb patch overview.json '{"title":"새 제목"}'` |
| `patch {file} --json-patch {file}` | JSON Patch 적용 | `jsondb patch overview.json --json-patch changes.json` |

### 4. 서버 관리 명령어

| 명령어 | 설명 | 예시 |
|--------|------|------|
| `serve` | 서버 시작 | `jsondb serve` |
| `serve --port {port}` | 포트 지정 서버 시작 | `jsondb serve --port 8080` |
| `status` | 서버 상태 확인 | `jsondb status` |
| `stop` | 서버 중지 | `jsondb stop` |

## 상세 설계

### 1. get 명령어

**기본 사용법**:
```bash
# 전체 데이터 조회
$ jsondb get overview.json
{
  "title": "프로젝트 개요",
  "content": "이 프로젝트는...",
  "tags": ["important", "draft"]
}

# 특정 필드 조회
$ jsondb get overview.json title
"프로젝트 개요"

# 중첩된 필드 조회
$ jsondb get overview.json metadata.author
"홍길동"

# JSONPath 조회
$ jsondb get overview.json "$.items[?(@.active==true)].name"
["프로젝트1", "프로젝트2"]
```

**옵션**:
| 옵션 | 설명 | 예시 |
|------|------|------|
| `--pretty` | 보기 좋게 출력 | `jsondb get overview.json --pretty` |
| `--raw` | 원시 값 출력 | `jsondb get overview.json title --raw` |
| `--output {file}` | 파일로 저장 | `jsondb get overview.json --output result.json` |

### 2. set 명령어

**기본 사용법**:
```bash
# 필드 값 설정
$ jsondb set overview.json title "새 제목"

# 중첩된 필드 설정
$ jsondb set overview.json metadata.author "김철수"

# 배열 요소 설정
$ jsondb set overview.json tags[0] "updated"
```

**옵션**:
| 옵션 | 설명 | 예시 |
|------|------|------|
| `--type {type}` | 타입 지정 | `jsondb set overview.json count --type number 42` |
| `--create` | 필드가 없으면 생성 | `jsondb set overview.json newField "value" --create` |

### 3. delete 명령어

**기본 사용법**:
```bash
# 필드 삭제
$ jsondb delete overview.json description

# 중첩된 필드 삭제
$ jsondb delete overview.json metadata.author

# 배열 요소 삭제
$ jsondb delete overview.json tags[0]
```

### 4. patch 명령어

**기본 사용법**:
```bash
# JSON Merge Patch 적용
$ jsondb patch overview.json '{"title": "새 제목", "description": null}'

# JSON Patch 적용
$ jsondb patch overview.json --json-patch changes.json

# stdin에서 패치 읽기
$ echo '{"title": "새 제목"}' | jsondb patch overview.json -
```

**옵션**:
| 옵션 | 설명 | 예시 |
|------|------|------|
| `--json-patch` | JSON Patch 형식으로 처리 | `jsondb patch overview.json --json-patch changes.json` |
| `--dry-run` | 변경 사항 미리보기 | `jsondb patch overview.json '{"title":"새 제목"}' --dry-run` |

### 5. serve 명령어

**기본 사용법**:
```bash
# 기본 포트로 서버 시작
$ jsondb serve
Server started on :8080

# 포트 지정
$ jsondb serve --port 9090
Server started on :9090

# 루트 디렉토리 지정
$ jsondb serve --root /path/to/data
Server started on :8080, root: /path/to/data

# 백그라운드 실행
$ jsondb serve --daemon
Server started in background (PID: 12345)
```

**옵션**:
| 옵션 | 설명 | 기본값 |
|------|------|--------|
| `--port {port}` | 서버 포트 | 8080 |
| `--root {path}` | 데이터 루트 디렉토리 | 현재 디렉토리 |
| `--daemon` | 백그라운드 실행 | false |
| `--log-level {level}` | 로그 레벨 | info |
| `--log-file {file}` | 로그 파일 경로 | 없음 |

## 입출력 형식

### 기본 출력

```bash
# JSON 데이터는 JSON 형식으로 출력
$ jsondb get overview.json
{
  "title": "프로젝트 개요"
}

# 필드 값은 원시 값으로 출력
$ jsondb get overview.json title
"프로젝트 개요"
```

### 에러 출력

```bash
# 에러 시 exit code 1
$ jsondb get nonexistent.json
Error: File not found: nonexistent.json
Exit code: 1

# 상세 에러 (verbose 모드)
$ jsondb get nonexistent.json --verbose
Error: File not found: nonexistent.json
  at: /path/to/file
  reason: 파일이 존재하지 않습니다
Exit code: 1
```

### 성공 메시지

```bash
# 변경 성공 시
$ jsondb set overview.json title "새 제목"
OK

# 상세 출력
$ jsondb set overview.json title "새 제목" --verbose
OK: overview.json.title updated
  old: "기존 제목"
  new: "새 제목"
```

## 서버 모드 동작

### CLI → Server 요청

```bash
# Server 모드에서 실행 시
$ jsondb --server http://localhost:8080 get overview.json title

# 내부 동작:
# 1. HTTP GET 요청 생성
# 2. Server로 전송
# 3. 응답 수신
# 4. 결과 출력
```

### 서버 자동 시작

```bash
# Server가 실행 중이지 않으면 자동 시작
$ jsondb get overview.json title
Starting server on :8080...
"프로젝트 개요"
```

## 환경 변수

| 변수 | 설명 | 예시 |
|------|------|------|
| `JSONDB_SERVER` | Server URL | `JSONDB_SERVER=http://localhost:8080` |
| `JSONDB_ROOT` | 데이터 루트 디렉토리 | `JSONDB_ROOT=/path/to/data` |
| `JSONDB_LOG_LEVEL` | 로그 레벨 | `JSONDB_LOG_LEVEL=debug` |

## 설정 파일

**위치**: `~/.jsondb/config.json` 또는 `./.jsondb/config.json`

```json
{
  "server": {
    "url": "http://localhost:8080",
    "timeout": "30s"
  },
  "default": {
    "pretty": true,
    "verbose": false
  }
}
```

## Downstream Concerns
- [ ] **Shell 자동 완성**: bash/zsh/fish 자동 완성 지원
- [ ] **버전 정보**: `--version` 옵션
- [ ] **도움말**: `--help` 옵션 상세 도움말
- [ ] **설정 관리**: 설정 파일 형식과 위치

## Tags
`cli`, `command-line`, `interface`

## Approved
- 2026-07-21: CLI 명령어 체계 설계, 사용자 승인