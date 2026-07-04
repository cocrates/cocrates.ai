# Q&A Plan — Episode 03

## Anticipated Audience Questions

### Q1: "curl | bash는 보안에 안전한가요?"

**Answer:**
좋은 질문입니다. `curl | bash` 패턴은 실제로 보안 우려가 제기되곤 합니다. opencode의 경우:

1. **HTTPS 사용:** curl이 HTTPS로 다운로드하므로 전송 중 변조 방지
2. **서명된 릴리스:** opencode는 서명된 바이너리로 배포
3. **대안:** 설치 스크립트를 먼저 다운로드해서 검토 후 실행 가능
   ```bash
   curl -fsSL https://opencode.ai/install -o install.sh
   # install.sh 내용을 검토
   less install.sh
   # 검토 후 실행
   bash install.sh
   ```

**관련 슬라이드:** 05, 06

---

### Q2: "Homebrew를 이미 쓰고 있는데, Homebrew로 opencode를 설치할 수 없나요?"

**Answer:**
공식 설치 스크립트(`curl | bash`)가 가장 권장되는 방법입니다. Homebrew를 통해 설치할 수 있는 경우도 있지만, 공식 문서에서는 설치 스크립트를 첫 번째 방법으로 권장합니다.

중요한 점은 macOS에서 Homebrew를 설치하기 위해 **먼저 Xcode Command Line Tools를 설치**해야 하는데, 그 과정이 `curl | bash`보다 더 복잡합니다. curl과 bash는 macOS에 **이미** 설치되어 있으니, 추가 설치 없이 바로 사용 가능합니다.

**관련 슬라이드:** 05

---

### Q3: "Windows에서 WSL을 쓰고 있는데, Chocolatey 대신 WSL을 통해 설치할 수 있나요?"

**Answer:**
네, 가능합니다! WSL(Windows Subsystem for Linux)을 사용 중이라면 WSL 내에서 Linux 설치 스크립트를 그대로 사용할 수 있습니다.

```bash
# WSL 터미널 (Ubuntu 등)에서
curl -fsSL https://opencode.ai/install | bash
```

다만 이 경우 opencode가 WSL 환경에서 실행된다는 점을 알아두세요. Windows 네이티브 환경이 필요하다면 Chocolatey 경로를 사용하세요.

**관련 슬라이드:** 06

---

### Q4: "VS Code 대신 IntelliJ나 다른 IDE를 쓰는데, OpenChamber 확장이 있나요?"

**Answer:**
현재 OpenChamber 확장은 VS Code 전용입니다. 대안으로는:

1. **opencode TUI 사용:** VS Code 통합 터미널 대신 일반 터미널에서 `opencode` 실행
2. **opencode Desktop (BETA):** 독립 실행형 GUI 앱 사용

어떤 방법을 선택하든 Cocrates의 핵심 기능은 동일하게 사용할 수 있습니다.

**관련 슬라이드:** 04, 07

---

### Q5: "설치했는데 'opencode: command not found'가 떠요"

**Answer:**
일반적인 문제입니다. 설치 스크립트가 PATH에 `~/.local/bin`을 추가했는지 확인하세요.

```bash
# PATH 확인
echo $PATH
# ~/.local/bin이 없다면 추가
export PATH="$HOME/.local/bin:$PATH"
# shell 설정 파일에 영구 등록 (.bashrc / .zshrc)
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

설치가 제대로 되었는지 다시 확인: `opencode --version`

**관련 슬라이드:** 05

---

### Q6: "Plugin 설정 파일(opencode.jsonc)이 어디 있는지 모르겠어요"

**Answer:**

운영체제별 경로:
- **macOS/Linux:** `~/.config/opencode/opencode.jsonc`
- **Windows:** `%APPDATA%\opencode\opencode.jsonc` (또는 `C:\Users\<username>\.config\opencode\opencode.jsonc`)

또는 AI에게 물어보세요: *"opencode.jsonc 파일이 어디 있어?"*

**관련 슬라이드:** 09

---

### Q7: "Cocrates Agent를 선택하라고 했는데, 그게 뭔가요?"

**Answer:**
opencode를 실행하면 여러 Agent(에이전트) 중에서 선택할 수 있습니다. Cocrates Agent는 Cocrates Harness 플러그인이 활성화된 상태로 opencode가 실행될 때 나타나는 옵션입니다.

- 기본 Agent: opencode의 기본 대화 모드
- **Cocrates Agent:** 플러그인이 적용된 모드 (Socratic 대화, 설계-검토-승인 흐름, 스킬 활용)

opencode 실행 후 나타나는 선택 화면에서 방향키로 "Cocrates"를 선택하거나, 설정 파일에서 기본 Agent로 지정할 수 있습니다.

**관련 슬라이드:** 09, 10

---

### Q8: "설치가 잘 되었는지 확인하는 가장 빠른 방법은?"

**Answer:**
3단계 확인:

1. **터미널에서:** `opencode --version` → 버전 출력 확인
2. **플러그인:** `opencode.jsonc` 파일에 `"cocrates-harness"` 포함 확인
3. **스킬:** `~/.config/opencode/skills/` 디렉토리에 스킬 파일들 존재 확인

이 세 가지만 확인하면 됩니다. 보통 30초도 안 걸립니다.

**관련 슬라이드:** 09

---

## Presenter Preparation Notes

### 핵심 포인트 (반드시 전달)
- Cocrates = Plugin + Skill (두 가지로 구성)
- opencode 설치가 선행되어야 함
- 설치 후 검토(U→A→E→A)가 진짜 AI-native Engineer

### 타이밍 주의
- Slide 05-07 (설치): OS별 사용자만 집중 (다른 OS 사용자는 잠시 대기)
- Slide 08 (치트키): "검토는 다음 슬라이드"에서 실제 검토가 나오므로 연결감 유지

### 데모 준비 (라이브 시)
- 사전에 opencode + Cocrates 설치 완료된 환경 준비
- "블룸의 분류학" 질문에 대한 Cocrates의 실제 반응 화면 캡처 준비
- Chocolatey 설치가 느릴 수 있으니 대비

### Q&A 타임 확보
- 설치 관련 질문이 가장 많이 나올 것으로 예상 (약 5분)
- OS별 차이에 대한 질문 대비 (약 3분)
- "다음 편은 언제?" — Ep4 일정 공유 가능

## Follow-up Resources

- 공식 설치 문서: https://opencode.ai/install
- Cocrates GitHub: https://github.com/cocrates/cocrates-harness
- VS Code + OpenChamber: VS Code 확장 마켓플레이스에서 "OpenChamber" 검색
- 문제 해결: opencode Discord / GitHub Issues
