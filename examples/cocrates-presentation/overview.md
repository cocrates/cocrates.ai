# Cocrates Harness 블로그 시리즈 — 14개 에피소드 발표자료 제작 계획

## 프로젝트 개요

`examples/blog/episodes/`에 있는 14편의 블로그 에피소드를 각각 **독립된 발표자료**로 제작한다. 각 발표자료는 `presentation-authoring` 스킬의 Snowflake 5단계(S1~S5)를 따라 단계별로 설계·검토·승인 후 생성한다.

### 발표자료의 목적

| 항목 | 내용 |
|------|------|
| **Topic** | Cocrates Harness 블로그 시리즈 — AI를 올바르게 쓰는 법 |
| **Purpose** | 14개 각 에피소드의 핵심 메시지와 스토리라인을 청중이 이해할 수 있도록 발표자료로 전달한다. 각 에피소드는 하나의 독립 발표 세션으로 설계한다. |
| **Audience** | 소프트웨어 개발자(초급~중급), AI 도구를 업무에 활용하는 지식 노동자. LLM 사용 경험은 있으나 구조적 접근이 필요한 독자. |
| **Format** | 각 에피소드별 10~20분 분량의 발표 자료. 슬라이드 + 대본 + Q&A 자료 포함. |

---

## 디렉토리 구조

```
examples/presentation/
├── overview.md                         # 이 파일 — 전체 계획
│
├── docs/                               # 공통 개념/참고 문서 (전 에피소드 공유)
│   ├── ai-native-engineer.md           # AI-native Engineer 정의
│   └── ...                             # 발표자료 제작 중 발견된 공통 개념 추가
│
├── episode-01/                         # 같은 LLM, 다른 결과
│   ├── overview.md                     # S1: 발표 정의
│   ├── outline.md                      # S2: 발표 개요
│   ├── slides.md                       # S3: 슬라이드 계획
│   ├── scripts/                        # S4: 대본 (슬라이드당 1개)
│   │   ├── script-01.md
│   │   └── ...
│   ├── slides/                         # S5: 슬라이드 콘텐츠 명세
│   │   ├── slide-01.md
│   │   └── ...
│   ├── verification.md                 # S5: 검증 보고서
│   └── q-and-a.md                      # S5: 예상 질의응답
│
├── episode-02/                         # The Unexamined Code Is Not Worth Generating
│   └── ...                             # (동일 구조)
│
├── episode-03/                         # Cocrates Harness 설치
│   └── ...
│
... (episode-04 ~ episode-13)
│
└── episode-14/                         # Cocrates 사용자 선언문
    └── ...                             # (동일 구조)
```

> **`presentation/docs/`** 는 모든 에피소드에서 공통으로 참조할 개념 정의, 용어 정리, 배경 지식 등을 저장한다. 발표자료 제작 과정에서 발견된 공통 개념은 이곳에 문서화하여 에피소드 간 일관성을 유지한다.

---

## 에피소드 목록

14개 에피소드는 총 4개 파트로 구성된 시리즈 아크를 따른다.

### Part 1. 소개 — Cocrates Harness의 필요성과 원칙

| # | 슬러그 | 제목 | 예상 분량 | 핵심 메시지 |
|---|--------|------|-----------|------------|
| 01 | `episode-01` | 같은 LLM, 다른 결과 | 12분 | 모델이 아니라 사용자의 태도가 차이를 만든다 |
| 02 | `episode-02` | The Unexamined Code Is Not Worth Generating | 10분 | 검토되지 않은 산출물은 생성할 가치가 없다 |

### Part 2. 실습 — 먼저 경험한다

| # | 슬러그 | 제목 | 예상 분량 | 핵심 메시지 |
|---|--------|------|-----------|------------|
| 03 | `episode-03` | Cocrates Harness 설치 | 12분 | 설치는 단순하지만, 진짜 시작은 설치 후다 |
| 04 | `episode-04` | 소크라테스식 학습 실습 | 18분 | Learning 파이프라인의 3단계를 실제 대화로 경험 |
| 05 | `episode-05` | 구조 기반 생성 실습 | 22분 | 구조 기반 생성은 적극적인 설계 과정이다 |
| 06 | `episode-06` | 구조 기반 스킬 생성 실습 | 20분 | 스킬 생성은 공동 설계 과정이다 |

### Part 3. 구조 — 원리를 이해한다

| # | 슬러그 | 제목 | 예상 분량 | 핵심 메시지 |
|---|--------|------|-----------|------------|
| 07 | `episode-07` | Cocrates Harness 구조 | 15분 | Agent + Skills 구조로 산출물별 최적화 + 지속적 개선 가능 |
| 08 | `episode-08` | 소크라테스식 학습 원리 | 12분 | Learning 파이프라인: Education → Knowledge Capture → Reflection |
| 09 | `episode-09` | 소크라테스식 학습 스킬 | 15분 | 세 스킬(education/knowledge-capture/reflection)의 연결과 순환 |
| 10 | `episode-10` | 구조 기반 생성 원리 | 15분 | Artifact Generation 파이프라인: ASR → ADR → Spec → 생성 → 검증 |
| 11 | `episode-11` | 구조 기반 생성 스킬 | 18분 | 4개 스킬(adr/spec/generation/verification)의 역할과 원칙 |
| 12 | `episode-12` | 구조 기반 스킬 생성 원리 | 16분 | generating-skill-creation: 스킬을 생성하는 메타-스킬 |

