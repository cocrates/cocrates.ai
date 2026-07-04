# Slide 09: ⚡ Generation — 오직 Spec만 보고 달린다

## Key Message
초기 프롬프트는 잊고, 오직 합의된 Spec만으로 코드를 생성한다. Spec의 품질이 결과물의 품질을 결정한다.

## Content
- **Headlines / text:**
  - **모드:** ⚡ `spec-driven-generation` 활성화
  - **원칙:** 🚫 초기 프롬프트("jsondb 만들어줘")는 잊는다. ✅ **오직 합의된 Spec만** 유일한 법전
  - **결과물 (Go 언어):**
    - 📦 코어 라이브러리 Go 파일 7개
    - 🌐 REST API 서버
    - ⌨️ CLI 도구
    - 🧪 단위 테스트 + 통합 테스트
    - ✅ **컴파일 오류 0개** (go build, go vet 패스)
- **핵심 메시지:** "Spec을 잘 설계하면 AI가 코드를 잘 만든다. **Spec의 품질이 결과물의 품질을 결정한다.** "

## Visual Elements
- 맨 왼쪽에 Spec 문서 → 중앙에 번개 아이콘(Generation) → 오른쪽에 결과물 아이콘들
- 결과물 아이콘(파일 7개, 서버, CLI, 테스트)이 한 번에 나타나는 효과
- "컴파일 오류 0개"는 큰 녹색 체크마크로 강조

## Layout Notes
- 좌→우 흐름: Spec → Generation → Result
- 결과물은 아이콘 그리드로 표시 (2×2 또는 1×4)
- 하단에 핵심 메시지 강조 표시

## Reference
- Script: `scripts/script-09.md`
- Slide plan: `slides.md` → Slide 09
