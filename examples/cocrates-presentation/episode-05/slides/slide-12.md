# Slide 12: 🔄 구조는 고정되지 않는다 — Living Cycle

## Key Message
구조 기반 개발은 Waterfall이 아니다. 검증/요구사항 변경 시 구조 설계(ADR)로 돌아가는 순환 구조다. ADR과 Spec이 함께 진화한다.

## Content
- **Headlines / text:**
  - **오해: Waterfall** — 구조 설계 → 구현 → 검증 (끝) ❌
  - **현실: Living Cycle** — 구조가 개발을 이끌지만, **구조 자체도 계속 진화한다** ✅
  - **Living Cycle 다이어그램:**
    ```
           (새 요구사항 / 검증 발견)
                    ↑
    구조 설계(ADR) → Spec → Generation → Verification
        ↑                                    ↓
        └────────── (피드백 반영) ←───────────┘
    ```
  - **사이클의 흐름:**
    - 🔄 검증 → 구조 설계(ADR) 회귀 → 결정 재검토
    - 🔄 요구사항 변경 → ADR 변경 → Spec 갱신 → 재생성 → 재검증
  - **핵심 개념:**
    - 📄 **ADR도 살아있는 문서다** — 요구사항 변화에 따라 결정 변경 가능
    - 📄 **Spec도 살아있는 문서다** — ADR 변경이나 검증 결과 반영
    - 🏛️ **구조가 정해지면 변하지 않는 것이 아니다**
- **💡 포인트:** "구조 기반 개발은 Waterfall이 아니다. **구조가 개발을 이끌고, 개발 경험이 구조를 개선한다.** "

## Visual Elements
- 중앙: 순환 다이어그램 (화살표가 명확한 사이클)
- 좌측: Waterfall 이미지 (직선, 단방향 화살표, X 표시)
- 우측: Cycle 이미지 (원형, 순환 화살표, 체크)
- ADR과 Spec 아이콘이 함께 빛나는 효과

## Layout Notes
- 좌측 Waterfall(작게, 흐릿하게, X 표시) / 우측 Cycle(크게, 선명하게)
- 사이클 다이어그램은 중앙에 크게 배치
- 하단에 핵심 메시지 강조

## Reference
- Script: `scripts/script-12.md`
- Slide plan: `slides.md` → Slide 12
