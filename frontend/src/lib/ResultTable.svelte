

<script>
    // NOTE: this assumes that data will be a non-empty JSON array
    // populate with objects of all the same format.

    export let data;
    let colNames = Object.keys(data[0])
</script>

<table>
    <tr class="header-row">
        {#each colNames as colName}
            <th class="header-cell">
                {colName}
            </th>
        {/each}
    </tr>
    {#each data as row}
        <tr class="data-row">
            {#each colNames as colName} 
                <th class="data-cell">
                    {row[colName]}
                </th>
            {/each}
        </tr>
    {/each}
</table>

<style>
    table {
    }

    .header-row {

    }

    .header-cell {

    }

    .data-row {

    } 

    .data-cell {

    }
</style>