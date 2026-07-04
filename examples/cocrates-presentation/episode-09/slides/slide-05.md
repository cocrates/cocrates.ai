# Slide 05: 🔥 3블록 출력 포맷 + DIP 실제 사례

## Key Message
Education 스킬의 3블록 출력 포맷(Concept Briefing → Thought Lab → MISSION)이 모든 교육 대화를 생성하는 단일 엔진이다. DIP 예시로 실제 작동을 확인한다.

## Content
### Headlines / text
- **3블록 출력 포맷 (SKILL.md 실제 규칙):**
  ```markdown
  ### 💡 [Concept Briefing]
  (핵심 원리를 1~3문장, 일상적 비유로 전달. 응답의 20% 이하)

  ### 💻 [Thought Lab]
  (일부러 결함을 심어두었거나 비어있는 불완전한 예시/시나리오)

  ### 🔥 [MISSION]
  (사용자가 다음 턴에 생각하고 답해야 할 정확히 하나의 인지적 과제)
  ```
- **각 블록 역할:**
  - Briefing: 최소한의 맥락 — "이게 왜 중요한가" (비유 사용)
  - Lab: 결함 있는 예시 — 사용자가 직접 오류 발견 및 수정
  - MISSION: 정확히 하나의 인지적 과제
- **DIP 실제 예시:**
  - 💡 **Briefing:** "DIP는 전기 플러그와 같습니다. 고수준 모듈이 저수준 세부사항에 직결되지 않고, 둘 다 공유 인터페이스에 의존해야 합니다."
  - 💻 **Thought Lab:**
    ```typescript
    class OrderService {
      constructor(private db: MySQLDatabase) {} // flaw
    }
    ```
  - 🔥 **MISSION:** "MySQLDatabase를 PostgresDatabase로 바꾸면 무엇이 깨질지 설명하고, DIP를 위반하는 의존 방향을 지적하세요."
- **💡 포인트:** 세 블록이 하나의 완전한 학습 턴. 규칙은 단순하지만 적용은 무한하다.

### Visual elements
- 상단: 3블록 코드 포맷 (회색 배경 코드 블록)
- 중앙: DIP 예시를 3블록 흐름도로 표시 (화살표 연결)
- 각 블록에 아이콘: 💡 Briefing / 💻 Lab / 🔥 Mission
- 하단: 💡 강조 박스

## Layout Notes
- 코드 블록은 실제 SKILL.md 포맷을 그대로 표시 (신뢰감)
- DIP 예시는 흐름도 형태로 시각화
- Briefing은 짧게, Lab은 코드로, MISSION은 질문 형태로

## Reference
- Script: `scripts/script-05.md`
- Slide plan: `slides.md` → Slide 05
