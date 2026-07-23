# jsondb 개발 대화 기록

> **프로젝트**: jsondb
> **날짜**: 2026-07-21
> **목적**: 에이전트와의 협업에서 파일 I/O 비효율성을 해결하기 위한 JSON 데이터베이스

---

## 대화 기록

### 1. 프로젝트 시작

**[user]** jsondb 만들어줘

**[cocrates]** jsondb 프로젝트를 시작하겠습니다. 먼저 현재 작업 공간을 확인하고, 필요한 폴더 구조를 만들겠습니다. jsondb가 어떤 것인지 이해하기 위해 몇 가지 질문을 드립니다: 1. 목적, 2. 사용 언어/플랫폼, 3. 핵심 기능, 4. 대상 사용자

---

### 2. 프로젝트 목적 설명

**[user]** 내가 jsondb를 만들려는 이유는 에이전트와 개발하다보니 파일로 데이터를 다루는데, 파일 I/O는 쉽고 명확하지만, 효율적이지는 않다. 이런 문제를 해결하기 위해서 데이터를 구조화해서 DB에 저장하고 에이전트가 효율적으로 데이터를 다룰 수 있도록 하기 위함이다.

**[cocrates]** 좋은 맥락 감사합니다. 좀 더 구체적으로 이해하기 위해 몇 가지 질문을 드립니다: 1. 데이터 유형, 2. "효율적"의 의미, 3. 스케일, 4. 언어/플랫폼, 5. 인터페이스

---

### 3. 문제점 구체화

**[user]** 예를 들어, overview에 제목이 있는데, 제목을 바꾸기 위해서 overview를 모두 읽고 파싱해야 하는 비효율이 가장 크지 않을까 싶다.

**[cocrates]** 아, 그 예시가 매우 명확합니다! 현재 문제점: overview 파일에서 title만 변경하려고 해도 전체 파일을 읽고 → 파싱하고 → 수정하고 → 다시 쓰는 과정이 필요. 좀 더 구체적으로 이해하기 위해: 1. 현재 파일 형식, 2. 파일 구조, 3. 빈도, 4. 다른 데이터

---

### 4. 데이터 형식 결정

**[user]** 현재는 markdown으로 작성했었는데, json 포맷으로 구조화 해서 DB에 넣으려고 하는 거지.

**[cocrates]** 이해했습니다! 현재 상태: Markdown 파일 → 목표 상태: JSON 구조화 → DB 저장. 이제 구체적인 요구사항을 정리하겠습니다: 1. 데이터 구조, 2. 필요한 연산, 3. 스케일, 4. 언어/플랫폼, 5. 인터페이스

---

### 5. 핵심 요구사항 정리

**[user]** 1. 데이터 구조: 복합, 2. 필요한 연산: key에 대한 값을 조회하고 변경하는 것이 가장 빈번하겠지. 3. 스케일은 크지 않아.

**[cocrates]** 핵심 요구사항이 명확해졌습니다: 복합 JSON 구조 저장, 키/필드 레벨 조회 및 수정, 소규모 스케일. 마지막으로 몇 가지 결정이 필요합니다: 1. 언어/플랫폼, 2. 인터페이스, 3. 추가 요구사항

---

### 6. 언어/플랫폼 비교 분석 요청

**[user]** 언어/플랫폼을 비교 분석해줘.

**[cocrates]** ADR 파일을 생성했습니다: `jsondb/adr/language-platform.md`. Python, Node.js, Go, Rust를 비교 분석하고 Python을 추천했습니다. 에이전트 통합 용이성, 개발 생산성, JSON 처리 편의성 등을 고려한 결과입니다.

---

### 7. Go 언어 선택

**[user]** 성능 및 배포 용이성이 중요하다. Go 언어로 하자.

**[cocrates]** ADR과 ASR을 업데이트했습니다. Go 언어 선택이 확정되었습니다. 다음 항목들을 결정해야 합니다: 1. 인터페이스 형태, 2. 데이터베이스 엔진

---

### 8. 인터페이스 유형 논의

**[user]** 어떤 인터페이스 유형이 좋을까?

**[cocrates]** 각 인터페이스 유형(라이브러리, CLI, 서버)의 장단점을 비교 분석하고, 라이브러리 + CLI 조합을 추천했습니다. 에이전트 통합, 성능, 유연성 등을 고려한 결과입니다.

---

### 9. 동시성 문제 제기

**[user]** 그런데, 복수 서브 에이전트가 동시에 접근하면?

