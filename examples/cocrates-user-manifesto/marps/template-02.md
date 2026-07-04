# Marp Template 02: Content

<!-- theme hint: uncover -->
<!-- root class: tpl-02 -->

```css
section.tpl-02 {
  background: #0a0a0a;
  color: #ffffff;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  padding: 6% 10%;
  font-family: 'Noto Sans KR', 'Pretendard', sans-serif;
  justify-content: flex-start;
}

section.tpl-02 h1 {
  font-size: 32pt;
  font-weight: 800;
  color: #ffffff;
  margin: 0 0 0.4em 0;
  text-align: center;
  line-height: 1.3;
}

section.tpl-02 .tpl-02__title-divider {
  width: 40%;
  height: 1px;
  background: #333333;
  margin: 0 auto 1.2em auto;
}

section.tpl-02 .tpl-02__body {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
  max-width: 90%;
  gap: 0.3em;
}

section.tpl-02 .tpl-02__footer {
  font-size: 11pt;
  font-weight: 300;
  color: #666666;
  text-align: center;
  padding-top: 1em;
  margin-top: auto;
}

/* Commandment text (large) */
section.tpl-02 .cmd {
  font-size: 42pt;
  font-weight: 900;
  color: #ffffff;
  text-align: center;
  line-height: 1.2;
  margin: 0.2em 0;
}

/* Korean subtitle for commandments */
section.tpl-02 .cmd-ko {
  font-size: 24pt;
  font-weight: 500;
  color: #999999;
  text-align: center;
  margin: 0.1em 0;
}

/* Description text */
section.tpl-02 .desc {
  font-size: 16pt;
  font-weight: 400;
  color: #cccccc;
  text-align: center;
  line-height: 1.5;
  margin: 0.3em 0;
  max-width: 85%;
}

/* Point / punchline */
section.tpl-02 .point {
  font-size: 15pt;
  font-weight: 400;
  color: #888888;
  text-align: center;
  margin: 0.5em 0;
}

/* Emphasis (accent red for rejection commandments) */
section.tpl-02 .accent-red {
  color: #ff3b30;
  font-weight: 800;
}

/* Emphasis (green for aspiration commandments) */
section.tpl-02 .accent-green {
  color: #34c759;
  font-weight: 800;
}

/* Highlight keyword */
section.tpl-02 .highlight {
  font-weight: 900;
  color: #ffffff;
}

/* Large transition word (e.g., 하지만...) */
section.tpl-02 .transition-lg {
  font-size: 54pt;
  font-weight: 900;
  color: #ffffff;
  text-align: center;
  margin: 0.3em 0;
}

/* Declaration text */
section.tpl-02 .declaration {
  font-size: 24pt;
  font-weight: 700;
  color: #ffffff;
  text-align: center;
  line-height: 1.4;
  margin: 0.5em 0;
  max-width: 85%;
}

/* Signal text (small, muted) */
section.tpl-02 .signal {
  font-size: 14pt;
  font-weight: 300;
  color: #666666;
  text-align: center;
  letter-spacing: 0.1em;
}

/* Manifesto list items */
section.tpl-02 .manifesto-item {
  font-size: 15pt;
  font-weight: 700;
  color: #ffffff;
  text-align: center;
  line-height: 1.6;
}

/* Equation elements */
section.tpl-02 .eq-left {
  font-size: 24pt;
  font-weight: 700;
  color: #888888;
}

section.tpl-02 .eq-op {
  font-size: 22pt;
  font-weight: 400;
  color: #666666;
}

section.tpl-02 .eq-center {
  font-size: 24pt;
  font-weight: 700;
  color: #ffffff;
}

section.tpl-02 .eq-right {
  font-size: 24pt;
  font-weight: 900;
  color: #ffffff;
}

/* Dense mode for compact slides */
section.tpl-02.compact .cmd {
  font-size: 36pt;
}

section.tpl-02.compact .manifesto-item {
  font-size: 13pt;
  line-height: 1.4;
}

section.tpl-02.compact .tpl-02__body {
  gap: 0.15em;
}
```

## Regions (Marp)

| Region | Markup / class |
|--------|----------------|
| Title | h1 |
| Title divider | .tpl-02__title-divider (hr) |
| Body | .tpl-02__body |
| Footer | .tpl-02__footer |

## Notes

- Content template — now center-aligned for all elements
- Uses utility classes (cmd, cmd-ko, desc, point, etc.) for consistent styling
- Compact class for dense slides (slide-13)
