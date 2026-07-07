---
sidebar_position: 3
---
# EP3. Cocrates Harness 설치

![Cocrates Install Procedure](/img/03-installation.png)

---

Ep2에서 우리는 하나의 강력한 선언을 했다.

> "The unexamined code is not worth generating."
> 검토되지 않은 산출물은 생성할 가치가 없다.

원칙은 분명하다. AI가 만든 모든 것은 검토(U→A→E→A)를 거쳐야 하고, 무지를 방치하지 말아야 한다.

그런데, 한 가지 현실적인 질문이 남아 있다.

*"원칙은 알겠다. 그런데 도대체 어떻게 시작하지?"*

오늘 우리는 그 질문에 답한다. 설치는 5분이면 끝난다. 하지만 진짜 시작은 설치 후라는 것, 그 점을 꼭 기억하길 바란다.

---

## 🎭 Cocrates Harness의 정체

먼저 Cocrates Harness의 정체부터 명확히 하고 넘어가자.

Cocrates Harness는 단독으로 실행되는 프로그램이 아니다. **opencode**라는 플랫폼 위에서 동작하는 **플러그인**이다.

비유하자면 이렇다. opencode는 **무대**이고, Cocrates Harness는 그 무대 위에서 공연하는 **배우**다. 무대가 없으면 배우도 설 자리가 없다.

그래서 Cocrates의 설치는 3단계로 이루어진다.

**하나,** opencode를 설치하고.
**둘,** Cocrates 플러그인을 설정하고.
**셋,** 스킬 파일을 준비한다.

이 세 가지만 기억하라. 설치는 생각보다 훨씬 간단하다.

---

## 🖥️ 3가지 UI 선택지

opencode를 사용하는 방법은 세 가지가 있다. 자신의 환경에 맞게 선택하면 된다.

### 1) opencode 터미널 (TUI)
터미널 환경에서 키보드 중심으로 사용한다. CLI에 익숙한 사용자에게 추천.

### 2) opencode 데스크탑 (베타)
설치형 GUI 앱. 아직 베타 버전이라 기능이나 안정성 면에서 불편한 점이 있을 수 있다.

### 3) VS Code 또는 Cursor + OpenChamber 확장 ⭐ (추천)
VS Code나 Cursor에 익숙하다면, Copilot이나 Cursor를 쓰는 것처럼 자연스럽게 Cocrates를 사용할 수 있다. VS Code + OpenChamber도 opencode가 설치되어 있어야 동작한다는 점을 기억하자.

---

## ⚙️ Step 1 — opencode 설치

### macOS / Linux

터미널을 열고 다음 한 줄만 입력하면 된다.

```bash
curl -fsSL https://opencode.ai/install | bash
```

curl이 설치 스크립트를 안전하게 다운로드하고, bash가 그 스크립트를 바로 실행한다. 그러면 opencode 바이너리가 자동으로 설치된다.

설치가 끝나면 확인해보자.

```bash
opencode --version
```

버전이 출력되면 성공이다. 혹시 *"command not found"*가 뜨면 PATH에 `~/.local/bin`을 추가하면 된다.

### Windows

Windows에도 curl은 기본 내장되어 있지만, bash는 기본적으로 없다. 대신 **Chocolatey**라는 Windows용 패키지 매니저를 사용한다.

