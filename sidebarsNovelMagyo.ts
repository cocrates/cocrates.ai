import fs from 'node:fs';
import path from 'node:path';
import type {SidebarsConfig} from '@docusaurus/plugin-content-docs';

const manuscriptsDir = path.join(
  __dirname,
  'examples/마교-교주가-정파-맹주로-사기-친다/manuscripts',
);

const INTRO_ID = '마교-교주가-정파-맹주로-사기-친다';

function listChapterIds(): string[] {
  return fs
    .readdirSync(manuscriptsDir)
    .filter((filename) => /^\d{3}-.+\.md$/.test(filename))
    .map((filename) => filename.replace(/\.md$/, ''))
    .sort((a, b) => a.localeCompare(b, 'ko'));
}

function parseChapterId(id: string): {chapterNumber: number | null; title: string} {
  if (id === INTRO_ID) {
    return {chapterNumber: null, title: '마교 교주가 정파 맹주로 사기 친다'};
  }

  const numberedMatch = id.match(/^(\d{3})-(.+)$/);
  if (numberedMatch) {
    return {
      chapterNumber: Number.parseInt(numberedMatch[1], 10),
      title: numberedMatch[2].replace(/-/g, ' '),
    };
  }

  return {chapterNumber: null, title: id.replace(/-/g, ' ')};
}

function sidebarLabelFromId(id: string): string {
  const {chapterNumber, title} = parseChapterId(id);

  if (chapterNumber === null) {
    return title;
  }

  return `${chapterNumber}화. ${title}`;
}

const novelMagyoChapterIds = [INTRO_ID, ...listChapterIds()];

const sidebars: SidebarsConfig = {
  novelMagyoSidebar: novelMagyoChapterIds.map((id) => ({
    type: 'doc' as const,
    id,
    key: id,
    label: sidebarLabelFromId(id),
  })),
};

export default sidebars;
