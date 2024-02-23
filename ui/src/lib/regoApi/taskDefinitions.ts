import { AxiosResponse } from "axios";
import { RegoApi } from "./baseApiInvoker";
import { TaskDefinition } from "../models/taskDefinition";

import { validate } from "class-validator";

export class TaskDefinitions extends RegoApi {
  public async getAllTaskDefinitions(api_key: string) {
    const response: AxiosResponse = await this.invokeApi(
      "get",
      "/api/v1/task",
      api_key,
    );

    validate();
    const tasks: TaskDefinition[] = response.data.forEach((element: any) => {
      return validate(element);
    });
    return response.data;
  }
}
