<script>
    import Choice from "./Choice.svelte";
    import { TYPES, TYPE_IDS } from "./Data";

    // This property should be a function which
    // receives a copy of the represented structure.
    export let structureHandler;

    let structure = [];

    function addRow() {
        structure.push({"name": "Name", "type_id": 0});
        structure = structure;  // Rerender pls.
    }

    function removeRow(i) {
        structure.splice(i, 1)
        structure = structure;
    }

    function submitStructure() {
        // The map below ensures that we send a deep copy to whomever is
        // using this structure.
        structureHandler(
            structure.map(c => ({"name": c.name, "type_id": c.type_id}))
        );
    }
</script>

<div class="struct-container wide">
    <h1 class="header">
        Specify Structure
    </h1>
    <div class="struct-table">
        {#each structure as column, i}
            <div class="struct-row divider">
                <div class="flexible">
                    <input bind:value={column.name} />
                </div>
                <div class="inflexible">
                    <Choice choices={TYPES} 
                        ci={structure[i].type_id}
                        choiceHandler={(tn) => structure[i].type_id = TYPE_IDS[tn]} />
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
    .struct-container {
        border-radius: 10px;
        overflow: hidden;
    }
    
    .header {
        text-align: center;
        font-size: 1.5em;
    }

    .struct-row {
        height: 1.5em;
        display: flex;
    }
</style>
