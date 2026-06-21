import { type Plugin } from '@opencode-ai/plugin';
import { registerAgents } from './agents/index.js';

export const CocratesPlugin: Plugin = async ({ client, directory, project }) => {
  await client.app.log({
    body: {
      service: 'cocrates-harness',
      level: 'info',
      message: 'Cocrates Harness plugin initialized',
      extra: {
        directory,
        projectId: project.id,
      },
    },
  });

  return {
    config: async (config) => {
      registerAgents(config);
    },
    event: async ({ event }) => {
      if (event.type === 'session.created') {
        await client.app.log({
          body: {
            service: 'cocrates-harness',
            level: 'debug',
            message: 'New session created',
            extra: {
              sessionId: event.properties.info.id,
            },
          },
        });
      }
    },
  };
};

export { registerAgents, parseAgentMd, specToAgentConfig, getAgentSlugs, loadAgentSpec } from './agents/index.js';
export type { AgentSpec, AgentMode } from './agents/index.js';
