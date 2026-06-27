# Script: Slide 09 — 설치 확인 — Cocrates = Plugin + Skill

**Time:** ~1분

**[SLIDE 09: 2열 레이아웃. 좌측 "Plugin" (설정 파일 아이콘 + 체크리스트), 우측 "Skill" (폴더 아이콘 + 체크리스트). 상단에 "Cocrates = Plugin + Skill"]**

자, AI가 설치를 끝냈습니다.
이제 우리가 검토할 차례입니다.

Cocrates Harness는 두 가지로 구성됩니다.
**Plugin**과 **Skill.**

이 두 가지만 확인하면 설치가 제대로 되었는지 알 수 있습니다.
U→A→E→A로 검토해보세요.

(좌측을 가리키며)

**첫째, Plugin 확인.**
파일을 열어보세요: `~/.config/opencode/opencode.jsonc`
그 안에 `"plugin": ["@cocrates/cocrates-harness"]`가 있는가?
있으면 통과.

(우측을 가리키며)

**둘째, Skill 확인.**
디렉토리를 열어보세요: `~/.config/opencode/skills/`
adr-writing, spec-writing, education 등
여러 스킬 파일들이 있는가?
있으면 통과.

(두 손을 모으며)

이 두 가지만 확인하면 됩니다.
AI가 알아서 잘 설치했다면, 여러분은 확인하고 승인만 하면 끝.

(단호하게)

이것이 U→A→E→A의 힘입니다.
AI가 일을 수행하고, 인간이 검토하고 승인합니다.

둘 다 정상이면,
opencode를 재시작하고 Cocrates Agent를 선택하세요.
활성화 완료입니다.

→ [SLIDE 10: 설치 완료! 첫 대화]
