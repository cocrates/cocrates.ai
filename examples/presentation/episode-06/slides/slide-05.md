# Slide 05: 🧱 Frame — 산출물의 뼈대와 작업 파일 트리

## Key Message
Snowflake 점진적 확장의 핵심: 작게 쪼개고, 파일에 저장하고, 검토하고, 승인받고, 다음 단계로.

## Content
- **문서 레이아웃:** 메타정보 → 개요 → 본문(섹션/서브섹션) → 결론 → 부록
- **Snowflake 파일 트리:**
  ```
  docs/{slug}/
  ├── outline.md              # 1단계
  ├── sections.md             # 2단계
  ├── sections/               # 3단계
  └── {slug}.md               # 4단계 (최종 통합본)
  ```
- **Cocrates의 제안:** "한 번에 통으로 쓰는 것보다 섹션 단위로 쪼개어 검토하며 작성하는 게 안전하겠죠?"
- **포인트:** 채팅창에서 모든 걸 주고받지 않는다. 파일이 작업의 단위다.

## Visual Elements
- Left: 문서 레이아웃 (horizontal flow with arrows)
- Right: File tree structure (indented hierarchy)
- Bottom: Snowflake ❄️ icon with 4 principles

## Layout Notes
- Two-column layout (left: layout, right: file tree)
- File tree should use monospace font
- Cocrates quote in a speech bubble

## Reference
- Script: `scripts/script-05.md`
- Slide plan: `slides.md` → Slide 05
