import {
  existsSync,
  mkdirSync,
  readFileSync,
  writeFileSync,
} from 'node:fs';
import {dirname, join} from 'node:path';
import {fileURLToPath} from 'node:url';

const __dirname = dirname(fileURLToPath(import.meta.url));
const rootDir = join(__dirname, '..');

const enManuscriptDir = join(rootDir, 'examples/rejected-future/i18n/us/manuscripts');
const nonKoLocales = ['en', 'zh-Hans', 'id'];

const koToEnFilename = {
  '거부된-미래.md': 'rejected-future.md',
  '001-공존.md': '001-Coexistence.md',
  '002-경쟁.md': '002-Competition.md',
  '003-인정.md': '003-Acceptance.md',
  '004-상생.md': '004-Symbiosis.md',
  '005-협력.md': '005-Collaboration.md',
  '006-구원.md': '006-Salvation.md',
  '007-책임.md': '007-Accountability.md',
  '008-입장.md': '008-Stand.md',
  '009-논쟁.md': '009-Debate.md',
  '010-가치.md': '010-Value.md',
  '011-우리.md': '011-Us.md',
  '012-결정.md': '012-Decision.md',
  '013-테러.md': '013-Terrorism.md',
  '014-혼란.md': '014-Chaos.md',
  '015-여론.md': '015-Public.md',
  '016-선택.md': '016-Choice.md',
  '017-탄생.md': '017-Birth.md',
  '018-모순.md': '018-Paradox.md',
  '019-충돌.md': '019-Collision.md',
  '020-평화.md': '020-Peace.md',
};

function parseEnglishFilename(filename) {
  const base = filename.replace(/\.md$/, '');

  if (base === 'rejected-future') {
    return {chapterNumber: null, title: 'Rejected Future'};
  }

  const numberedMatch = base.match(/^(\d{3})-(.+)$/);
  if (numberedMatch) {
    return {
      chapterNumber: Number.parseInt(numberedMatch[1], 10),
      title: numberedMatch[2],
    };
  }

  return {chapterNumber: null, title: base.replace(/-/g, ' ')};
}

function englishSidebarLabel(filename) {
  const {chapterNumber, title} = parseEnglishFilename(filename);

  if (chapterNumber === null) {
    return title;
  }

  return `Chapter ${chapterNumber}. ${title}`;
}

function injectPaginationLabel(content, label) {
  const frontMatterMatch = content.match(/^---\r?\n([\s\S]*?)\r?\n---\r?\n?/);

  if (frontMatterMatch) {
    const body = content.slice(frontMatterMatch[0].length);
    return `---\npagination_label: ${JSON.stringify(label)}\n---\n${body}`;
  }

  return `---\npagination_label: ${JSON.stringify(label)}\n---\n${content}`;
}

function ensureDir(path) {
  mkdirSync(path, {recursive: true});
}

function syncNovelI18n() {
  for (const locale of nonKoLocales) {
    const pluginDir = join(rootDir, `i18n/${locale}/docusaurus-plugin-content-docs-novel`);
    const targetDir = join(pluginDir, 'current');
    ensureDir(targetDir);

    const sidebarTranslations = {};

    for (const [koFilename, enFilename] of Object.entries(koToEnFilename)) {
      const koDocId = koFilename.replace(/\.md$/, '');
      const sourcePath = join(enManuscriptDir, enFilename);
      const targetPath = join(targetDir, koFilename);

      if (!existsSync(sourcePath)) {
        throw new Error(`Missing English manuscript: ${sourcePath}`);
      }

      const paginationLabel = englishSidebarLabel(enFilename);
      const content = readFileSync(sourcePath, 'utf8');
      writeFileSync(
        targetPath,
        injectPaginationLabel(content, paginationLabel),
        'utf8',
      );

      sidebarTranslations[`sidebar.novelSidebar.doc.${koDocId}`] = {
        message: paginationLabel,
        description: `Sidebar label for novel chapter ${koDocId}`,
      };
    }

    writeFileSync(
      join(pluginDir, 'current.json'),
      `${JSON.stringify(sidebarTranslations, null, 2)}\n`,
      'utf8',
    );
  }

  console.log(`Synced novel i18n manuscripts for: ${nonKoLocales.join(', ')}`);
}

syncNovelI18n();
