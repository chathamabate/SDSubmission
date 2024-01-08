<script>
    import ResultTable from "./ResultTable.svelte";

    const QUERY_URL = "http://localhost:3000/query"
    let queryText = "";
    let dataPromise = null;

    async function query() {
        let resp = await fetch(QUERY_URL + "?q=" + queryText, {
            method: "GET"
        });
        let respJSON = await resp.json();

        if (!resp.ok) {
            alert(respJSON.message);
            return [];
        }

        return respJSON.data;
    }

    function submit() {
        if (queryText.length == 0) {
            alert("Please enter a query.");
            return;
        }

        dataPromise = query();
    }
</script>

<div>
    <div class="title-bar">
        Query Data
    </div>
    <div class="flex">
        <textarea class="flexible padded-element" placeholder="Enter Query..." bind:value={queryText}/>
    </div>
    <button class="simple-button padded-element wide" on:click={submit}>
        Submit 
    </button>

    {#await dataPromise}
        <div></div>
    {:then data}
        {#if data != null && data.length > 0}
            <ResultTable data={data} />
        {/if}
    {/await}
</div>