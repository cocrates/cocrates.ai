# Slide Plan — Episode 03

Total slides: 12
Estimated pace: ~1분/슬라이드 (총 ~12분)

---

## Slide 1: 타이틀 페이지
- **Section:** 도입
- **Purpose/goal:** Ep2의 원칙 선언을 복습하고, 이제 그 원칙을 실천할 도구를 설치할 차례임을 전달한다.
- **Content:**
  - Ep2 복습: "The unexamined code is not worth generating"
  - "원칙은 알겠다. 그런데 어떻게 시작하지?"
  - 오늘의 답변: **Cocrates Harness 설치 — AI-native Engineer의 첫걸음**
  - 핵심 메시지: "설치는 5분이면 끝나지만, 진짜 시작은 설치 후다"
  - 시각: Ep2의 선언문 이미지를 작게 리마인더로 배치, 아래에 설치 아이콘(⚙️)
- **Type:** 타이틀 / 맥락 연결

---

## Slide 2: 학습 목표
- **Section:** 도입
- **Purpose/goal:** 이 에피소드의 3가지 학습 목표를 체크리스트로 제시한다.
- **Content:**
  - 학습 목표 3가지:
    - opencode(플랫폼)와 Cocrates(플러그인)의 관계를 이해한다
    - 3단계 설치 과정(opencode → plugin → skills)을 설명할 수 있다
    - 자신의 환경에 맞는 방법으로 Cocrates Harness를 직접 설치한다
  - 시각: 체크리스트 형식, 각 항목 빈 체크박스
- **Type:** 학습 목표 / 체크리스트

---

## Slide 3: Cocrates Harness의 정체
- **Section:** 본론 — 구조 이해
- **Purpose/goal:** Cocrates Harness가 단독 실행 프로그램이 아니라 opencode 플랫폼의 플러그인임을 명확히 전달한다.
- **Content:**
  - **opencode (플랫폼):** AI 에이전트를 실행하고 관리하는 무대
  - **Cocrates Harness (플러그인):** 그 무대 위에서 특별한 능력을 발휘하는 배우
  - 비유: "극장(플랫폼)과 배우(플러그인)"
  - 설치 = 3단계: opencode 설치 → plugin 설정 → skill 파일 복사
  - 시각: 2층 구조 다이어그램 — 아래층 "opencode Platform", 위층 "Cocrates Harness Plugin". 또는 극장과 배우 아이콘.
- **Type:** 개념 설명 / 비유

---

## Slide 4: 3가지 UI 선택지
- **Section:** 본론 — 구조 이해
- **Purpose/goal:** opencode를 사용하는 3가지 UI 방식을 소개하고, 초보자에게 VS Code + OpenChamber를 권장한다.
- **Content:**
  - **① opencode 터미널 (TUI):** 터미널 기반, 키보드 중심. 힙스터 취향.
  - **② opencode 데스크탑:** 설치형 GUI 앱. **아직 베타** — 기능/안정성 불안정.
  - **③ VS Code + OpenChamber 확장:** 기존 VS Code 환경에 GUI 추가. **초보자 추천.**
  - 단, 어떤 방식을 선택하든 **opencode CLI 설치는 공통**
  - 시각: 3개의 카드 (각각 아이콘 + 간단한 특징). 3번째 카드(OpenChamber)를 강조 표시.
- **Type:** 옵션 소개 / 추천

---

## Slide 5: Step 1 — opencode 설치 (macOS / Linux)
- **Section:** 본론 — 설치
- **Purpose/goal:** macOS와 Linux 사용자에게 가장 간단한 설치 방법(install script)을 안내한다.
- **Content:**
  - macOS는 Terminal.app에 `curl`과 `bash`가 기본 내장
  - Linux도 대부분 `curl` 기본 포함 (없으면 `sudo apt install curl`)
  - **명령어 (따라 치기 쉽게 크게 표시):**
    ```
    curl -fsSL https://opencode.ai/install | bash
    ```
  - 설치 확인: `opencode --version`
  - "Homebrew를 먼저 설치할 필요 없습니다. curl과 bash로 바로 됩니다."
  - 시각: 터미널 창 이미지 또는 명령어를 큰 폰트로 중앙 배치
