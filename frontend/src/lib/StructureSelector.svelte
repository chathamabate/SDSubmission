<script>
    import Choice from "./Choice.svelte";

    export const TYPES = [
        "REAL", "TEXT"
    ];
    export const TYPE_IDS = {
        "REAL": 0,  // ID = index in type array.
        "TEXT": 1,
    };

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

        console.log(structure)
    }
</script>

<div class="struct-container wide">
    <h1 class="header">
        Specify Structure
    </h1>
    <div class="struct-table">
        {#each structure as column, i}
            <div class="struct-row">
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
    <div class="footer">
        <button class="flexible padded-element simple-button" on:click={addRow}>
            Add Column
        </button>
        <button class="flexible padded-element simple-button" on:click={addRow}>
            Create Insert Request
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

    .footer {
        display: flex;
    }
</style>
