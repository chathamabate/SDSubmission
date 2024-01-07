<script>
    // choices should be a non-empty list of strings.
    export let choices;

    // index of current choice.
    export let ci;

    // This is a function which takes a string.
    // It is called whenever the choice changes.
    export let choiceHandler;

    function switchChoice(i) {
        choiceHandler(choices[i]);
        ci = i;
    }
</script>

<div class="tall">
    {#each choices as choice, i}
        {#if i == ci} 
            <button class="tall padded-element toggled" on:click={() => switchChoice(i)}>
                {choice}
            </button>
        {:else}
            <button class="tall padded-element untoggled" on:click={() => switchChoice(i)}>
                {choice}
            </button>
        {/if}
    {/each}
</div>