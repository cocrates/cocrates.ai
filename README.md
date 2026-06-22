# Cocrates

Docusaurus 기반 문서 사이트입니다. GitHub Pages에 배포됩니다.

- **사이트 URL**: https://cocrates.ai/

## 로컬 개발

```bash
npm install
npm start
```

브라우저에서 http://localhost:3000/ 로 접속합니다.

## 빌드

```bash
npm run build
npm run serve
```

## 배포

`main` 브랜치에 push하면 GitHub Actions가 자동으로 GitHub Pages에 배포합니다.

최초 배포 시 GitHub 저장소 설정에서 **Settings → Pages → Build and deployment → Source**를 **GitHub Actions**로 선택해야 합니다.

사용자 지정 도메인(`cocrates.ai`) 연결 후 **Settings → Pages → Enforce HTTPS**를 활성화하세요. 인증서 발급까지 최대 24시간 걸릴 수 있습니다.
