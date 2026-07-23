# 데이터 모델 — 파일 기반 접근

## Concern
jsondb의 데이터 모델은 어떤 구조로 설계해야 하는가?

## Status
approved

## Context
- 기존 Markdown 파일에서 JSON 구조화된 DB로 전환
- 파일 시스템처럼 직관적인 접근 방식 선호
- 복합 JSON 구조 지원 필요 (중첩 객체, 배열)
- 필드 레벨 접근이 핵심 기능

## Decision
**파일 기반 접근 (Option A)** 사용자 승인: 파일 시스템과 유사한 직관적인 구조, 단순성 및 Git 호환성 근거.

## Options

### Option A — 파일 기반 접근 (파일 시스템 유사)

**구조**:
```
{jsondb-root}/
├── overview.json          ← 단일 JSON 문서
├── settings.json
├── logs/                  ← 디렉토리 (문서 그룹)
│   ├── 2026-07-21.json
│   └── 2026-07-20.json
└── projects/
    ├── project1.json
    └── project2.json
```

**접근 방식**:
```bash
# 파일 + JSON 경로로 접근
$ jsondb get overview.json title
"기존 제목"

$ jsondb set overview.json tags[0] "important"

# 전체 파일 읽기
$ jsondb read overview.json

# 디렉토리 내 파일 목록
$ jsondb list logs/
```

**장점**:
- **직관성**: 파일 시스템과 유사한 접근 방식
- **호환성**: 기존 파일 관리 도구와 호환 (ls, cp, mv)
- **Git 친화적**: JSON 파일을 그대로 커밋 가능
- **유연성**: 디렉토리 구조로 자유로운 조직화
- **단순성**: 추가 메타데이터 없음

**단점**:
- **메타데이터 부재**: 생성일, 수정일 등 자동 관리 안 됨
- **문서 간 관계**: 참조, 조인 등 관계형 기능 없음
- **인덱싱 없음**: 특정 값으로 검색 시 전체 스캔 필요

---

### Option B — 컬렉션 기반 접근

**구조**:
```
{jsondb-root}/
├── db.json                 ← 데이터베이스 메타데이터
└── collections/
    ├── overview/           ← 컬렉션
    │   ├── doc1.json       ← 문서 (ID 기반)
    │   └── doc2.json
    └── settings/
        └── config.json
```

**접근 방식**:
```bash
# 컬렉션 + 문서 ID로 접근
$ jsondb get overview doc1 title
"기존 제목"

# 문서 생성
$ jsondb create overview --id new-doc

# 컬렉션 전체 조회
$ jsondb list overview
```

**장점**:
- **메타데이터 지원**: 문서별 타임스탬프, 버전 관리 가능
- **구조화**: 컬렉션으로 논리적 그룹화
- **ID 관리**: 문서별 고유 ID로 안정적 접근

**단점**:
- **복잡성**: 추가 구조 (컬렉션, ID) 필요
- **비직관적**: 파일 시스템과 다른 접근 방식
- **제약**: ID 생성 및 관리 필요

---

### Option C — 하이브리드 접근

**구조**:
```
{jsondb-root}/
├── .jsondb/                ← jsondb 메타데이터
│   ├── index.json          ← 필드 인덱스
│   └── meta.json           ← 시스템 메타데이터
├── overview.json
└── projects/
    └── project1.json
```

**접근 방식**:
```bash
# 파일 기반 접근 (기본)
$ jsondb get overview.json title

# 인덱싱된 검색 (선택적)
$ jsondb search "title == '특정 제목'"
```

**장점**:
- **유연성**: 파일 기반 + 선택적 인덱싱
- **확장성**: 필요 시 메타데이터 추가 가능
- **성능**: 인덱스로 빠른 검색 가능

**단점**:
- **복잡성**: 인덱스 관리 필요
- **구현 부담**: 두 가지 방식 모두 구현

## Tradeoffs

| 기준 | 파일 기반 | 컬렉션 기반 | 하이브리드 |
|------|----------|------------|----------|
| **직관성** | ★★★★★ | ★★★☆☆ | ★★★★☆ |
| **단순성** | ★★★★★ | ★★★☆☆ | ★★★☆☆ |
| **Git 호환** | ★★★★★ | ★★★☆☆ | ★★★★☆ |
| **메타데이터** | ★★☆☆☆ | ★★★★★ | ★★★★☆ |
| **검색 성능** | ★★☆☆☆ | ★★★☆☆ | ★★★★★ |
| **구현 난이도** | ★★★★★ | ★★★☆☆ | ★★☆☆☆ |

## Recommendation
**파일 기반 접근 (Option A)** 추천 이유:

1. **직관성**: 기존 파일 시스템 경험과 일치
2. **단순성**: 추가 구조 없이 빠른 구현
3. **Git 호환**: JSON 파일을 그대로 버전 관리 가능
4. **에이전트 친화적**: 에이전트가 파일 경로로 자연스럽게 접근

필요 시 향후 Option C로 확장 가능.

## Downstream Concerns
- [ ] **인덱싱 방식**: 특정 필드로 검색할 때 어떤 방식을 사용할지
- [ ] **파일명 충돌 처리**: 동일한 파일명이 있을 때의 처리
- [ ] **대용량 파일 처리**: 매우 큰 JSON 파일 처리 방식
- [ ] **메타데이터 확장**: 필요 시 생성일/수정일 등 추가 방법

## Tags
`data-model`, `file-based`, `structure`

## Approved
- 2026-07-21: Option A (파일 기반), 사용자 승인