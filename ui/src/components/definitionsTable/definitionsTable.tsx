import { api } from "~/trpc/server";
import { DataTable } from "../ui/dataTable";
import { taskDefinitionsColumns } from "./columns";

const DefinitionsTable = async () => {
  console.log("test1");
  const definitions = await api.taskDefinitions.getTasks.query();
  console.log("test2");
  console.log(definitions[0]);

  return (
    <div className="container mx-auto py-10">
      <DataTable columns={taskDefinitionsColumns} data={definitions} />
    </div>
  );
};

export default DefinitionsTable;
