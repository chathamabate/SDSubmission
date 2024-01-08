<script>
    import Inserter from "./lib/Inserter.svelte";
    import QueryViewer from "./lib/QueryViewer.svelte";
</script>
   
<main>
    <div class="title-bar">
        ScratchDB Lite
    </div>

    <div class="container">
        <div class="rounded section">
            <QueryViewer />
        </div>
    
        <div class="break"></div>
    
        <div class="rounded section">
            <Inserter />
        </div>
    </div>
</main>

