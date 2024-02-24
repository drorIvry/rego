import { AxiosResponse } from "axios";
import { RegoApi } from "./baseApiInvoker";
import { TaskDefinition } from "../models/taskDefinition";

import { validate } from "class-validator";

export class TaskDefinitionsApi extends RegoApi {
  public async getAllTaskDefinitions(
    api_key: string,
  ): Promise<TaskDefinition[]> {
    const response: AxiosResponse = await this.invokeApi(
      "get",
      "/api/v1/task",
      api_key,
    );

    const tasks: TaskDefinition[] = response.data;
    console.log(tasks[0]);
    return tasks;
  }
}
