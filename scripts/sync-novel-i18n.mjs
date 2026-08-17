import {
  existsSync,
  mkdirSync,
  readdirSync,
  readFileSync,
  writeFileSync,
} from 'node:fs';
import {dirname, join} from 'node:path';
import {fileURLToPath} from 'node:url';

const __dirname = dirname(fileURLToPath(import.meta.url));
const rootDir = join(__dirname, '..');

const enManuscriptDir = join(rootDir, 'examples/rejected-future/i18n/us/manuscripts');
const magyoManuscriptDir = join(
  rootDir,
  'examples/마교-교주가-정파-맹주로-사기-친다/manuscripts',
);
const novel120ManuscriptDir = join(rootDir, 'examples/120통의-편지/manuscripts');
const novelAbeojiManuscriptDir = join(
  rootDir,
  'examples/아버지를-죽인-건-나다/manuscripts',
);
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

const MAGYO_INTRO_ID = '마교-교주가-정파-맹주로-사기-친다';
const NOVEL_120_INTRO_ID = '120통의-편지';
const NOVEL_ABEOJI_INTRO_ID = '아버지를-죽인-건-나다';

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

function parseMagyoFilename(filename) {
  const base = filename.replace(/\.md$/, '');

  if (base === MAGYO_INTRO_ID) {
    return {chapterNumber: null, title: '마교 교주가 정파 맹주로 사기 친다'};
  }

  const numberedMatch = base.match(/^(\d{3})-(.+)$/);
  if (numberedMatch) {
    return {
      chapterNumber: Number.parseInt(numberedMatch[1], 10),
      title: numberedMatch[2].replace(/-/g, ' '),
    };
  }

  return {chapterNumber: null, title: base.replace(/-/g, ' ')};
}

function magyoSidebarLabel(filename) {
  const {chapterNumber, title} = parseMagyoFilename(filename);

  if (chapterNumber === null) {
    return title;
  }

  return `${chapterNumber}화. ${title}`;
}

function parseNovel120Filename(filename) {
  const base = filename.replace(/\.md$/, '');

  if (base === NOVEL_120_INTRO_ID) {
    return {chapterNumber: null, title: '120통의 편지'};
  }

  const numberedMatch = base.match(/^(\d{3})-(.+)$/);
  if (numberedMatch) {
    return {
      chapterNumber: Number.parseInt(numberedMatch[1], 10),
      title: numberedMatch[2].replace(/-/g, ' '),
    };
  }

  return {chapterNumber: null, title: base.replace(/-/g, ' ')};
}

function novel120SidebarLabel(filename) {
  const {chapterNumber, title} = parseNovel120Filename(filename);

  if (chapterNumber === null) {
    return title;
  }

  return `${chapterNumber}화. ${title}`;
}

function parseNovelAbeojiFilename(filename) {
  const base = filename.replace(/\.md$/, '');

  if (base === NOVEL_ABEOJI_INTRO_ID) {
    return {chapterNumber: null, title: '아버지를 죽인 건 나다'};
  }

  const numberedMatch = base.match(/^(\d{3})-(.+)$/);
  if (numberedMatch) {
    return {
      chapterNumber: Number.parseInt(numberedMatch[1], 10),
      title: numberedMatch[2].replace(/-/g, ' '),
    };
  }

  return {chapterNumber: null, title: base.replace(/-/g, ' ')};
}

function novelAbeojiSidebarLabel(filename) {
  const {chapterNumber, title} = parseNovelAbeojiFilename(filename);

  if (chapterNumber === null) {
    return title;
  }

  return `${chapterNumber}화. ${title}`;
}

function injectPaginationLabel(content, label) {
  const frontMatterMatch = content.match(/^---\r?\n([\s\S]*?)\r?\n---\r?\n?/);

  if (frontMatterMatch) {
    const existing = frontMatterMatch[1];
    const body = content.slice(frontMatterMatch[0].length);
    if (/^pagination_label:/m.test(existing)) {
      const updated = existing.replace(
        /^pagination_label:.*$/m,
        `pagination_label: ${JSON.stringify(label)}`,
      );
      return `---\n${updated}\n---\n${body}`;
    }
    return `---\npagination_label: ${JSON.stringify(label)}\n${existing}\n---\n${body}`;
  }

  return `---\npagination_label: ${JSON.stringify(label)}\n---\n${content}`;
}

