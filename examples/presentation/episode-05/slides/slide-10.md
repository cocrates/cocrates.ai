# Slide 10: 🔍 Verification — 설계 검증 + 동시에 설계 재검토

## Key Message
Verification은 두 가지 역할을 한다. (1) Spec대로 구현되었는지 확인, (2) 동시에 설계 자체가 올바른지 재검토. 복잡도가 높을수록 점진적/반복적 개발이 필요하다.

## Content
- **Headlines / text:**
  - **역할 1: "사양대로 구현되었는가?"**
    - Spec에서 **72개 체크 항목** 식별
    - 각 항목에 대해 설계(Spec)에 맞춰 구현되었는지 확인
    - 결과: ✅ 71개 PASS, ❌ **1개 FAIL** (Deviation)
    - 발견된 Deviation:
      - 📄 Spec: `PUT /schema/blog/episode`
      - 💻 코드: `PUT /schema//blog/episode` (이중 슬래시)
  - **📈 복잡도가 높을수록 AI가 사양을 놓칠 가능성은 높아진다**
    - → **작은 단위로 점진적/반복적 개발** 필요
    - 큰 Spec 한 번 구현 → 검증 부담 폭증
    - 작은 Spec씩 구현 → 검증+피드백 선순환
  - **역할 2: "설계 자체가 올바른가?"**
    - 검증하면서 동시에 설계의 올바름을 재검토
    - "사양대로 안 되면 AI 잘못" ❌ → **"설계 자체를 변경해야 하는 상황"** 일 수 있음
- **💡 포인트:** "Verification은 단순한 체크리스트가 아니다. **Spec 확인 + 설계 재검토의 기회다.** "

## Visual Elements
- 상단: 72개 체크리스트 이미지 (71개 초록, 1개 빨강)
- 중앙 좌측: Deviation 하이라이트 (이중 슬래시 시각화)
- 중앙 우측: 점진적 개발 사이클 아이콘 (작은 Spec → 구현 → 검증 → 다음 Spec)
- 하단: 저울 이미지 (Spec 대로 확인 ↔ 설계 재검토)

## Layout Notes
- 정보량이 많은 슬라이드이므로 3단 분할 권장
- 좌측: Verification 역할 1 (검증 결과)
- 중앙: 점진적 개발 메시지
- 우측: Verification 역할 2 (설계 재검토)
- 하단: 통합 포인트 메시지

## Reference
- Script: `scripts/script-10.md`
- Slide plan: `slides.md` → Slide 10
