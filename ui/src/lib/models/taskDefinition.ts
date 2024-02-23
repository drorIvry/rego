import { TaskStatus } from "./taskStatus";

export type TaskDefinition = {
  id: string;
  OrganizationId: string;
  Image: string;
  Name: string;
  Namespace: string;
  LatestStatus: TaskStatus;
  ExecutionInterval: number;
  ExecutionsCounter: number;
  NextExecutionTime: Date;
  Enabled: boolean;
  Deleted: boolean;
  Cmd: string[];
  Metadata: any;
};
