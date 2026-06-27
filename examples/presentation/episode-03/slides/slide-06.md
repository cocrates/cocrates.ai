# Slide 06 — Step 1: opencode 설치 (Windows)

**Section:** 본론 — 설치
**Script reference:** script-06.md

---

## Layout

```
┌─────────────────────────────────────────────────┐
│  Step 1 — opencode 설치 (Windows)                │
│                                                   │
│  Step A: Chocolatey 설치                          │
│  ┌─────────────────────────────────────┐         │
│  │  PS C:\> Set-ExecutionPolicy ...    │         │
│  │  PS C:\> iex ((New-Object ...))     │         │
│  └─────────────────────────────────────┘         │
│                                                   │
│  Step B: opencode 설치                            │
│  ┌─────────────────────────────────────┐         │
│  │  PS C:\> choco install opencode     │         │
│  │  PS C:\> opencode --version         │         │
│  └─────────────────────────────────────┘         │
│                                                   │
│  ⚠️ PowerShell을 관리자 권한으로 실행하세요         │
│  💡 choco = Windows의 앱스토어 같은 패키지 매니저   │
└─────────────────────────────────────────────────┘
```

## Visual Elements

| Element | Description | Source |
|---------|-------------|--------|
| 배경 | 다크 템플릿 | |
| 제목 | "Step 1 — opencode 설치 (Windows)" — 28pt | |
| PowerShell 창 | 블루 테마 PowerShell 창 | |
| Step A/B | 구분선으로 분리 (A: Chocolatey, B: opencode) | |
| 경고 택 | ⚠️ 관리자 권한 실행 강조 | |
| 설명 택 | 💡 choco 설명 | |

## Audio / Narration

- **Tone:** 튜토리얼, 친절하고 인내심 있게
- **Key phrases:**
  - "Windows 사용자는 과정이 조금 다릅니다"
  - "PowerShell을 관리자 권한으로 실행하세요"
  - "choco install opencode"

## Transitions

- **Enter:** Step A → Step B 순차적으로 나타남
- **Exit:** → Slide 07

## Notes

- Windows에서 `curl | bash`가 동작하지 않는 이유 (bash 없음)를 명시.
- Chocolatey 설치 명령어가 길지만, 복붙만 하면 되는 점 강조.
