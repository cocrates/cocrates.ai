# Cocrates OpenCode Plugin

Cocrates 저장소의 OpenCode **V2** 플러그인입니다. Docusaurus 문서 사이트와 **별도로** 빌드·배포합니다.

## 기능

- `src/agents/prompts/`의 OpenCode 에이전트 명세(md)를 빌드 시 임베드하고, `ctx.agent.transform`으로 에이전트 등록
- Frontmatter 지원 (`mode`, `permissions`, `description`, `color`, `steps` 등; body → `system`)
- `session.created` 이벤트 로깅

## OpenCode V2 요구사항

OpenCode **2.x**가 필요합니다. V1 함수형 플러그인 API(`@opencode-ai/plugin`)는 더 이상 로드되지 않습니다.

V2는 default export에 `{ id, setup }` (또는 `{ id, effect }`)가 있어야 합니다. 설정 키도 `plugin` → `plugins`로 바뀌었습니다.

## 개발

```bash
cd opencode
npm install
npm run build
```

에이전트 md를 수정한 뒤에는 `npm run generate:agents`로 `specs.gen.ts`를 재생성할 수 있습니다. `npm run build`가 이 단계를 자동으로 실행합니다.

파일 변경을 감시하며 빌드:

```bash
npm run dev
```

## OpenCode에서 사용

### npm 패키지로 (배포 후)

```json
{
  "$schema": "https://opencode.ai/config.json",
  "plugins": ["@cocrates/cocrates-harness"]
}
```

버전 없이 패키지 이름만 적으면 OpenCode가 `@latest`로 해석합니다. 첫 설치 시점의 버전이 캐시되며, npm에 새 버전이 올라와도 **자동 갱신되지 않을 수 있습니다**.

특정 버전을 고정하려면:

```json
{
  "$schema": "https://opencode.ai/config.json",
  "plugins": ["@cocrates/cocrates-harness@0.2.0"]
}
```

최신 버전으로 갱신하려면:

```bash
opencode plugin add @cocrates/cocrates-harness -f
```

또는 캐시를 삭제한 뒤 OpenCode를 다시 시작합니다.

```bash
rm -rf ~/.cache/opencode/packages/@cocrates/cocrates-harness@latest
```

### 로컬 빌드 결과로

```bash
npm run build
```

로컬 디렉터리로 등록할 때는 **디렉터리**를 지정합니다 (V2는 단일 `.js` 파일 경로를 건너뜁니다):

```json
{
  "$schema": "https://opencode.ai/config.json",
  "plugins": ["/home/you/work/cocrates.ai/opencode"]
}
```

`package.json`의 `exports["."]`가 `dist/index.js`를 가리키므로, 빌드 후 위처럼 디렉터리를 넣으면 됩니다.

## 배포

### GitHub Actions (CI)

`opencode/**` 변경 시 `.github/workflows/opencode-plugin.yml`이 빌드를 검증합니다.

### GitHub Release

`opencode-v*` 태그를 push하면 빌드 산출물이 Release에 첨부됩니다.

```bash
git tag opencode-v0.2.0
git push origin opencode-v0.2.0
```

### npm 배포

```bash
npm publish --access public
```
