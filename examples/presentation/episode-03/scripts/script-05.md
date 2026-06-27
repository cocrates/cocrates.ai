# Script: Slide 05 — Step 1 — opencode 설치 (macOS / Linux)

**Time:** ~1분

**[SLIDE 05: 터미널 창 이미지. 중앙에 큰 폰트로 설치 명령어: "curl -fsSL https://opencode.ai/install | bash". 아래에 "설치 확인: opencode --version"]**

자, 이제 진짜 설치를 시작해볼까요?

먼저 macOS와 Linux 사용자부터.

(자신 있게)

macOS를 쓰고 계신가요?
좋습니다. Terminal.app만 여세요.
curl과 bash는 macOS에 기본 내장되어 있습니다.
따로 설치할 게 아무것도 없어요.

Linux도 마찬가지입니다.
대부분의 배포판에 curl이 기본으로 들어 있습니다.

(명령어를 가리키며)

터미널에 다음 한 줄만 입력하면 됩니다.

**curl -fsSL https://opencode.ai/install | bash**

(설명하며)

curl이 설치 스크립트를 안전하게 다운로드하고,
bash가 그 스크립트를 바로 실행합니다.
그러면 opencode 바이너리가 자동으로 설치됩니다.

설치가 끝나면 확인해보세요.

**opencode --version**

버전이 출력되면 성공입니다.

(잠시 멈추며)

혹시 "command not found"가 뜨면
PATH에 `~/.local/bin`을 추가하면 됩니다.
설치 스크립트가 이 과정을 자동으로 처리해주긴 하지만,
가끔 놓치는 경우가 있어요.

자, macOS와 Linux 사용자는 이것으로 설치 끝!
Homebrew 같은 걸 먼저 설치할 필요 없습니다.

→ [SLIDE 06: Step 1 — opencode 설치 (Windows)]
