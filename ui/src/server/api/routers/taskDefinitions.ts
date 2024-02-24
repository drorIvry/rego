import { TaskDefinitionsApi } from "~/lib/regoApi/taskDefinitionsApi";
import { createTRPCRouter, publicProcedure } from "~/server/api/trpc";
import { getServerAuthSession } from "~/server/auth";

const taskDefinitionsApi = new TaskDefinitionsApi();

export const taskDefinitionsRouter = createTRPCRouter({
  getTasks: publicProcedure.query(async ({ ctx }) => {
    const session = await getServerAuthSession();
    const apiKey: string = "delete_me"; //getUserApiKey(session?.user);

    return await taskDefinitionsApi.getAllTaskDefinitions(apiKey);
  }),
});
