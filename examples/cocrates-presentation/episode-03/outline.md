# Presentation Outline — Episode 03

## Metadata
- **Title:** Cocrates Harness 설치 — AI-native Engineer의 첫걸음
- **Total estimated time:** ~12분
- **Number of slides:** 12

## Structure

### 1. 도입 (슬라이드 1~2) — 약 2분

#### 슬라이드 1: 타이틀 페이지
- **Key message:** Ep2에서 "검토되지 않은 산출물은 생성할 가치가 없다"는 원칙을 배웠다. 이제 그 원칙을 실천할 도구를 설치할 차례다.
- **Estimated time:** 1분
- **Key points:**
  - Ep2 복습: "The unexamined code is not worth generating"
  - "원칙은 알겠다. 그런데 어떻게 시작하지?"
  - 오늘의 답변: **Cocrates Harness 설치**
  - "설치는 5분이면 끝나지만, 진짜 시작은 설치 후다"

#### 슬라이드 2: 학습 목표
- **Key message:** 이 에피소드를 마치면 (1) opencode와 Cocrates의 관계를 이해하고 (2) 3단계 설치 과정을 알고 (3) 직접 설치할 수 있다.
- **Estimated time:** 1분
- **Key points:**
  - opencode(플랫폼)와 Cocrates(플러그인)의 관계를 이해한다
  - 3단계 설치 과정(opencode → plugin → skills)을 설명할 수 있다
  - 자신의 환경에 맞는 방법으로 Cocrates Harness를 직접 설치한다

---

### 2. 본론 (슬라이드 3~10) — 약 8분

#### 2.1 구조 이해 (슬라이드 3~4) — 2분

##### 슬라이드 3: Cocrates Harness의 정체
- **Key message:** Cocrates Harness는 단독 실행 프로그램이 아니라, opencode 플랫폼 위에서 동작하는 플러그인이다.
- **Estimated time:** 1분
- **Key points:**
  - **opencode:** AI 에이전트를 실행하고 관리하는 플랫폼 (무대)
  - **Cocrates Harness:** 그 플랫폼 위에서 특별한 능력을 발휘하는 플러그인 (배우)
  - 비유: "극장(플랫폼)과 배우(플러그인)"
  - 설치 = 3단계: opencode 설치 → plugin 설정 → skill 파일 복사

##### 슬라이드 4: 3가지 UI 선택지
- **Key message:** opencode는 3가지 방식으로 사용할 수 있다. 초보자에게는 VS Code + OpenChamber를 권장한다.
- **Estimated time:** 1분
- **Key points:**
  - **opencode 터미널 (TUI):** 터미널 기반, 키보드 중심, 힙스터 취향
  - **opencode 데스크탑:** 설치형 GUI 앱. **단, 아직 베타** — 기능/안정성 불안정
  - **VS Code + OpenChamber 확장:** 기존 VS Code 환경에 GUI 추가. **초보자 추천**
  - 하지만 어떤 방식을 선택하든 **opencode CLI 설치는 공통**

#### 2.2 Step 1: opencode 설치 (슬라이드 5~6) — 2분

##### 슬라이드 5: Step 1 — opencode 설치 (macOS / Linux)
- **Key message:** macOS와 Linux는 `curl`과 `bash`가 기본 내장되어 있어, install script 한 줄이면 설치 끝.
- **Estimated time:** 1분
- **Key points:**
  - macOS는 Terminal.app에 curl/bash 기본 내장
  - Linux도 대부분 curl 기본 포함
  - 명령어 한 줄: `curl -fsSL https://opencode.ai/install | bash`
  - 설치 확인: `opencode --version`
  - **"Homebrew를 먼저 설치할 필요 없습니다"**

##### 슬라이드 6: Step 1 — opencode 설치 (Windows)
- **Key message:** Windows는 Chocolatey로 설치한다. Chocolatey 자체 설치가 필요하다.
- **Estimated time:** 1분
- **Key points:**
  - PowerShell을 관리자 권한으로 실행
  - Chocolatey 설치 명령어 (1줄 복붙)
  - `choco install opencode`
  - 대안: WSL + install script (opencode 공식 권장)
  - **"choco는 Windows의 앱스토어 같은 패키지 매니저입니다"**

#### 2.3 GUI 환경 구성 (슬라이드 7) — 1분

##### 슬라이드 7: VS Code + OpenChamber 확장 설치
- **Key message:** TUI가 익숙하지 않다면 VS Code에 OpenChamber 확장을 설치하여 GUI 환경에서 opencode를 사용할 수 있다.
- **Estimated time:** 1분
- **Key points:**
  - **VS Code 설치:** https://code.visualstudio.com 에서 다운로드 → Next만 누르면 끝
  - **OpenChamber 확장 설치:** VS Code 실행 → 확장(Ctrl+Shift+X) → "OpenChamber" 검색 → 설치
  - "마치 Copilot이나 Cursor처럼 VS Code에서 Cocrates를 사용할 수 있습니다"