**[cocrates]** 아주 좋은 질문입니다! 동시성은 반드시 고려되어야 하는 아키텍처적 요구사항입니다. 새로운 ASR(ASR-006)을 등록하고, 동시성 해결 방안(파일 레벨 잠금, 트랜잭션, 단일 서버 프로세스 등)을 분석했습니다.

---

### 10. 기존 DB 사용 반대

**[user]** 아, SQLite 등 DB를 backend로 사용하면 jsondb를 개발하는 의미가 없지 않나.

**[cocrates]** 맞습니다! 기존 DB를 백엔드로 사용하면 단순 래퍼가 되어버리죠. jsondb의 의미: 기존 DB의 복잡성 없이, JSON 구조를 네이티브로 다루면서, 필드 레벨 접근이 가능한 경량 솔루션. 저장 방식에 대한 대안 방식을 제시했습니다.

---

### 11. 저장 방식과 동시성 관계

**[user]** 데이터를 어떻게 저장할지도 당연히 검토되어야 하겠지만... 인터페이스부터 다시 결정해 보자. CLI로 이 문제를 해결할 수 있나?

**[cocrates]** CLI alone으로는 근본적인 문제를 해결할 수 없습니다. CLI는 인터페이스일 뿐, 저장 방식이 변하지 않으면 효율성이 개선되지 않습니다. 인터페이스와 저장 방식은 별개의 관심사입니다.

---

### 12. 저장 방식과 동시성 관계 질문

**[user]** 저장 방식에 대한 접근이 성능에는 영향을 미치지만, 프로세스간의 동시성 문제도 해결이 되나?

**[cocrates]** 저장 방식과 동시성은 별개의 문제입니다. 저장 방식은 성능에, 동시성 제어는 데이터 무결성에 영향을 미칩니다. 동시성 해결 방안(파일 잠금, WAL, 단일 서버 프로세스 등)을 분석했습니다.

---

### 13. CLI vs 서버 비교 분석

**[user]** CLI와 서버 인터페이스의 장단점을 분석해줘.

**[cocrates]** ADR 파일을 생성했습니다: `jsondb/adr/interface-type.md`. CLI와 서버의 장단점을 비교 분석하고, 하이브리드 방식(CLI + Server)을 추천했습니다.

---

### 14. 하이브리드 방식 결정

**[user]** 하이브리드로 하자. CLI는 Server에 요청하고, Server는 jsondb 엔진 모듈과 HTTP 인터페이스를 제공하는 서버 모듈로 구분된다.

**[cocrates]** 훌륭한 결정입니다! 명확한 아키텍처 분리가 되었습니다. ADR과 ASR을 업데이트하고, 다음 단계로 엔진 아키텍처 설계를 제안했습니다.

---

### 15. 데이터 모델 검토

**[user]** 먼저 데이터 모델을 검토해 보자.

**[cocrates]** 데이터 모델 비교 분석 ADR을 생성했습니다: `jsondb/adr/data-model.md`. 파일 기반 접근, 컬렉션 기반 접근, 하이브리드 접근을 비교하고, 파일 기반 접근을 추천했습니다.

---

### 16. 파일 기반 데이터 모델 결정

**[user]** 파일 기반이 적절한 것 같아. 파일 시스템처럼.

**[cocrates]** 파일 기반 데이터 모델이 결정되었습니다. ADR과 ASR을 업데이트하고, 다음 단계로 파일 I/O 방식 분석을 제안했습니다.

---

### 17. 파일 I/O 방식 분석

**[user]** 파일 I/O 방식에 대해서 분석해줘.

**[cocrates]** ADR 파일을 생성했습니다: `jsondb/adr/file-io-method.md`. 직접 I/O, mmap, 버퍼링, mmap+인덱스, WAL 방식을 비교 분석하고, 단계적 접근(직접 I/O → mmap → WAL)을 추천했습니다.

---

### 18. Storage Backend 추상화 결정

**[user]** 지금은 직접 I/O로 하고 나중에 mmap이나 WAL로 변경하기 용이하도록 storage backend를 추상화하는게 좋겠어.

**[cocrates]** 훌륭한 결정입니다! Storage Backend 추상화는 확장성과 유지보수성에 매우 중요합니다. ADR에 Storage Backend 인터페이스 설계를 추가하고, 업데이트했습니다.

---

### 19. JSON 작업 방식 검토

**[user]** json object는 매우 복잡한 복합 객체이다. json object/file의 조회, 검색, 변경 기능을 제공해야 하는데, 이 부분에 설계 결정해야 할 것은 없나?

