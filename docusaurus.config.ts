import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

const config: Config = {
  title: 'Cocrates',
  tagline: 'AI에게 맡기지 말고, AI를 지휘하라',
  favicon: 'img/logo.svg',

  future: {
    v4: true,
  },

  url: 'https://cocrates.ai',
  baseUrl: '/',

  organizationName: 'cocrates',
  projectName: 'cocrates',

  onBrokenLinks: 'throw',

  markdown: {
    mermaid: true,
    format: 'detect',
  },
  themes: ['@docusaurus/theme-mermaid'],

  i18n: {
    defaultLocale: 'ko',
    locales: ['ko', 'en', 'zh-Hans', 'id'],
    localeConfigs: {
      ko: {
        label: '한국어',
      },
      en: {
        label: 'English',
      },
      'zh-Hans': {
        label: '中文',
        htmlLang: 'zh-Hans',
        baseUrl: '/cn/',
      },
      id: {
        label: 'Bahasa Indonesia',
        htmlLang: 'id',
        baseUrl: '/id/',
      },
    },
  },

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.ts',
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
      } satisfies Preset.Options,
    ],
  ],

  plugins: [
    [
      '@docusaurus/plugin-content-docs',
      {
        id: 'novel',
        path: 'examples/rejected-future/manuscripts',
        routeBasePath: 'novel',
        sidebarPath: './sidebarsNovel.ts',
        numberPrefixParser: false,
      },
    ],
  ],

  themeConfig: {
    image: 'img/logo.svg',
    colorMode: {
      respectPrefersColorScheme: true,
    },
    navbar: {
      title: 'Cocrates',
      logo: {
        alt: 'Cocrates Logo',
        src: 'img/logo.svg',
      },
      items: [
        {
          type: 'docSidebar',
          sidebarId: 'docsSidebar',
          position: 'left',
          label: '문서',
        },
        {
          type: 'docSidebar',
          sidebarId: 'novelSidebar',
          docsPluginId: 'novel',
          position: 'left',
          label: '소설',
        },
        {
          type: 'localeDropdown',
          position: 'right',
        },
        {
          href: 'https://github.com/cocrates/cocrates.ai',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [],
      copyright: `Copyright © ${new Date().getFullYear()} Cocrates.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
