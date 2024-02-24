import { TaskStatus } from "./taskStatus";

export type TaskDefinition = {
  id: string;
  organization_id: string;
  image: string;
  name: string;
  namespace: string;
  latest_status: TaskStatus;
  execution_interval: number;
  execution_counter: number;
  next_execution_time: Date;
  enabled: boolean;
  deleted: boolean;
  cmd: string[];
  metadata: any;

  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt: Date;
};