**[cocrates]** 맞습니다! JSON 객체의 복잡성을 고려하면, 조회/검색/변경 방식에 대한 설계 결정이 매우 중요합니다. ADR을 생성했습니다: `jsondb/adr/json-operations.md`. 단순 경로, JSONPath, 혼합, 함수형 접근을 비교 분석하고, 혼합 접근을 추천했습니다.

---

### 20. JSONPath 지원 질문

**[user]** 1. 에이전트가 JSONPath를 사용하는데 문제가 있을까? 2. Go 언어로 JSONPath를 구현하는 방법은?

**[cocrates]** 1. JSONPath는 에이전트가 이미 사용하는 표준이므로 큰 문제 없음. 2. Go 언어에서는 `github.com/ohler55/ojg` 라이브러리를 사용하면 됨. 검증된 구현체가 존재합니다.

---

### 21. JSON Patch/Merge Patch 지원 결정

**[user]** 검토한 내용도 adr에 보완해 주고, 기본 경로 + JSONPath를 모두 지원하는 것으로 결정해줘.

**[cocrates]** ADR에 JSONPath 연구 결과와 Go 구현 방법을 추가하고, 기본 경로 + JSONPath 혼합 지원으로 결정했습니다.

---

### 22. 엔진 아키텍처 설계

**[user]** 진행해줘

**[cocrates]** ADR 파일을 생성했습니다: `jsondb/adr/engine-architecture.md`. 엔진 모듈의 전체 구조, 모듈 구성, 상세 설계(Engine 인터페이스, Query Parser, Query Executor, Storage Backend), 데이터 흐름, 에러 처리 등을 설계했습니다.

---

### 23. 엔진 아키텍처 승인

**[user]** 오케이

**[cocrates]** 엔진 아키텍처가 결정되었습니다. ADR과 ASR을 업데이트하고, 다음 단계로 동시성 제어 방식 결정을 제안했습니다.

---

### 24. 동시성 제어 방식 결정

**[user]** 진행해줘

**[cocrates]** ADR 파일을 생성했습니다: `jsondb/adr/concurrency-control.md`. 파일 잠금, RWLock, 혼합, 버전 기반 잠금 방식을 비교 분석하고, 파일 잠금 + 메모리 잠금 혼합(Option C)을 추천했습니다.

---

### 25. 동시성 제어 승인

**[user]** 좋아 Option C로 하자

**[cocrates]** 동시성 제어 방식이 결정되었습니다. ADR과 ASR을 업데이트하고, 다음 단계로 HTTP API 설계를 제안했습니다.

---

### 26. HTTP API 설계

**[user]** 오케이

**[cocrates]** ADR 파일을 생성했습니다: `jsondb/adr/http-api.md`. RESTful API 설계, 엔드포인트 구조, 요청/응답 형식, 에러 처리 등을 설계했습니다.

---

### 27. 부분 업데이트 사양 질문

**[user]** 부분 업데이트 사양에 대해 좀 더 구체적으로 알려줘.

**[cocrates]** 부분 업데이트(PATCH)의 상세 사양을 설명했습니다. 지원 연산(Set, Delete, Append, Remove, Merge), 중첩된 필드 변경, 배열 조작, 충돌 처리 등을 상세히 설명했습니다.

---

### 28. 공식 표준 질문

**[user]** 부분 변경에 대한 사양이 있나?

**[cocrates]** 공식 표준인 RFC 6902(JSON Patch)와 RFC 7396(JSON Merge Patch)를 설명했습니다. 두 표준의 비교와 장단점을 분석했습니다.

---

### 29. Go 구현 가능성 질문

**[user]** go 언어로 구현하는데 문제가 없나?

**[cocrates]** Go 언어로 구현하는 데 큰 문제는 없습니다. `github.com/evanphx/json-patch` 라이브러리를 사용하면 JSON Patch와 Merge Patch 모두 지원됩니다. 검증된 라이브러리가 존재합니다.

---

### 30. JSON Patch ADR 작성 요청

**[user]** json patch 관련해서 adr을 작성하고, 두가지 모두 지원하고, json-patch 라이브러리를 사용하는 것으로 결정해줘.

**[cocrates]** ADR 파일을 생성했습니다: `jsondb/adr/json-patch.md`. JSON Patch(RFC 6902)와 JSON Merge Patch(RFC 7396)를 모두 지원하고, `github.com/evanphx/json-patch` 라이브러리를 사용하는 것으로 결정했습니다.