먼저 PowerShell을 **관리자 권한**으로 실행하고, Chocolatey를 설치한다.

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```

Chocolatey가 설치되었다면, opencode 설치는 단 한 줄이다.

```powershell
choco install opencode
```

그리고 `opencode --version`으로 확인한다.

### VS Code + OpenChamber 확장 설치

opencode CLI 설치는 끝났다. 그런데 터미널 UI가 어색하다면, VS Code + **OpenChamber** 확장을 설치하자.

1. 아직 VS Code가 없다면 [code.visualstudio.com](https://code.visualstudio.com)에서 다운로드해 설치한다.
2. VS Code를 실행하고, 확장 아이콘(Ctrl+Shift+X)을 클릭한 다음 검색창에 **"OpenChamber"**를 입력하고 설치 버튼을 누른다.

OpenChamber를 이용하면, 마치 VS Code의 Copilot이나 Cursor를 쓰는 것처럼 VS Code 안에서 Cocrates를 사용할 수 있다.

![OpenChamber Extension](/img/openchamber-extension.png)

---

## 🔧 Step 2-3 — Plugin 설정 + Skill 파일

opencode를 설치한 후, AI에게 Cocrates Harness 설치를 요청하는 방법이 가장 간단하다.

Cocrates를 실행하고 다음과 같이 요청한다.

> *"https://cocrates.ai/install.md 파일을 근거로 Cocrates Harness를 설치해 줘."*

그럼 AI가 알아서 처리한다. GitHub 저장소에서 스킬 파일을 찾아 다운로드하고, opencode.jsonc에 플러그인 설정을 추가하고, 모든 걸 알아서 처리한다.

이미 설치되어 있다면, 같은 요청을 다시 해도 된다. AI는 [install.md](https://cocrates.ai/install.md)에 따라 **업그레이드**로 처리한다.

- **Plugin:** 설치된 버전과 최신 버전을 비교한다. 새 버전이 있으면 갱신하고, 이전·최신 버전 번호를 알려준다. 이미 최신이면 그대로임을 알려준다.
- **Skill:** 스킬마다 공식 저장소와 로컬 파일을 비교한다. 차이가 없으면 건너뛰고, 차이가 있으면 변경 요약을 보여준 뒤 **공식 버전으로 교체할지, 로컬을 유지할지** 선택을 받는다.

별도의 "업그레이드" 명령을 외울 필요는 없다. 평소처럼 설치를 요청하면 된다.

![Cocrates Harness Installed](/img/cocrates-harness-installed.png)

**하지만. 여기서 끝이 아니다.**

AI가 설치했다고 해서 그냥 넘어가면, 그것은 AI-assisted일 뿐이다. **AI가 설치한 내용을 검토하고 승인하는 것.** 그것이 진짜 AI-native Engineer다.

### 설치 확인

Cocrates Harness는 두 가지로 구성된다. **Plugin**과 **Skill.** 이 두 가지만 확인하면 설치가 제대로 되었는지 알 수 있다.

**Plugin 확인.** `~/.config/opencode/opencode.jsonc` 파일을 열어보자. 그 안에 `"plugin": ["@cocrates/cocrates-harness"]`가 있는가? 있으면 통과.

**Skill 확인.** `~/.config/opencode/skills/` 디렉토리를 열어보자. adr-writing, spec-writing, education 등 여러 스킬 파일들이 있는가? 있으면 통과.

이 두 가지만 확인하면 된다. AI가 알아서 잘 설치했다면, 여러분은 확인하고 승인만 하면 끝.

둘 다 정상이면, opencode를 재시작하고 Cocrates Agent를 선택하라. 활성화 완료다.

---

## 🦉 첫 대화 — 진짜는 지금부터

설치 완료! 이제 모든 준비가 끝났다. 진짜를 경험할 시간이다.

Cocrates에게 첫 질문을 던져보자.

> **"블룸의 분류학에 대해 알려줘."**

그런데, 한 가지 미리 말씀드리자면... Cocrates는 일반 AI와 다를 것이다.

보통 AI는 바로 장문의 설명을 내놓겠지만, Cocrates는 그러지 않는다. 은근슬쩍 오류가 섞인 예시와 함께, 여러분에게 첫 번째 미션을 던질 것이다.

**설치만으로 끝나지 않는다. 진짜는 지금부터다.**

---

## 📌 오늘의 핵심

1. **Cocrates는 opencode 플러그인이다.** opencode라는 플랫폼 위에서 동작하며, Plugin + Skill 두 가지로 구성된다.
2. **권장 설치는 opencode + VS Code + OpenChamber 확장이다.** VS Code와 OpenChamber로 GUI 환경에서 Cocrates를 사용할 수 있다.
3. **설치는 AI에게 요청하고, 설치 후 확인하라.** "install.md를 근거로 설치해 줘" 한 마디면 된다. 이미 설치된 경우에도 같은 요청으로 업그레이드된다. 단, 설치·업그레이드 후 Plugin과 Skill이 제대로 있는지 반드시 확인하라. AI의 동작을 확인하고 검토하는 습관이 필요하다.

---

**스스로에게 질문하자.**

- 지금 opencode가 설치되어 있는가?
- Cocrates Agent가 활성화되어 있는가?
- 첫 질문을 던질 준비가 되었는가?

세 질문에 모두 Yes라면, 여러분은 이미 AI-native Engineer의 첫걸음을 내디딘 것이다.

---

## 🎬 다음 편 예고

오늘 우리는 도구를 준비했다. Cocrates Harness를 설치하고, Plugin과 Skill을 확인했으며, 첫 질문을 던질 준비를 마쳤다.

하지만 원칙을 아는 것과 실천하는 것은 다르다. 도구를 가지는 것과 사용하는 것도 다르다.

다음 에피소드, Ep4에서는 드디어 Cocrates와의 **첫 번째 진짜 대화**를 경험한다.

"블룸의 분류학에 대해 알려줘"라고 물었을 때, Cocrates가 어떻게 반응하는지, Education → Knowledge Capture → Reflection으로 이어지는 3단계 학습 파이프라인을 실제 대화와 함께 생생하게 보여주겠다.

생각 유발 유도형 AI와의 첫 만남. 기대하셔도 좋다.

---

*이 시리즈는 Cocrates Harness 프레임워크를 소개합니다. Cocrates는 소크라테스식 대화로 사용자가 주도권을 잡고 성장하도록 설계된 에이전트 하네스입니다.*
