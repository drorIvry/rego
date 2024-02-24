"use client";

import { ColumnDef } from "@tanstack/react-table";
import { Badge } from "~/components/ui/badge";
import { TaskDefinition } from "~/lib/models/taskDefinition";
import { TaskStatus } from "~/lib/models/taskStatus";

const STATUS_TO_COLOR_MAP: Record<TaskStatus, string> = {
  [TaskStatus.READY]: "gray",
  [TaskStatus.JOB_DEPLOYED]: "gray",
  [TaskStatus.PENDING]: "gray",
  [TaskStatus.RUNNING]: "yellow",
  [TaskStatus.TIMEOUT]: "red",
  [TaskStatus.PROC_ERROR]: "red",
  [TaskStatus.APP_ERROR]: "red",
  [TaskStatus.ABORTED]: "red",
  [TaskStatus.SUCCESS]: "green",
};

export const taskDefinitionsColumns: ColumnDef<TaskDefinition>[] = [
  {
    accessorKey: "id",
    header: "id",
  },
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "image",
    header: "Image",
  },
  {
    accessorKey: "latest_status",
    header: "Latest Status",
    cell: ({ row }) => {
      const task_status: TaskStatus = row.getValue("latest_status");
      const color = STATUS_TO_COLOR_MAP[task_status] ?? "white";

      return (
        <Badge
          style={{
            backgroundColor: color,
          }}
        >
          {task_status}
        </Badge>
      );
    },
  },
  {
    accessorKey: "execution_interval",
    header: "Interval",
  },
  {
    accessorKey: "execution_counter",
    header: "Counter",
  },
  {
    accessorKey: "next_execution_time",
    header: "Next Execution Time",
  },
  {
    accessorKey: "cmd",
    header: "Cmd",
  },
  {
    accessorKey: "metadata",
    header: "Metadata",
  },
];