- **Type:** 튜토리얼 / 명령어 안내

---

## Slide 6: Step 1 — opencode 설치 (Windows)
- **Section:** 본론 — 설치
- **Purpose/goal:** Windows 사용자에게 Chocolatey를 통한 설치 방법을 안내한다. Chocolatey 설치 과정까지 포함한다.
- **Content:**
  - **Step A: Chocolatey 설치** (PowerShell을 관리자 권한으로 실행)
    ```
    Set-ExecutionPolicy Bypass -Scope Process -Force; [Net.ServicePointManager]::SecurityProtocol = [Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
    ```
  - **Step B: opencode 설치**
    ```
    choco install opencode
    ```
  - 설치 확인: `opencode --version`
  - "choco는 Windows의 앱스토어 같은 패키지 매니저입니다"
  - 대안: WSL + install script (공식 권장)
  - 시각: PowerShell 창 이미지. Chocolatey 로고와 opencode 로고.
- **Type:** 튜토리얼 / 명령어 안내

---

## Slide 7: VS Code + OpenChamber 확장 설치
- **Section:** 본론 — 설치
- **Purpose/goal:** TUI가 익숙하지 않은 사용자를 위해 VS Code + OpenChamber 확장으로 GUI 환경을 구성하는 방법을 안내한다.
- **Content:**
  - **VS Code 설치:** https://code.visualstudio.com 에서 다운로드 → 설치 (Next만 누르면 끝)
  - **OpenChamber 확장 설치:** VS Code 실행 → 확장(Ctrl+Shift+X) → "OpenChamber" 검색 → 설치
  - "마치 Copilot이나 Cursor처럼 VS Code에서 Cocrates를 사용할 수 있습니다"
  - 시각: VS Code 화면 캡처(OpenChamber 확장이 활성화된 모습). 또는 VS Code 로고 + OpenChamber 로고 + opencode 로고 연결 다이어그램.
- **Type:** 튜토리얼 / 설치 안내

---

## Slide 8: AI에게 설치 시키기
- **Section:** 본론 — 설치 + 검증
- **Purpose/goal:** 직접 설치 명령어를 외울 필요 없이, AI에게 설치를 요청하는 방법을 알려준다. 설치는 AI가 하고, 검토는 인간이 한다는 원칙을 강조한다.
- **Content:**
  - opencode만 설치되어 있다면, 한 마디면 끝
  - 👤 **유저:** *"[https://cocrates.ai/install.md](https://cocrates.ai/install.md) 파일을 근거로 Cocrates Harness를 설치해 줘."*
  - AI가 plugin 설정부터 skill 파일 다운로드까지 알아서 처리
  - **하지만 "설치조차 검토하고 승인하세요"** — AI가 설치한 내용을 확인(U→A→E→A)하는 것이 진짜 AI-native Engineer
  - 다음 슬라이드에서 확인 방법 안내
  - 시각: 말풍선 대화 형식. 유저(👤) → AI 메시지 강조. 하단에 "🔍 검토는 다음 슬라이드에서" 안내.
- **Type:** 사례 연구 / 팁

---

