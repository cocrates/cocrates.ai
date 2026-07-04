# opencode 설치 가이드

> **목적:** AI 에이전트 플랫폼인 opencode를 설치하는 모든 방법을 OS별로 정리한다.
> 초보자도 따라 할 수 있도록 패키지 매니저 설치 과정까지 포함한다.
>
> 참고: opencode 공식 문서 — https://opencode.ai/docs

---

## 1. opencode 설치 방식 개요

opencode를 설치하는 방법은 크게 3가지 경로로 나뉜다.

| 경로 | 방법 | 난이도 | 추천 대상 |
|------|------|--------|----------|
| **Install script** | `curl -fsSL https://opencode.ai/install \| bash` | 중급 | 터미널 사용에 익숙한 개발자 |
| **npm / bun / pnpm / yarn** | `npm install -g opencode-ai` | 중급 | Node.js 환경이 이미 있는 개발자 |
| **패키지 매니저** | choco (Windows), brew (macOS/Linux), pacman (Arch), scoop (Windows) | 중급 | 해당 패키지 매니저가 설치된 환경 |

---

## 2. 공통 사전 준비

모든 설치 방식에 공통으로 필요한 것:

1. **터미널(명령 프롬프트)** — Windows는 PowerShell, macOS는 Terminal.app, Linux는 기본 터미널
2. **인터넷 연결** — 설치 스크립트와 패키지를 다운로드
3. **LLM API 키** — opencode 사용을 위한 AI 모델 제공자 키 (기본 모델 사용 시 불필요)

---

## 3. Windows — 설치 가이드

### 3.1 Chocolatey로 설치 (권장)

Chocolatey는 Windows용 패키지 매니저로, opencode를 명령어 한 줄로 설치할 수 있게 해준다.
단, Chocolatey 자체가 기본 설치되어 있지 않으므로 먼저 설치해야 한다.

#### Step 1: Chocolatey 설치

PowerShell을 **관리자 권한**으로 실행한 후, 다음 명령어를 입력한다.

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
```

> **설명:** 이 명령어는 실행 정책을 임시로 변경하고, Chocolatey 공식 설치 스크립트를 다운로드하여 실행한다.

설치 확인:
```powershell
choco --version
```

#### Step 2: opencode 설치

```powershell
choco install opencode
```

#### Step 3: 설치 확인

```powershell
opencode --version
```

### 3.2 Scoop으로 설치

Scoop도 Chocolatey와 유사한 Windows 패키지 매니저다.

```powershell
scoop install opencode
```

### 3.3 npm으로 설치

Node.js가 이미 설치된 환경이라면 npm으로 설치할 수 있다.

```powershell
npm install -g opencode-ai
```

### 3.4 WSL (Windows Subsystem for Linux) 활용

> opencode 공식 문서는 Windows에서는 WSL 사용을 **강력 권장**한다.
> WSL이 제공하는 Linux 환경에서 opencode의 모든 기능이 완전하게 동작한다.

```powershell
# 1. WSL 설치
wsl --install

# 2. WSL 터미널 실행 후 Linux 방식으로 설치
# WSL 내에서:
curl -fsSL https://opencode.ai/install | bash
```

---

## 4. macOS / Linux — 설치 가이드 (공통)

> macOS와 Linux는 `curl`과 `bash`를 기본적으로 사용할 수 있어 **동일한 install script로 설치**할 수 있다.
> - **macOS:** Terminal.app에 `curl`과 `bash`가 기본 내장되어 있다.
> - **Linux:** 대부분의 배포판에 `curl`이 기본 포함되어 있다. (없으면 `sudo apt install curl` 또는 `sudo dnf install curl`)

### 4.1 Install script로 설치 (권장, 최단 경로)

터미널을 열고 다음 한 줄을 입력한다.

```bash
curl -fsSL https://opencode.ai/install | bash
```

> **설명:**
> - `curl -fsSL` — 설치 스크립트를 안전하게 다운로드
> - `| bash` — 다운로드한 스크립트를 즉시 실행
> - opencode 바이너리가 `~/.local/bin/opencode`에 설치된다

설치 확인:
```bash
opencode --version
```

> **팁:** `command not found`가 뜨면 `~/.local/bin`이 PATH에 없는 것.
> macOS는 `export PATH="$HOME/.local/bin:$PATH"`를 `~/.zshrc`에,
> Linux는 `~/.bashrc`에 추가하면 된다.

### 4.2 Homebrew로 설치 (대안)

Homebrew가 이미 설치된 환경에서만 사용 가능하다.
(Homebrew는 기본 설치되어 있지 않으며, 설치하려면 위 4.1과 동일하게 `curl | bash` 방식이 필요하다.)

```bash
brew install anomalyco/tap/opencode
```

> `brew install opencode`로도 설치 가능하지만, 공식 tap(`anomalyco/tap/opencode`)이 더 최신 릴리즈를 제공한다.

### 4.3 npm으로 설치 (대안)

Node.js가 이미 설치된 환경이라면 npm으로도 설치할 수 있다.

```bash
npm install -g opencode-ai
```

### 4.4 Arch Linux (Linux 전용)

```bash
sudo pacman -S opencode        # Stable 버전
paru -S opencode-bin           # AUR 최신 버전
```

---

## 5. VS Code + OpenChamber로 GUI 사용하기

opencode는 기본적으로 TUI(Text-based User Interface, 터미널 환경)로 동작한다.
터미널 사용이 익숙하지 않은 사용자는 VS Code + OpenChamber 확장을 통해 GUI 환경에서 사용할 수 있다.

### Step 1: opencode CLI 설치

> VS Code + OpenChamber를 사용하더라도 opencode CLI는 별도로 설치해야 한다.
> 앞서 설명한 OS별 설치 방법 중 하나를 선택하여 설치한다.

**초보자 추천 경로:**
- **Windows:** Chocolatey 설치 → `choco install opencode`
- **macOS:** `curl -fsSL https://opencode.ai/install | bash`
- **Linux:** `curl -fsSL https://opencode.ai/install | bash`

### Step 2: VS Code 설치

1. https://code.visualstudio.com 에 접속한다.
2. 운영체제에 맞는 설치 파일을 다운로드한다.
3. 설치 파일을 실행하고 화면 안내에 따라 설치한다 (Next만 누르면 됨).

### Step 3: OpenChamber 확장 설치

1. VS Code를 실행한다.
2. 왼쪽 사이드바에서 **확장 아이콘**(Extensions, `Ctrl+Shift+X`)을 클릭한다.
3. 검색창에 **"OpenChamber"** 를 입력한다.
4. 검색 결과에서 OpenChamber 확장을 찾아 **설치(Install)** 버튼을 클릭한다.
5. 설치 완료 후 VS Code를 재시작한다.

---

## 6. 설치 방식 비교표

| 방법 | Windows | macOS | Linux | CLI 필요 | 사전 설치 필요 |
|------|---------|-------|-------|----------|---------------|
| Install script | △ (WSL 필요) | ✅ | ✅ | 예 | curl |
| npm | ✅ | ✅ | ✅ | 예 | Node.js + npm |
| Chocolatey | ✅ | — | — | 예 | Chocolatey 설치 필요 |
| Scoop | ✅ | — | — | 예 | Scoop 설치 필요 |
| Homebrew | — | ✅ | ✅ | 예 | Homebrew 설치 필요 |
| pacman | — | — | ✅ (Arch) | 예 | — |
| VS Code + OpenChamber | ✅ | ✅ | ✅ | **예** (opencode CLI 별도 설치) | VS Code |

---

*참고: opencode 공식 문서 기준 (2026-06-25 확인)*
