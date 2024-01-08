<script>
    import { TYPES, TYPE_DEFAULTS, TYPE_IDS } from "./Data";

    // Should be in the format return from StructureSelector.
    export let tableName;
    export let structure;
    export let dataHandler;

    let data = [];

    function addRow() {
        let newRow = structure.map((c) => TYPE_DEFAULTS[c.typeID]);
        data.push(newRow);
        data = data;
    }

    function removeRow(i) {
        data.splice(i, 1);
        data = data;
    }

    function submitRequest() {
        if (data.length == 0) {
            alert("Cannot submit empty request.");
            return;
        }

       let dataCopy = new Array(data.length);

        for (let ri = 0; ri < data.length; ri++) {
            let entry = {};
            for (let ci = 0; ci < structure.length; ci++) {
                let expType = structure[ci].typeID;
                
                if (expType == TYPE_IDS["REAL"]) {
                    let val = parseFloat(data[ri][ci]);
                    if (isNaN(val)) {
                        alert("REAL required at (" + ri + ", " + ci + ")");
                        return;
                    }
                    entry[structure[ci].name] = val;
                } else {
                    entry[structure[ci].name] = data[ri][ci];
                }
            }

            dataCopy[ri] = entry;
        }

        dataHandler(tableName, dataCopy);
    }
</script>

<div class="wide">
    <div class="title-bar">
        {tableName}
    </div>
    <div class="flex divider">
        {#each structure as col}
            <div class="flexible">
                <div class="centered-text col-name">
                    {col.name}
                </div>
                <div class="centered-text col-type">
                    {TYPES[col.typeID]}
                </div>
            </div>
        {/each}
        <div class="unflexible padded-element placeholder">
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
    .placeholder {
        color: transparent;
    }

    .col-type {
        font-style: italic;
        color: grey;
        font-size: .7em;
    }
</style>