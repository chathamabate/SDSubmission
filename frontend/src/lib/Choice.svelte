<script>
    // choices should be a non-empty list of strings.
    export let choices;

    // This is a function which takes a string.
    // It is called whenever the choice changes.
    export let choiceHandler;

    let ci = 0; // index of current choice.

    function switchChoice(i) {
        choiceHandler(choices[i]);
        ci = i;
    }
</script>

<div>
    {#each choices as choice, i}
        {#if i == ci} 
            <button class="toggled" on:click={() => switchChoice(i)}>
                {choice}
            </button>
        {:else}
            <button class="untoggled" on:click={() => switchChoice(i)}>
                {choice}
            </button>
        {/if}
    {/each}
</div>

<style>
    @import "../global.css";
</style>