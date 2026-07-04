# Slide Content Specification — Slide 9

## Slide 9: 🔗 파이프라인 연결 + 요약 매트릭스

**Section:** 본론 3 — 연결과 요약
**Type:** 요약/시각화
**Time:** ~2분

### Visual Design
- **Background:** 동일 템플릿 (다크 계열, Slide 1과 유사)
- **Layout:** 상단 50% 파이프라인 다이어그램 + 하단 50% 요약 매트릭스
  - 상단: 전체 파이프라인 플로우 (가로형)
  - 하단: 4개 스킬 요약 매트릭스 + 협력 관계

### Content Elements
| 요소 | 내용 | 비고 |
|------|------|------|
| 제목 | 🔗 파이프라인 연결 + 요약 매트릭스 | 상단 좌측 |
| 다이어그램 | [Routing] → [Gate] → [adr-writing] → [spec-writing] → [gen] → [ver] → [순환] | 상단 50%, 가로 플로우 |
| 매트릭스 | 스킬 / 핵심 역할 / 절대 금지 사항 (4행) | 하단 50% |
| 협력 관계 | adr→spec / spec→gen / gen→ver / ver→spec (순환) | 매트릭스 하단, 작은 텍스트 |
| Key Point | 💡 각 스킬은 각자의 역할에 집중, **인터페이스(파일)로만 소통** | 하단 강조 박스 |

### Pipeline Diagram Detail
```
[Routing] → [spec-driven-gen Gate] → [adr-writing] → [spec-writing]
                                          ↓                ↓
                                     사용자 결정        Spec 완성
                                          ↓                ↓
                                 [spec-driven-gen 생성] → [spec-driven-verification]
                                                                  ↓
                                                           [순환: 검증 결과 → Spec 개정]
```

### Summary Matrix Detail
| 스킬 | 핵심 역할 | 절대 금지 사항 |
|------|---------|--------------|
| adr-writing | 대안 분석 & 결정 기록 | 가짜 대안, 산문 스타일, 무단 승인 |
| spec-writing | 결정 통합 설계도 | ADR 참조, 장황한 줄글 |
| spec-driven-gen | Spec 기준 생성 | Spec 없이 생성 (Gate 통과 필수) |
| spec-driven-ver | Deviation + Undoc ASR 발견 | 사용자 컨펌 없이 독단적 수정 |

### Visual Notes
- 매트릭스의 '절대 금지 사항' 열에 빨간색 ❌ 아이콘 추가
- 다이어그램의 화살표에 입력/출력 파일명 표시 (선택사항)

### Transitions
- **In:** 다이어그램 먼저 → 매트릭스 아래에서 등장
- **Out:** Key Point 표시 (이 슬라이드가 본론 마지막)