#### 2.4 설치 + 검증 (슬라이드 8~9) — 2분

##### 슬라이드 8: AI에게 설치 시키기
- **Key message:** 직접 설치하기 귀찮다면 AI에게 시키면 된다. 우리는 AI를 지휘하는 디렉터니까.
- **Estimated time:** 1분
- **Key points:**
  - opencode 설치 후, 한 마디만 던지면 된다
  - 👤 **유저:** *"[https://cocrates.ai/install.md](https://cocrates.ai/install.md) 파일을 근거로 Cocrates Harness를 설치해 줘."*
  - AI가 plugin 설정부터 skill 파일 다운로드까지 알아서 처리
  - 하지만 **"설치조차 검토하고 승인하세요"** — AI가 설치한 내용을 확인하는 것이 진짜
  - 이제 다음 슬라이드에서 확인 방법을 알려드립니다

##### 슬라이드 9: 설치 확인 — Cocrates = Plugin + Skill
- **Key message:** Cocrates Harness는 Plugin + Skill로 구성된다. 설치가 제대로 되었는지 이 두 가지를 확인하면 된다.
- **Estimated time:** 1분
- **Key points:**
  - **Cocrates = Plugin + Skill** (두 가지로 구성, 하나만 있어도 불완전)
  - **Plugin 확인:** `~/.config/opencode/opencode.jsonc`에 `"plugin": ["@cocrates/cocrates-harness"]`가 있는가?
  - **Skill 확인:** `~/.config/opencode/skills/` 디렉토리에 skill 파일들이 있는가?
  - "AI가 설치했다면, 이 두 가지만 검토(U→A→E→A)하면 됩니다"
  - 둘 다 정상 → opencode 재시작 → Cocrates Agent 선택 → 활성화 완료

#### 2.5 첫 대화 (슬라이드 10) — 1분

##### 슬라이드 10: 설치 완료! 첫 대화
- **Key message:** 설치가 끝났다면 첫 질문을 던져보자. "블룸의 분류학에 대해 알려줘"
- **Estimated time:** 1분
- **Key points:**
  - opencode 재시작 → Cocrates Agent 선택 → 활성화
  - 첫 질문: 💬 **"블룸의 분류학에 대해 알려줘"**
  - 일반 AI와 다를 것 — Cocrates는 곧바로 답을 주지 않는다
  - **"설치만으로 끝나지 않습니다. 진짜는 지금부터입니다."**

---

### 3. 결론 (슬라이드 11~12) — 약 2분

#### 슬라이드 11: 핵심 요약 + 학습 확인
- **Key message:** 오늘의 3가지 핵심 포인트. 그리고 스스로에게 질문하라.
- **Estimated time:** 1분
- **Key points:**
  - **3가지 요약:**
    1. **Cocrates는 opencode 플러그인이다.** — opencode(플랫폼) 위에서 동작하는 플러그인. Plugin + Skill로 구성된다.
    2. **opencode의 권장 설치는 opencode + VS Code + OpenChamber 확장이다.** — TUI가 익숙하지 않다면 VS Code + OpenChamber로 GUI 환경에서 사용한다.
    3. **Cocrates의 설치는 AI에게 요청한다.** — "install.md를 근거로 Cocrates Harness를 설치해 줘" 한 마디면 끝. 단, 설치 후 Plugin + Skill을 반드시 확인(검토)한다.
  - **학습 확인 질문:**
    - "지금 opencode가 설치되어 있는가?"
    - "Cocrates Agent가 활성화되어 있는가?"
    - "첫 질문을 던질 준비가 되었는가?"

#### 슬라이드 12: 다음 에피소드 연결
- **Key message:** 설치가 끝났다면 이제 진짜 경험을 시작할 차례. Ep4에서는 Cocrates와의 첫 대화를 생생하게 경험한다.
- **Estimated time:** 1분
- **Key points:**
  - 오늘: 도구를 준비했다 (설치)
  - 다음: **경험한다** (Education 실습)
  - "블룸의 분류학에 대해 알려줘" — Cocrates가 어떻게 반응할까?
  - 기대하셔도 좋습니다

---

## 시간 배분 요약

| 구간 | 슬라이드 | 시간 | 누적 |
|------|---------|------|------|
| 도입 | 1~2 | 2분 | 2분 |
| 본론 — 구조 이해 | 3~4 | 2분 | 4분 |
| 본론 — Step 1 (macOS/Linux) | 5 | 1분 | 5분 |
| 본론 — Step 1 (Windows) | 6 | 1분 | 6분 |
| 본론 — GUI 환경 구성 | 7 | 1분 | 7분 |
| 본론 — 치트키 + 설치 확인 | 8~9 | 2분 | 9분 |
| 본론 — 첫 대화 | 10 | 1분 | 10분 |
| 결론 — 요약 + 학습 확인 | 11 | 1분 | 11분 |
| 결론 — 다음 편 연결 | 12 | 1분 | 12분 |
