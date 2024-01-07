<script>
    import Tailwind from "./Tailwind.svelte";
    let people = [
        {name: "Bob", age: 24},
        {name: "Dave", age: 21},
        {name: "Mark", age: 30}
    ];
</script>

<Tailwind />
<main class="bg-red-300">
    {#each people as person}
	    <h1>Hello {person.name}</h1>
    {/each}
</main>

<style>
	main {
        color: green;
    }
</style>
