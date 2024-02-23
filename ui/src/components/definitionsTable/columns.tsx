"use client";

import { ColumnDef } from "@tanstack/react-table";
import { TaskDefinition } from "~/lib/models/taskDefinition";

export const taskDefinitionsColumns: ColumnDef<TaskDefinition>[] = [
  {
    accessorKey: "Image",
    header: "Image",
  },
  {
    accessorKey: "Name",
    header: "Name",
  },
  {
    accessorKey: "LatestStatus",
    header: "LatestStatus",
  },
  {
    accessorKey: "ExecutionInterval",
    header: "ExecutionInterval",
  },
  {
    accessorKey: "ExecutionsCounter",
    header: "ExecutionsCounter",
  },
  {
    accessorKey: "NextExecutionTime",
    header: "NextExecutionTime",
  },
  {
    accessorKey: "Cmd",
    header: "Cmd",
  },
  {
    accessorKey: "Metadata",
    header: "Metadata",
  },
];
