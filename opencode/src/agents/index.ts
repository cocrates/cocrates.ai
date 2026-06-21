/**
 * Agent registration. Uses embedded specs (specs.gen.ts) so dist works without .md files.
 * Generate: npm run generate:agents
 */

import type { Config } from '@opencode-ai/plugin';
import { getAgentSpecs, getAgentSlugsFromSpecs } from './specs.gen.js';
import { specToAgentConfig, type AgentSpec } from './parse.js';

export { parseAgentMd, specToAgentConfig } from './parse.js';
export type { AgentSpec, AgentMode } from './parse.js';

/** Load agent spec by slug (from embedded specs). */
export function loadAgentSpec(slug: string): AgentSpec {
  const spec = getAgentSpecs()[slug];
  if (!spec) throw new Error(`Unknown agent slug: ${slug}`);
  return spec;
}

/** Get sorted list of agent slugs (cocrates first, then alphabetical). */
export function getAgentSlugs(): string[] {
  return getAgentSlugsFromSpecs();
}

/**
 * Register all agents into config.agent.
 * OpenCode plugin config callback receives Config; mutating config.agent is the intended API.
 */
export function registerAgents(config: Config): void {
  config.agent ??= {};

  for (const slug of getAgentSlugs()) {
    const spec = loadAgentSpec(slug);
    config.agent[spec.name] = specToAgentConfig(spec);
  }
}
