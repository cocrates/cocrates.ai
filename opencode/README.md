# Cocrates OpenCode Plugin

Cocrates 저장소의 OpenCode 플러그인입니다. Docusaurus 문서 사이트와 **별도로** 빌드·배포합니다.

## 기능

- `src/agents/prompts/`의 OpenCode 에이전트 명세(md)를 빌드 시 임베드하고, 플러그인 `config` 훅으로 `config.agent`에 등록
- Frontmatter 전체 지원 (`mode`, `permission`, `tools`, `temperature` 등)
- `session.created` 이벤트 로깅

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
  "plugin": ["@cocrates/cocrates-harness"]
}
```

### 로컬 빌드 결과로

```bash
npm run build
```

저장소 루트의 `opencode.jsonc`가 로컬 플러그인과 `skills/` 경로를 프로젝트 설정으로 등록합니다. 이 저장소에서 OpenCode를 실행하면 자동으로 적용됩니다.

다른 프로젝트에서 단독으로 쓸 때는 `opencode.json`에 절대 경로로 등록:

```json
{
  "$schema": "https://opencode.ai/config.json",
  "plugin": ["file:///home/you/work/cocrates.ai/opencode/dist/index.js"]
}
```

## 배포

### GitHub Actions (CI)

`opencode/**` 변경 시 `.github/workflows/opencode-plugin.yml`이 빌드를 검증합니다.

### GitHub Release

`opencode-v*` 태그를 push하면 빌드 산출물이 Release에 첨부됩니다.

```bash
git tag opencode-v0.1.0
git push origin opencode-v0.1.0
```

### npm 배포

```bash
cd opencode
npm publish --access public
```
