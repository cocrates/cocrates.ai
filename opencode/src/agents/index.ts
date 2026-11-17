/**
 * Agent registration. Uses embedded specs (specs.gen.ts) so dist works without .md files.
 * Generate: npm run generate:agents
 */

import { getAgentSpecs, getAgentSlugsFromSpecs } from './specs.gen.js';
import { applySpecToAgent, type AgentSpec } from './parse.js';

export { parseAgentMd, specToAgentConfig, applySpecToAgent } from './parse.js';
export type { AgentSpec, AgentMode, AgentConfigPatch } from './parse.js';

/** Minimal editor surface used by registerAgents (matches AgentEditor.update). */
export interface AgentRegistrationEditor {
  update(id: string, update: (agent: Record<string, unknown>) => void): void;
}

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
 * Register all agents via AgentEditor.update (upsert).
 * Missing agents are created from Agent.Info.default(id).
 */
export function registerAgents(editor: AgentRegistrationEditor): void {
  for (const slug of getAgentSlugs()) {
    const spec = loadAgentSpec(slug);
    editor.update(spec.name, (agent) => {
      applySpecToAgent(agent, spec);
    });
  }
}
