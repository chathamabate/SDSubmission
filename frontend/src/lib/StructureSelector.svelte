<script>
    import Choice from "./Choice.svelte";
    import { TYPES, TYPE_IDS } from "./Data";

    // This property should be a function which
    // receives a copy of the represented structure.
    export let structureHandler;

    let tableName = "";
    let structure = [];

    function addRow() {
        structure.push({"name": "", "typeID": 0});
        structure = structure;  // Rerender pls.
    }

    function removeRow(i) {
        structure.splice(i, 1)
        structure = structure;
    }

    function submitStructure() {
        // First check that everything is valid.
        if (tableName.length == 0) {
            alert("Please enter a table name.");
            return;
        }

        if (structure.length == 0) {
            alert("Please add columns.");
            return;
        }

        let keySet = new Set();
        for (const c of structure) {
            if (c.name.length == 0) {
                alert("Column names must be non-empty.");
                return;
            }

            if (keySet.has(c.name)) {
                alert("Column names must be unique.");
                return;
            }

            keySet.add(c.name);
        }


        // The map below ensures that we send a deep copy to whomever is
        // using this structure.
        structureHandler(
            tableName, structure.map(c => ({"name": c.name, "typeID": c.typeID}))
        );
    }
</script>

<div class="struct-container wide">
    <div class="struct-table">
        <div class="struct-row divider">
            <input bind:value={tableName} placeholder="Table" />
        </div>
        {#each structure as column, i}
            <div class="struct-row divider">
                <div class="flexible">
                    <input bind:value={column.name} placeholder="Column" />
                </div>
                <div class="inflexible">
                    <Choice choices={TYPES} 
                        ci={structure[i].typeID}
                        choiceHandler={(tn) => structure[i].typeID = TYPE_IDS[tn]} />
                </div>
                <div class="inflexible">
                    <button class="tall padded-element exit-button" on:click={() => removeRow(i)}>
                        X
                    </button>
                </div>
            </div>
        {/each}
        </div>
    <div class="flex">
        <button class="flexible padded-element simple-button" on:click={addRow}>
            Add Column
        </button>
        <button class="flexible padded-element simple-button" on:click={submitStructure}>
            Create Request
        </button>
    </div>
</div>

<style>
    .struct-row {
        height: 1.5em;
        display: flex;
    }
</style>
