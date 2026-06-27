# Script: Slide 06 — Step 1 — opencode 설치 (Windows)

**Time:** ~1분

**[SLIDE 06: PowerShell 창 이미지. Chocolatey 로고와 opencode 로고. 두 줄의 명령어를 구분해서 표시]**

Windows 사용자분들은 과정이 조금 다릅니다.

Windows에도 curl은 기본 내장되어 있지만,
안타깝게도 bash는 기본적으로 없습니다.
그래서 opencode 설치 스크립트를 그대로 쓸 수는 없어요.

대신 **Chocolatey**라는 Windows용 패키지 매니저를 사용합니다.

(Step A를 가리키며)

먼저 PowerShell을 **관리자 권한**으로 실행하세요.
시작 버튼에 마우스 오른쪽 버튼, "Windows PowerShell (관리자)"를 선택하면 됩니다.

그리고 다음 명령어를 입력합니다.

(명령어를 읽으며)

**Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))**

이 한 줄이 Chocolatey를 설치합니다.
choco는 Windows의 앱스토어 같은 존재라고 생각하세요.

(Step B를 가리키며)

Chocolatey가 설치되었다면,
이제 진짜 opencode 설치는 단 한 줄입니다.

**choco install opencode**

그리고 `opencode --version`으로 확인.

(강조하며)

choco 자체를 설치하는 게 더 복잡해 보일 수 있지만,
한 번만 설치해두면 앞으로 모든 프로그램을
이렇게 간단하게 설치할 수 있습니다.

→ [SLIDE 07: VS Code + OpenChamber 확장 설치]