## Slide 9: 설치 확인 — Cocrates = Plugin + Skill
- **Section:** 본론 — 설치 + 검증
- **Purpose/goal:** AI가 설치한 내용을 검토하는 방법을 안내한다. Cocrates가 Plugin + Skill로 구성되어 있음을 명확히 하고, 이 두 가지만 확인하면 된다는 점을 전달한다.
- **Content:**
  - **Cocrates = Plugin + Skill** (두 가지로 구성, 하나만 있어도 불완전)
  - **🔍 Plugin 확인:**
    - 파일: `~/.config/opencode/opencode.jsonc`
    - 확인할 내용: `"plugin": ["@cocrates/cocrates-harness"]`가 있는가?
  - **🔍 Skill 확인:**
    - 디렉토리: `~/.config/opencode/skills/`
    - 확인할 내용: skill 파일들(adr-writing, spec-writing, education 등)이 있는가?
  - "AI가 설치했다면, 이 두 가지만 검토(U→A→E→A)하면 됩니다"
  - 둘 다 정상 → opencode 재시작 → Cocrates Agent 선택 → 활성화 완료
  - 시각: 2열 레이아웃 — 좌측 "Plugin" (설정 파일 아이콘 + 체크리스트), 우측 "Skill" (폴더 아이콘 + 체크리스트). 상단에 "Cocrates = Plugin + Skill".
- **Type:** 프로세스 설명 / 체크리스트

---

## Slide 10: 설치 완료! 첫 대화
- **Section:** 본론 — 첫 대화
- **Purpose/goal:** 설치가 끝났음을 알리고, 첫 질문을 던지게 한다. Cocrates가 일반 AI와 다를 것이라는 기대감을 형성한다.
- **Content:**
  - 설치 완료 체크리스트:
    - ✅ opencode 설치 완료
    - ✅ Plugin 설정 확인
    - ✅ Skill 파일 확인
    - ✅ Cocrates Agent 활성화
  - 이제 첫 질문을 던져보자:
    - 💬 **"블룸의 분류학에 대해 알려줘"**
  - "일반 AI와 다를 겁니다 — Cocrates는 곧바로 답을 주지 않아요"
  - **"설치만으로 끝나지 않습니다. 진짜는 지금부터입니다."**
  - 시각: 큰 체크리스트 4개가 모두 ✅로 채워지는 애니메이션. 아래에 말풍선으로 첫 질문 표시.
- **Type:** 전환 / 기대감 형성

---

## Slide 11: 핵심 요약 + 학습 확인
- **Section:** 결론 — 요약
- **Purpose/goal:** 오늘의 3가지 핵심 포인트를 압축하고, 시청자가 직접 설치했는지 확인하는 질문을 던진다.
- **Content:**
  - **3가지 요약:**
    1. **Cocrates는 opencode 플러그인이다.** — Plugin + Skill로 구성
    2. **opencode의 권장 설치는 opencode + VS Code + OpenChamber 확장이다.** — GUI 환경
    3. **Cocrates의 설치는 AI에게 요청한다.** — 단, 설치 후 Plugin + Skill은 반드시 검토
  - **학습 확인 질문:**
    - "지금 opencode가 설치되어 있는가?"
    - "Cocrates Agent가 활성화되어 있는가?"
    - "첫 질문을 던질 준비가 되었는가?"
  - 시각: 3개의 숫자 아이콘 + 간결한 텍스트. 아래에 3개의 질문 마크 아이콘.
- **Type:** 요약 / 학습 확인

---

## Slide 12: 다음 에피소드 연결
- **Section:** 결론 — 연결
- **Purpose/goal:** 설치가 끝났으니 이제 진짜 경험을 시작할 차례임을 전달하고, Ep4에 대한 기대감을 형성한다.
- **Content:**
  - 오늘: **도구를 준비했다** (설치)
  - 다음: **경험한다** (Education 실습)
  - "블룸의 분류학에 대해 알려줘" — Cocrates가 어떻게 반응할까?
  - Cocrates는 곧바로 답을 주지 않는다. **은근슬쩍 오류가 섞인 예시와 함께 첫 번째 미션**을 던질 것이다.
  - "생각 유발 유도형 AI와의 첫 만남, 다음 편에서 만나요! 🦉"
  - 시각: 시리즈 맵 (Ep3 위치 표시, Ep4 강조). "오늘(설치) → 다음(실습)" 화살표.
- **Type:** 연결 / 기대감 형성
