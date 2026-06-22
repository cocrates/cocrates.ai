# Cocrates Harness 설치

---

지난 두 편에서 우리는 Cocrates가 무엇인지, 어떤 원칙으로 움직이는지 알아봤습니다.

> *"검토되지 않은 산출물은 생성할 가치가 없다."*

이 원칙을 실현하는 도구, Cocrates Harness. 그렇다면 어떻게 시작해야 할까요?

이제 직접 설치해봅시다.

---

## Cocrates Harness는 opencode 플러그인이다

Cocrates Harness는 독립 실행형 프로그램이 아닙니다. **opencode**라는 플랫폼 위에서 동작하는 **플러그인**입니다.

- **opencode**: AI 에이전트를 실행하고 관리하는 플랫폼
- **@cocrates/cocrates-harness**: opencode 위에서 Cocrates Agent와 Skills를 제공하는 플러그인

그러니까 설치 과정은 두 단계로 나뉩니다. opencode를 먼저 설치하고, 그 위에 Cocrates Harness 플러그인을 추가하는 것입니다. 거기에 Cocrates의 스킬 파일들까지 복사하면 설치 완료입니다.

간단하죠?

---

## 1단계: opencode 설치

opencode는 세 가지 방식으로 설치할 수 있습니다. 사용자 환경에 따라 골라잡으면 됩니다.

| 설치 방식 | 추천 대상 | 특징 |
|----------|----------|------|
| **opencode 터미널** | Claude/GPT TUI에 익숙한 사용자 | 터미널 기반, 키보드 중심 |
| **opencode desktop** | Cursor, VS Code 스타일을 선호하는 사용자 | GUI 에디터 환경 |
| **OpenCode 확장** | 기존 Cursor/Code 환경을 유지하고 싶은 사용자 | 확장 기능으로 추가 |

선택지는 셋이지만, 핵심은 같습니다. **어떤 방식을 선택하든 Cocrates Harness는 동일하게 동작합니다.** 자신에게 가장 편한 방법을 고르세요.

설치 방법은 공식 사이트를 참고하세요.

> **https://opencode.ai/download**

---

## 2단계: cocrates-harness plugin 설치

opencode 설치가 끝났다면, 이제 Cocrates Harness 플러그인을 추가할 차례입니다.

플러그인 설치는 opencode의 **config 파일**을 수정하는 것으로 이루어집니다.

### config 파일 위치

opencode는 설정 파일을 다음 위치에서 찾습니다.

- **기본 위치**: `~/.config/opencode/`
- **환경변수 설정 시**: `OPENCODE_CONFIG_DIR` 변수가 가리키는 경로

### opencode.jsonc 수정

config 디렉토리 안의 `opencode.jsonc` 파일을 열고, `plugin` 항목에 `"@cocrates/cocrates-harness"`를 추가합니다.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "plugin": ["@cocrates/cocrates-harness"]
}
```

---

## 3단계: skill 파일 복사

플러그인만으로는 Cocrates가 완전히 동작하지 않습니다. Cocrates의 핵심 역량은 **스킬(Skills)** 에 있기 때문입니다. 스킬은 각 작업(학습, 구조 기반 생성 등)을 위한 구체적인 워크플로우를 정의한 파일들입니다.

이 스킬 파일들을 opencode의 skills 디렉토리에 복사해야 합니다.

### 복사할 스킬 목록

Cocrates Harness는 현재 다음과 같은 스킬을 제공합니다.

| 스킬 | 역할 |
|------|------|
| `education` | 소크라테스식 1:1 학습 코치 |
| `knowledge-capture` | 학습 내용을 KB로 저장 |
| `reflection` | 이해도 평가 |
| `adr-writing` | 구조적 의사결정 기록 |
| `spec-writing` | 요구사항 명세 통합 |
| `spec-driven-generation` | Spec 기반 산출물 생성 |
| `spec-driven-verification` | Spec 기반 검증 |

### 복사 방법

```
https://github.com/cocrates/cocrates.ai/tree/main/skills
```

위 GitHub 저장소의 `skills/` 폴더에 있는 모든 스킬을 `~/.config/opencode/skills/` 디렉토리로 복사하면 됩니다.

결과적으로 디렉토리 구조는 다음과 같이 됩니다.

```
~/.config/opencode/
├── opencode.jsonc
└── skills/
    ├── education/SKILL.md
    ├── knowledge-capture/SKILL.md
    ├── reflection/SKILL.md
    ├── adr-writing/SKILL.md
    ├── spec-writing/SKILL.md
    ├── spec-driven-generation/SKILL.md
    ├── spec-driven-verification/SKILL.md
    └── ...
```

### AI에게 설치를 요청하는 방법

사실, 이 작업을 직접 하실 필요는 없습니다. opencode를 설치하고 실행한 다음, AI에게 이렇게 요청하세요.

> "https://cocrates.ai/install.md 파일을 근거로 Cocrates Harness를 설치해줘."

그러면 AI가 이 문서를 근거로 플러그인 설치와 스킬 복사를 도와줄 것입니다.

---

## 설치 확인

설치가 완료되었다면, opencode를 재시작하고 Cocrates Agent가 활성화되었는지 확인해보세요.

---

## 설치만으로는 아무 일도 일어나지 않는다

설치는 시작일 뿐입니다.

Cocrates를 설치했다고 해서 자동으로 당신의 AI 활용 방식이 바뀌지는 않습니다. Cocrates는 **당신의 참여**를 기다립니다. 질문에 답하고, 구조를 검토하고, 결정을 내리는 것은 언제나 당신의 몫입니다.

이 시리즈의 다음 편들에서 우리는 Cocrates와 함께 직접 실습해볼 것입니다.

- **Ep4**: Cocrates와 함께 블룸의 분류학을 학습하는 과정을 따라갑니다
- **Ep5**: ADR부터 검증까지 구조 기반 생성의 전 과정을 경험합니다
- **Ep6**: 직접 스킬을 만들어보는 워크플로우 생성 실습을 합니다

하지만 그 전에, 지금 당장 Cocrates와 첫 대화를 시작해보세요.

설치가 끝났다면, 이제 Cocrates에게 말을 걸 차례입니다. Cocrates는 당신의 참여를 기다리고 있습니다.

> **"블룸의 분류학에 대해 알려줘."**

어떤 일이 벌어질지 궁금하지 않나요?

---

*이 시리즈는 Cocrates Harness 프레임워크를 소개합니다. Cocrates는 소크라테스식 대화로 사용자가 주도권을 잡고 성장하도록 설계된 에이전트 하네스입니다.*

*→ 다음 편: 소크라테스식 학습 활동 — 블룸의 분류학에 대해 학습하는 과정을 경험한다*