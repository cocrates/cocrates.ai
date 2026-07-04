# Template 02: Content

## Derivation
- **Base:** template-01
- **Relationship:** extends
- **Changes from base:**
  - **Inherited:** background(#0a0a0a), font family, color palette, divider style
  - **Added:** title zone(top), content zone(middle), footer zone(bottom)
  - **Modified:** title smaller than cover, center-aligned

## Background
- **Color:** #0a0a0a (same as cover), full-bleed

## Typography
- **Font family:** 굵은 고딕 계열 sans-serif (일관성 유지)
- **Slide title:** 중간~큼(28~36pt), 굵기 700~800, 흰색(#ffffff), 중앙 정렬
- **Section heading:** 중간(20~24pt), 굵기 600, 흰색
- **Body:** 보통(16~18pt), 굵기 400, 흰색 또는 연회색
- **Caption / label:** 작게(12~14pt), 굵기 300, 회색(#888888)
- **Emphasis:** 굵기 700~900, 필요시 accent 컬러(#ffffff)
- **Commandment text:** 매우 큼(36~48pt), 굵기 800~900, 흰색, 중앙 정렬

## Color Palette
- **Background:** #0a0a0a
- **Primary text:** #ffffff
- **Secondary text:** #cccccc
- **Tertiary text:** #888888
- **Muted text:** #666666
- **Divider / border:** #333333
- **Accent (거부 계명):** #ff3b30 (강한 빨강 — ❌ 아이콘 등 강조)
- **Accent (지향 계명):** #34c759 (초록 — ✅ 아이콘 등 강조)
- **Highlight:** #ffffff (bold text on black)

## Regions
### Title zone
- **Position:** top, 중앙 정렬
- **Style:** slide title, 중앙 정렬
- **Max lines:** 1~2
- **Divider:** title 아래에 가는 선(#333333, 중앙 정렬, 전체 폭의 30~50%)

### Content zone
- **Bounds:** title 아래 ~ footer 위. 넉넉한 상하 여백
- **Default alignment:** center
- **Padding:** 좌우 10% 여백

### Footer zone
- **Position:** bottom, 중앙 정렬
- **Use:** slide number, 소스 표시 등 (선택적)
- **Style:** 작은 텍스트(10~12pt), #666666

## Supported content layouts
bullets | columns (2~3) | grid (2×2, 2×3) | rows / stack | focal | diagram | mixed | contrast

## Layout Defaults
- **Content alignment:** center
- **Whitespace:** generous — 선 굵은 미니멀 스타일, 넉넉한 여백이 곧 디자인

## Intended Use
content — default for body slides (Slides 02–13)
