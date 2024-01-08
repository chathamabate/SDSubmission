

<script>
    import { afterUpdate } from "svelte";
    import { TYPE_IDS, TYPES } from "./Data";

    // NOTE: this assumes that data will be a non-empty JSON array
    // populate with objects of all the same format.
    export let data;

    let initialData = null;

    let structure = [];
    let matrix = [];

    function generateStructure() {
        let newStructure = [];

        let entry = data[0];
        for (const colName in entry) {
            let val = entry[colName];

            // Determine type of given value.
            let typeID = TYPE_IDS["REAL"];
            if (typeof val == "string") {
                typeID = TYPE_IDS["TEXT"];
            } 

            newStructure.push({name: colName, typeID: typeID});
        }

        return newStructure;
    }

    // Generate a new matrix from the structure.
    function generateMatrix() {
        return data.map((entry) => structure.map((c) => entry[c.name]));
    }

    afterUpdate(() => {
        // Everytime our data changes, we need to rebuild
        // the structure array.
        if (initialData != data) {
            initialData = data;

            structure = generateStructure();
            matrix = generateMatrix();
        }
    });
</script>

<div>
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
    </div>

    {#each matrix as row}
        <div class="flex row-height divider">
            {#each row as cell} 
            <div class="tall flexible">    
                <input readonly bind:value={cell} />
            </div>
            {/each}
        </div>
    {/each}
</div>