import { Plugin } from '@opencode/plugin';
import { registerAgents } from './agents/index.js';

export default Plugin.define({
  id: 'cocrates-harness',
  async setup(ctx) {
    console.log(
      `[cocrates-harness] initialized (opencode ${ctx.app.version}, dir=${ctx.location.directory})`,
    );

    await ctx.agent.transform((editor) => {
      registerAgents(editor);
    });

    const controller = new AbortController();
    void (async () => {
      try {
        for await (const event of ctx.event.subscribe({
          signal: controller.signal,
        })) {
          if (event.type === 'session.created') {
            console.debug('[cocrates-harness] session.created', {
              sessionId: event.data.sessionID,
            });
          }
        }
      } catch (error) {
        if (!controller.signal.aborted) throw error;
      }
    })();

    return () => controller.abort();
  },
});

export {
  registerAgents,
  parseAgentMd,
  specToAgentConfig,
  getAgentSlugs,
  loadAgentSpec,
} from './agents/index.js';
export type { AgentSpec, AgentMode } from './agents/index.js';
