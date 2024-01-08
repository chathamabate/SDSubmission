<script>
    import { TYPES, TYPE_DEFAULTS, TYPE_IDS } from "./Data";

    // Should be in the format return from StructureSelector.
    export let structure;

    let data = [];

    function addRow() {
        let newRow = structure.map((c) => TYPE_DEFAULTS[c.type_id]);
        data.push(newRow);
        data = data;
    }

    function removeRow(i) {
        data.splice(i, 1);
        data = data;
    }
</script>

<div class="wide">
    <div class="flex header">
        {#each structure as col}
            <div class="flexible">
                <div class="centered-text">
                    {col.name}
                </div>
                <div class="centered-text">
                    {TYPES[col.type_id]}
                </div>
            </div>
        {/each}
        <div class="unflexible padded-element">
            X
        </div>
    </div>

    {#each data as row, i}
        <div class="flex divider">
            {#each row as cell} 
                <input bind:value={cell} />
            {/each}
            <button class="unflexible tall padded-element exit-button" on:click={() => removeRow(i)}>
                X
            </button>        
        </div>

    {/each}

    <div class="flex">
        <button class="flexible padded-element simple-button" on:click={addRow}>
            Add Row
        </button>
        <button class="flexible padded-element simple-button">
            Submit Request
        </button>
    </div>
</div>

<style>

</style>