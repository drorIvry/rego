import { postRouter } from "~/server/api/routers/post";
import { createTRPCRouter } from "~/server/api/trpc";
import { taskDefinitionsRouter } from "./routers/taskDefinitions";

/**
 * This is the primary router for your server.
 *
 * All routers added in /api/routers should be manually added here.
 */
export const appRouter = createTRPCRouter({
  post: postRouter,
  taskDefinitions: taskDefinitionsRouter,
});

// export type definition of API
export type AppRouter = typeof appRouter;
