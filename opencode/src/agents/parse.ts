/**
 * Parse agent .md content (frontmatter + body). Shared by runtime and codegen.
 */

import { parse as parseYaml } from 'yaml';

const FRONTMATTER_RE = /^---\r?\n([\s\S]*?)\r?\n---\r?\n([\s\S]*)$/;

export type AgentMode = 'primary' | 'subagent' | 'all';

export interface AgentSpec {
  /** Agent key in OpenCode config.agent */
  name: string;
  prompt: string;
  /** Frontmatter fields passed through to OpenCode agent config (description, mode, permission, tools, …) */
  frontmatter: Record<string, unknown>;
}

function normalizeMode(value: unknown): AgentMode | undefined {
  if (value === 'primary' || value === 'subagent' || value === 'all') return value;
  if (typeof value === 'string') {
    const normalized = value.trim().toLowerCase();
    if (normalized === 'primary' || normalized === 'subagent' || normalized === 'all') {
      return normalized;
    }
  }
  return undefined;
}

function normalizeFrontmatter(raw: unknown): Record<string, unknown> {
  if (raw == null) return {};
  if (typeof raw !== 'object' || Array.isArray(raw)) {
    throw new Error('Agent frontmatter must be a YAML mapping');
  }
  return raw as Record<string, unknown>;
}

/** Parse frontmatter and body (prompt) from OpenCode agent markdown. */
export function parseAgentMd(content: string, slug?: string): AgentSpec {
  const match = content.match(FRONTMATTER_RE);
  if (!match) {
    throw new Error('Agent md must start with YAML frontmatter delimited by ---');
  }

  const [, frontmatterYaml, body] = match;
  const parsed = normalizeFrontmatter(parseYaml(frontmatterYaml));

  const name = String(parsed.name ?? slug ?? '').trim();
  if (!name) {
    throw new Error('Agent md must have frontmatter "name" or a filename slug');
  }

  const frontmatter: Record<string, unknown> = {};
  for (const [key, value] of Object.entries(parsed)) {
    if (key === 'name') continue;
    if (key === 'mode') {
      const mode = normalizeMode(value);
      if (mode) frontmatter.mode = mode;
      continue;
    }
    frontmatter[key] = value;
  }

  return {
    name,
    prompt: body.trim(),
    frontmatter,
  };
}

/** Convert parsed spec to OpenCode config.agent entry. */
export function specToAgentConfig(spec: AgentSpec): Record<string, unknown> {
  const entry: Record<string, unknown> = {
    ...spec.frontmatter,
    prompt: spec.prompt,
  };

  if (entry.model === 'inherit') {
    delete entry.model;
  }

  return entry;
}
