<script>
    import DataTable from "./DataTable.svelte";
    import StructureSelector from "./StructureSelector.svelte";

    const INSERT_URL = "http://localhost:3000/data";

    let tableName = "";
    let structure = [];

    function structureHandler(n, s) {
        tableName = n;
        structure = s;
    }

    // NOTE: we don't need to do much checking on
    // table name and data here.
    // If the Data table is visible and calling this handler,
    // it is gauranteed that tablenName and structure are valid.
    async function insertData(tName, data) {
        console.log(data);
        let resp = await fetch(
            INSERT_URL + "?table=" + tName,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                body: JSON.stringify(data)
            }
        )
        let respJSON = await resp.json();

        if (!resp.ok) {
            alert(respJSON.message);
            return;
        }

        alert("Insert Success!");
    }
</script>

<div class="title-bar">
    Insert Data
</div>
<div class="container">
    <div class="tight">
        <div class="rounded">
            <StructureSelector structureHandler={structureHandler} />
        </div>
    </div>
    
    {#if structure.length > 0}
        <div class="break">
        </div>

        <div class="rounded">
            <DataTable tableName={tableName} structure={structure} dataHandler={insertData} />
        </div>
    {/if}
</div>
