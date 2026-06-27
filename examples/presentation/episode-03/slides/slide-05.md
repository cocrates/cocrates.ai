# Slide 05 — Step 1: opencode 설치 (macOS / Linux)

**Section:** 본론 — 설치
**Script reference:** script-05.md

---

## Layout

```
┌─────────────────────────────────────────────────┐
│  Step 1 — opencode 설치                          │
│              (macOS / Linux)                     │
│                                                   │
│  ┌─────────────────────────────────────┐         │
│  │  💻 터미널 창                         │         │
│  │                                     │         │
│  │  $ curl -fsSL                       │         │
│  │    https://opencode.ai/install      │         │
│  │    | bash                            │         │
│  │                                     │         │
│  │  $ opencode --version               │         │
│  │  opencode x.y.z                     │         │
│  └─────────────────────────────────────┘         │
│                                                   │
│  ✅ macOS: curl/bash 기본 내장 (추가 설치 불필요)   │
│  ✅ Linux: 대부분 curl 기본 탑재                   │
│  💡 PATH: ~/.local/bin 자동 등록 (확인 필요 시 수동) │
└─────────────────────────────────────────────────┘
```

## Visual Elements

| Element | Description | Source |
|---------|-------------|--------|
| 배경 | 다크 템플릿 | |
| 제목 | "Step 1 — opencode 설치 (macOS / Linux)" — 28pt | |
| 터미널 창 | 다크 테마 터미널 (초록색 프롬프트) — 명령어 표시 | |
| 확인 메시지 | "opencode x.y.z" — 버전 출력 모사 | |
| 팁 박스 | 하단 3개의 요약 체크 | |

## Audio / Narration

- **Tone:** 튜토리얼, 자신감 있고 명확하게
- **Key phrases:**
  - "macOS를 쓰고 계신가요? Terminal.app만 여세요"
  - "curl -fsSL https://opencode.ai/install | bash"
  - "Homebrew 같은 걸 먼저 설치할 필요 없습니다"

## Transitions

- **Enter:** 터미널 창이 타이핑되는 애니메이션
- **Exit:** → Slide 06

## Notes

- macOS/Linux 사용자용. Windows 사용자는 이 슬라이드를 건너뛰거나 참고만.
- "curl | bash"가 macOS에서 작동하는 이유(둘 다 기본 내장)를 설명.
