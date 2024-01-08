<script>
    import { TYPES, TYPE_DEFAULTS, TYPE_IDS } from "./Data";

    // Should be in the format return from StructureSelector.
    export let structure;
    export let dataHandler;

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

    function submitRequest() {
        // 2D copy.
        let dataCopy = data.map(row => {
            let entry = {}
            structure.forEach((v, i) => entry[v.name] = row[i]);
            return entry;
        });

        dataHandler(dataCopy);
    }
</script>

<div class="wide">
    <div class="flex header divider">
        {#each structure as col}
            <div class="flexible">
                <div class="centered-text col-name">
                    {col.name}
                </div>
                <div class="centered-text col-type">
                    {TYPES[col.type_id]}
                </div>
            </div>
        {/each}
        <div class="unflexible padded-element corner">
            X
        </div>
    </div>

    {#each data as row, i}
        <div class="flex row-height divider">
            {#each row as cell} 
                <div class="tall flexible">    
                    <input bind:value={cell} />
                </div>
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
        <button class="flexible padded-element simple-button" on:click={submitRequest}>
            Submit Request
        </button>
    </div>
</div>

<style>
    .corner {
        color: transparent;
    }

    .col-type {
        font-style: italic;
        color: grey;
        font-size: .7em;
    }
</style>