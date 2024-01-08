<script>
    import Choice from "./lib/Choice.svelte";
    import DataTable from "./lib/DataTable.svelte";
    import ResultTable from "./lib/ResultTable.svelte";
    import StructureSelector from "./lib/StructureSelector.svelte";

    let st = [
        {name: "Name", type_id: 1},
        {name: "Age", type_id: 0},
        {name: "Job", type_id: 1},
    ];
    let people = [
        {name: "Bob Allen", age: 24, id: 1},
        {name: "Dave Smith", age: 21, id: 2},
        {name: "Mark Lip", age: 30, id: 3}
    ];

    let c = ["choice1", "choice2", "choice3"];
    function handler(c) {
        console.log(c);
    }

</script>
   
<main>
    <DataTable structure={st} dataHandler={handler} />
</main>

