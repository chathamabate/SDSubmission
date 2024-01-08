<script>
    import { TYPES, TYPE_DEFAULTS, TYPE_IDS } from "./Data";
    import { afterUpdate } from "svelte";

    // Should be in the format return from StructureSelector.
    export let tableName;
    export let structure;
    export let dataHandler;

    // NOTE: Unlike in other places of this code base,
    // data will be a matrix (list of lists).
    // NOT a list of dictionaries.
    let matrix = [];

    let initialTableName = tableName;
    let initialStructure = structure;

    // NOTE: this is a little hacky.
    // The idea here is that when the structure of the request changes in any way,
    // we need to reset the data matrix.
    afterUpdate(() => {
        if (tableName != initialTableName || structure != initialStructure) {
            matrix = [];  // Clear matrix.

            initialTableName = tableName;
            initialStructure = structure;
        }
    });

    function addRow() {
        let newRow = structure.map((c) => TYPE_DEFAULTS[c.typeID]);
        matrix.push(newRow);
        matrix = matrix;
    }

    function removeRow(i) {
        matrix.splice(i, 1);
        matrix = matrix;
    }

    function submitRequest() {
        if (matrix.length == 0) {
            alert("Cannot submit empty request.");
            return;
        }

       let dataCopy = new Array(matrix.length);

        for (let ri = 0; ri < matrix.length; ri++) {
            let entry = {};
            for (let ci = 0; ci < structure.length; ci++) {
                let expType = structure[ci].typeID;
                
                if (expType == TYPE_IDS["REAL"]) {
                    let val = parseFloat(matrix[ri][ci]);
                    if (isNaN(val)) {
                        alert("REAL required at (" + ri + ", " + ci + ")");
                        return;
                    }
                    entry[structure[ci].name] = val;
                } else {
                    entry[structure[ci].name] = matrix[ri][ci];
                }
            }

            dataCopy[ri] = entry;
        }

        dataHandler(tableName, dataCopy);
    }
</script>

<div>
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

    {#each matrix as row, i}
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