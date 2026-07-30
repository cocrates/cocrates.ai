# TODO: 치카치카 티라노 (영상)

> **Project root:** `/home/drajin/cocrates/chika-chika-t-rex-video`
> **Updated:** 2026-07-30

## Snapshot

| Done | In progress | Pending | Blocked | Skipped |
|------|-------------|---------|---------|---------|
| 13   | 0           | 1       | 0       | 0       |

**Current focus:** T-005 — G5 gate 대기 (사용자 승인)
**Recommended next:** 사용자 최종 영상 승인 → G5 → 프로젝트 완료

## Completed

- [x] **T-001** `done` — Define: overview.md G1 승인
- [x] **T-002** `done` — Design: references.md, sequence.md G2 승인
- [x] **T-003** `done` — Evaluate: evaluations/001-main.md G3 design lock
- [x] **T-004.0** `done` — 참조 이미지 복사 및 clip key 이미지 생성 (13장)
- [x] **T-004.1a** `done` — voice-spec.md 작성
- [x] **T-004.1b** `done` — transition-plan.md 작성
- [x] **T-004.2a** `done` — 13개 비디오 YAML 작성 (전환 로직 포함)
- [x] **T-004.2b** `done` — Generate Batch 1 (Clip 00~03): 4개 클립 생성 ✅
- [x] **T-004.2c** `done` — Generate Batch 2 (Clip 04~08): 5개 클립 생성 ✅
- [x] **T-004.2d** `done` — Generate Batch 3 (Clip 09~12): 4개 클립 생성 ✅

## Backlog

- [x] **T-004.2e** `done` — 13개 클립 생성 완료 (사용자 승인)
  - Phase: Generate Components (Stage ④)

- [ ] **T-005** `pending` — G5 gate: 최종 영상 사용자 승인
  - Phase: Assemble (Stage ⑤)
  - Artifact: `output/치카치카-티라노.mp4` ✅ 생성됨 (68MB, ~112초)

## Notes
- Short 모드 (단일 세그먼트)
- 모든 13개 비디오 클립 생성 완료 (`videos/001-main/*.mp4`)
- 각 클립에 대사/효과음/BGM이 내장됨 (omni 모델 자체 생성)
- 모델: `gemini-omni-flash-preview`
- 최종 mp4: ffmpeg concatenation으로 13개 클립 순차 합성
- 예상 총 길이: ~110초 (약 1분 50초)