function ensureDir(path) {
  mkdirSync(path, {recursive: true});
}

function syncRejectedFutureI18n() {
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
}

function syncMagyoI18n() {
  const magyoFiles = readdirSync(magyoManuscriptDir).filter((filename) =>
    filename.endsWith('.md'),
  );

  for (const locale of nonKoLocales) {
    const pluginDir = join(
      rootDir,
      `i18n/${locale}/docusaurus-plugin-content-docs-novel-magyo`,
    );
    const targetDir = join(pluginDir, 'current');
    ensureDir(targetDir);

    const sidebarTranslations = {};

    for (const filename of magyoFiles) {
      const docId = filename.replace(/\.md$/, '');
      const sourcePath = join(magyoManuscriptDir, filename);
      const targetPath = join(targetDir, filename);
      const paginationLabel = magyoSidebarLabel(filename);
      const content = readFileSync(sourcePath, 'utf8');

      writeFileSync(
        targetPath,
        injectPaginationLabel(content, paginationLabel),
        'utf8',
      );

      sidebarTranslations[`sidebar.novelMagyoSidebar.doc.${docId}`] = {
        message: paginationLabel,
        description: `Sidebar label for magyo novel chapter ${docId}`,
      };
    }

    writeFileSync(
      join(pluginDir, 'current.json'),
      `${JSON.stringify(sidebarTranslations, null, 2)}\n`,
      'utf8',
    );
  }
}

function syncNovel120I18n() {
  const novel120Files = readdirSync(novel120ManuscriptDir).filter((filename) =>
    filename.endsWith('.md'),
  );

  for (const locale of nonKoLocales) {
    const pluginDir = join(
      rootDir,
      `i18n/${locale}/docusaurus-plugin-content-docs-novel-120`,
    );
    const targetDir = join(pluginDir, 'current');
    ensureDir(targetDir);

    const sidebarTranslations = {};

    for (const filename of novel120Files) {
      const docId = filename.replace(/\.md$/, '');
      const sourcePath = join(novel120ManuscriptDir, filename);
      const targetPath = join(targetDir, filename);
      const paginationLabel = novel120SidebarLabel(filename);
      const content = readFileSync(sourcePath, 'utf8');

      writeFileSync(
        targetPath,
        injectPaginationLabel(content, paginationLabel),
        'utf8',
      );

      sidebarTranslations[`sidebar.novel120Sidebar.doc.${docId}`] = {
        message: paginationLabel,
        description: `Sidebar label for 120 novel chapter ${docId}`,
      };
    }

    writeFileSync(
      join(pluginDir, 'current.json'),
      `${JSON.stringify(sidebarTranslations, null, 2)}\n`,
      'utf8',
    );
  }
}

function syncNovelAbeojiI18n() {
  const novelAbeojiFiles = readdirSync(novelAbeojiManuscriptDir).filter(
    (filename) => filename.endsWith('.md'),
  );

  for (const locale of nonKoLocales) {
    const pluginDir = join(
      rootDir,
      `i18n/${locale}/docusaurus-plugin-content-docs-novel-abeoji`,
    );
    const targetDir = join(pluginDir, 'current');
    ensureDir(targetDir);

    const sidebarTranslations = {};

    for (const filename of novelAbeojiFiles) {
      const docId = filename.replace(/\.md$/, '');
      const sourcePath = join(novelAbeojiManuscriptDir, filename);
      const targetPath = join(targetDir, filename);
      const paginationLabel = novelAbeojiSidebarLabel(filename);
      const content = readFileSync(sourcePath, 'utf8');

      writeFileSync(
        targetPath,
        injectPaginationLabel(content, paginationLabel),
        'utf8',
      );

      sidebarTranslations[`sidebar.novelAbeojiSidebar.doc.${docId}`] = {
        message: paginationLabel,
        description: `Sidebar label for abeoji novel chapter ${docId}`,
      };
    }

    writeFileSync(
      join(pluginDir, 'current.json'),
      `${JSON.stringify(sidebarTranslations, null, 2)}\n`,
      'utf8',
    );
  }
}

function syncNovelI18n() {
  syncRejectedFutureI18n();
  syncMagyoI18n();
  syncNovel120I18n();
  syncNovelAbeojiI18n();
  console.log(
    `Synced novel i18n manuscripts for: ${nonKoLocales.join(', ')} (rejected-future + magyo + 120 + abeoji)`,
  );
}

syncNovelI18n();