### Part 4. 결론 — 지속적 개선과 선언

| # | 슬러그 | 제목 | 예상 분량 | 핵심 메시지 |
|---|--------|------|-----------|------------|
| 13 | `episode-13` | 올바른 Cocrates Harness 활용 | 12분 | 사용자도 에이전트도 함께 진화해야 한다 |
| 14 | `episode-14` | Cocrates 사용자 선언문 | 10분 | 검토되지 않은 산출물은 생성할 가치가 없다 — 7계명 |

---

## 제작 프로세스 (Presentation-authoring Skill 적용)

각 에피소드별 발표자료는 `presentation-authoring` 스킬의 Snowflake 5단계를 순차적으로 적용한다.

| 단계 | 이름 | 주요 활동 | 산출물 | 승인 필요 |
|------|------|-----------|--------|----------|
| **S1** | Define | 발표 정의: 주제, 목적, 청중, 시간, 성공 기준 확정 | `overview.md` | ✅ 사용자 승인 |
| **S2** | Plan | 에피소드 내용 분석 → 발표 개요 설계 (도입-본론-결론) | `outline.md` | ✅ 사용자 승인 |
| **S3** | Architecture Design | 개요를 슬라이드 단위로 분해, 각 슬라이드 목적과 내용 정의 | `slides.md` | ✅ 사용자 승인 |
| **S4** | Detail Design | 슬라이드별 대본 작성 (영화/드라마 대본 형식, 1:1 매칭) | `scripts/script-NN.md` | ✅ 사용자 승인 |
| **S5** | Generation | 슬라이드 콘텐츠 명세 생성 + 검증 + Q&A + 부록 | `slides/slide-NN.md`, `verification.md`, `q-and-a.md`, `appendix/` | ✅ 사용자 최종 승인 |

### 제작 순서

1. **episode-01**부터 **episode-14**까지 순차적으로 제작
2. 각 에피소드 내에서는 S1 → S2 → S3 → S4 → S5 단계를 순차적으로 진행
3. 각 단계 완료 후 사용자 승인을 받은 후 다음 단계로 진행
4. 모든 에피소드의 S5 완료 후 전체 최종 검토

---

## 고정 슬라이드 템플릿 (모든 에피소드 공통)

각 에피소드 발표자료는 다음 슬라이드 구성을 **고정 템플릿**으로 따른다.

```
슬라이드 1:    타이틀 페이지
               - 이전 에피소드 언급 (시리즈 맥락)
               - 이번 에피소드 제목 및 간단한 소개

슬라이드 2:    학습 목표
               - 이번 에피소드를 통해 알게 될 내용 / 체크리스트

슬라이드 3~N:  에피소드 내용
               - 블로그 에피소드의 스토리라인과 구조를 따름
               - 내용에 따라 슬라이드 내부 구성은 유연하게 (강제 구조 없음)

슬라이드 N+1:  학습 목표 확인
               - 슬라이드 2의 체크리스트를 질문/퀴즈/미션 형태로 전환
               - 소크라테스 방식: 시청자가 스스로 답을 생각해보도록 유도

슬라이드 N+2:  다음 에피소드 연결
               - 다음 에피소드 주제 및 시리즈 내 위치 안내
```

### 공통 원칙

- **Part 1 (소개):** 청중의 공감과 문제 의식을 형성 — "지금 나는 A님일까 B님일까?"라는 질문을 던짐
- **Part 2 (실습):** 실제 대화 기반 스토리텔링으로 현장감 전달 — 구체적 사례와 수치(예: 72개 항목 검증, 70 pass)를 시각화
- **Part 3 (구조):** 개념적 깊이를 유지하되 비유(주방, 전원주택)로 추상성을 낮춤 — Agent + Skills 구조, 파이프라인 다이어그램 활용
- **Part 4 (결론):** 7계명을 카드 형식으로 제시 — 시리즈 전체를 관통하는 핵심 원칙 3가지 재강조

---

## 진행 현황

- [x] **episode-01** — 같은 LLM, 다른 결과
- [x] **episode-02** — The Unexamined Code Is Not Worth Generating
- [x] **episode-03** — Cocrates Harness 설치
- [x] **episode-04** — 소크라테스식 학습 활동 - 실습
- [x] **episode-05** — 구조 기반 산출물 생성 활동 - 실습습
- [x] **episode-06** — 구조 기반 워크플로우 생성 실습
- [x] **episode-07** — Cocrates Harness 구조
- [x] **episode-08** — 소크라테스식 학습 활동 - 원리
- [x] **episode-09** — 소크라테스식 학습 활동 - 스킬
- [x] **episode-10** — 구조 기반 산출물 생성 활동 - 원리
- [x] **episode-11** — 구조 기반 산출물 생성 활동 - 스킬
- [x] **episode-12** — 구조 기반 워크플로우 생성 스킬
- [x] **episode-13** — 올바른 Cocrates Harness 활용
- [x] **episode-14** — Cocrates 사용자 선언문

---

*이 계획은 `presentation-authoring` 스킬의 Snowflake 5단계(S1: Define → S2: Plan → S3: Architecture Design → S4: Detail Design → S5: Generation + Verification)를 각 에피소드 발표자료에 적용한다.*