---

### 31. 스펙 기반 구현 검증 요청

**[user]** spec대로 구현되었는지 확인해줘.

**[cocrates]** `verification/full-verification.md`를 생성하여 PRD, engine, server, cli 스펙과 실제 구현을 대조 검증했습니다. 28 pass, 14 fail, 6 partial, 2 not-verifiable 결과. 핵심 이슈: CLI가 서버를 거치지 않고 엔진을 직접 사용하는 아키텍처 위반.

---

### 32. CLI 서버 연결 수정 요청

**[user]** CLI가 서버를 사용하도록 수정해줘.

**[cocrates]** `cmd/main.go`를 수정하여 모든 CLI 명령어(get, set, delete, patch, ls, create, rm)가 HTTP 요청(`httpRequest()`)을 통해 서버와 통신하도록 변경했습니다. `serve` 명령어만이 Engine을 직접 생성하여 서버를 시작합니다. 검증 완료.

---

### 33. CLI-Engine 분리 수정 확인

**[user]** spec-driven-development 스킬을 사용해서 jsondb를 개발했음. 현재 CLI가 engine을 직접 사용하는 문제에 대해서 수정중이었는데, 수정이 완료되었는지 다시 확인해줘.

**[cocrates]** 현재 코드를 검토한 결과, CLI-Engine 분리 수정은 이미 완료되었습니다. 모든 데이터 명령어(get/set/delete/patch/ls/create/rm)가 `httpRequest()` 함수를 통해 HTTP API로 서버와 통신하며, `serve` 명령어만이 Engine을 직접 생성합니다. 단일 바이너리 배포 모델에 부합하는 정상 상태입니다.

---

### 34. 매 요청마다 Open/Close 이슈 검토

**[user]** 매 요청마다 Open/Close 하는 이슈에 대해 검토해줘.

**[cocrates]** 서버 핸들러의 모든 요청이 `s.engine.Open(path)` → 작업 → `defer s.engine.Close()` 패턴을 사용하고 있음을 발견했습니다. 두 가지 핵심 문제: 1) **성능** — 읽기 전용 요청도 매번 파일 전체를 디스크에서 읽고 JSON 파싱, 2) **정합성** — 단일 Engine 인스턴스의 `data map` 공유로 Race Condition 발생. 이 문제를 해결하기 위한 ADR 작성 필요.

---

### 35. 서버 엔진 생명주기 ADR 작성

**[user]** adr을 작성해서 검토해줘.

**[cocrates]** ADR 파일을 생성했습니다: `jsondb/adr/server-engine-lifecycle.md`. 4가지 옵션 비교: A) 파일별 Long-lived Engine, B) 요청별 독립 Engine Pool, C) In-Memory Cache, D) Long-lived + Write-through Cache. Option D 추천: 파일별 Engine 유지 + Write-through로 성능과 정합성 동시 해결.

---

### 36. Option D 승인

**[user]** 좋아. Option D로 결정해줘.

**[cocrates]** Option D(Long-lived Engine + Write-through Cache)를 승인했습니다. ADR 상태를 `approved`로 변경하고, 관련 ASR(ASR-006 동시성 제어, ASR-007 엔진 아키텍처)을 `designed`로 업데이트했습니다. Downstream Concerns로 4가지 후속 결정 항목 식별.

---

### 37. 즉시 구현 필수 항목 진행

**[user]** 즉시 구현에 필요한 항목 먼저 진행해줘.

**[cocrates]** Downstream Concerns 중 즉시 구현에 필요한 2가지 항목에 대해 ADR을 작성: 1) `concurrent-write-serialization.md` — 파일 경로 기반 RWLock(Option B) 승인, 2) `server-shutdown-cleanup.md` — signal 핸들러 + Graceful Shutdown(Option A) 승인. ASR 레지스트리 업데이트 완료.

---

### 38. 승인 및 구현 요청

**[user]** 모두 승인하고 구현해줘.

**[cocrates]** 두 ADR을 모두 승인 처리하고 서버 모듈을 전면 재구현했습니다. 주요 변경: 1) Server 구조체에 `engines`, `fileLocks` 맵 추가, 2) `getOrCreateEngine()` — lazy Engine 생성 및 재사용, 3) `getFileLock()` — 파일별 RWMutex 관리, 4) 모든 핸들러에서 Open/Close 제거, 5) Graceful Shutdown — signal 수신 후 30초 타임아웃으로 정리. 빌드 및 테스트 통과.

---
