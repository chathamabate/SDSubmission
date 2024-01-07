<script>
    import Choice from "./Choice.svelte";

    export const TYPE_MAP = {
        "REAL": 0,
        "TEXT": 1,
    };

    let structure = [];

    function addRow() {
        structure.push({"name": "Name", "type": 0});
        structure = structure;  // Rerender pls.

        console.log(structure)
    }

    function choiceHandler(i) {
        return (typeName) => {
            structure[i].type = TYPE_MAP[typeName]
        };
    }

    function removeRow(i) {
        structure.splice(i, 1)
        structure = structure;
    }
</script>

<div>
    <table class="struct-table">
        <tr>
            <th>Column Name</th>
            <th>Type</th>
            <th>.</th>
        </tr>
        {#each structure as column, i}
            <tr>
                <th>
                    <input bind:value={column.name} />
                </th>
                <th>
                    <Choice choices={Object.keys(TYPE_MAP)} choiceHandler={(tn) => structure[i].type = TYPE_MAP[tn]} />
                </th>
                <th>
                    <button class="remove-row" on:click={() => removeRow(i)}>
                        -
                    </button>
                </th>
            </tr>
        {/each}
    </table>
    <button class="add-row" on:click={addRow}>
        +
    </button>
</div>

<style>
    .struct-table {
        width: 100%;
    }

    .add-row {
        width: 100%;
        border: 2px solid black;
        border-radius: 10px;
        text-align: center;
    }

    .remove-row {
        border: 2px solid black;
        border-radius: 10px;
        text-align: center;
    }
</style>