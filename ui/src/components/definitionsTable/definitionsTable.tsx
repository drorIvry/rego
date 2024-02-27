import { api } from "~/trpc/server";
import { DataTable } from "../ui/dataTable";
import { taskDefinitionsColumns } from "./columns";

const DefinitionsTable = async () => {
  const definitions = await api.taskDefinitions.getTasks.query();

  return (
    <div className="container mx-auto py-10">
      <DataTable columns={taskDefinitionsColumns} data={definitions} />
    </div>
  );
};

export default DefinitionsTable;
