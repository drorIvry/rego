import { DataTable } from "../ui/dataTable";
import { taskDefinitionsColumns } from "./columns";

const DefinitionsTable = async () => {
  const definitions = await getTaskDefinitions();
  return (
    <div className="container mx-auto py-10">
      <DataTable columns={taskDefinitionsColumns} data={definitions} />
    </div>
  );
};

export default DefinitionsTable;
