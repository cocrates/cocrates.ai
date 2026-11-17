/**
 * Parse agent .md content (frontmatter + body). Shared by runtime and codegen.
 */

import { parse as parseYaml } from 'yaml';

const FRONTMATTER_RE = /^---\r?\n([\s\S]*?)\r?\n---\r?\n([\s\S]*)$/;

export type AgentMode = 'primary' | 'subagent' | 'all';

export interface AgentSpec {
  /** Agent key / id in OpenCode agents map */
  name: string;
  /** Markdown body → V2 `system` prompt */
  prompt: string;
  /** Frontmatter fields (description, mode, permissions, …) */
  frontmatter: Record<string, unknown>;
}

export type PermissionRule = {
  action: string;
  resource: string;
  effect: 'allow' | 'deny' | 'ask';
};

/** V2 agent fields applied through AgentEditor.update */
export interface AgentConfigPatch {
  system: string;
  name?: string;
  description?: string;
  mode?: AgentMode;
  color?: string;
  steps?: number;
  hidden?: boolean;
  model?: {
    providerID: string;
    id: string;
    variant?: string;
  };
  permissions?: PermissionRule[];
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

function parseModelRef(value: unknown): AgentConfigPatch['model'] | undefined {
  if (typeof value !== 'string') return undefined;
  const trimmed = value.trim();
  if (!trimmed || trimmed === 'inherit') return undefined;

  const [base, variant] = trimmed.split('#');
  const slash = base.indexOf('/');
  if (slash <= 0 || slash === base.length - 1) return undefined;

  const providerID = base.slice(0, slash);
  const id = base.slice(slash + 1);
  return variant ? { providerID, id, variant } : { providerID, id };
}

function parsePermissions(value: unknown): PermissionRule[] | undefined {
  if (!Array.isArray(value)) return undefined;
  const rules: PermissionRule[] = [];
  for (const item of value) {
    if (!item || typeof item !== 'object' || Array.isArray(item)) continue;
    const row = item as Record<string, unknown>;
    const action = typeof row.action === 'string' ? row.action : undefined;
    const resource = typeof row.resource === 'string' ? row.resource : undefined;
    const effect = row.effect;
    if (!action || !resource) continue;
    if (effect !== 'allow' && effect !== 'deny' && effect !== 'ask') continue;
    rules.push({ action, resource, effect });
  }
  return rules.length > 0 ? rules : undefined;
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

/** Convert parsed spec to a V2 agent field patch (`prompt` → `system`). */
export function specToAgentConfig(spec: AgentSpec): AgentConfigPatch {
  const fm = spec.frontmatter;
  const entry: AgentConfigPatch = {
    system: spec.prompt,
    name: spec.name,
  };

  if (typeof fm.description === 'string') entry.description = fm.description;

  const mode = normalizeMode(fm.mode);
  if (mode) entry.mode = mode;

  if (typeof fm.color === 'string') entry.color = fm.color;
  if (typeof fm.hidden === 'boolean') entry.hidden = fm.hidden;

  if (typeof fm.steps === 'number' && Number.isInteger(fm.steps) && fm.steps > 0) {
    entry.steps = fm.steps;
  } else if (typeof fm.maxSteps === 'number' && Number.isInteger(fm.maxSteps) && fm.maxSteps > 0) {
    entry.steps = fm.maxSteps;
  }

  const model = parseModelRef(fm.model);
  if (model) entry.model = model;

  const permissions = parsePermissions(fm.permissions);
  if (permissions) entry.permissions = permissions;

  return entry;
}

/** Apply a parsed spec onto a mutable agent object from AgentEditor.update. */
export function applySpecToAgent(agent: Record<string, unknown>, spec: AgentSpec): void {
  const patch = specToAgentConfig(spec);
  agent.system = patch.system;
  if (patch.name !== undefined) agent.name = patch.name;
  if (patch.description !== undefined) agent.description = patch.description;
  if (patch.mode !== undefined) agent.mode = patch.mode;
  if (patch.color !== undefined) agent.color = patch.color;
  if (patch.hidden !== undefined) agent.hidden = patch.hidden;
  if (patch.steps !== undefined) agent.steps = patch.steps;
  if (patch.model !== undefined) agent.model = patch.model;
  if (patch.permissions !== undefined) agent.permissions = patch.permissions;
}
