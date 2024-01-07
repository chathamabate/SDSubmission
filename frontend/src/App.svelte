<script>
    let people = [
        {name: "Bob", age: 24},
        {name: "Dave", age: 21},
        {name: "Mark", age: 30}
    ];
</script>

<main>
    {#each people as person}
	    <h1>Hello {person.name}</h1>
    {/each}
</main>

<style>
	main {
        color: black;
    }
</style>
