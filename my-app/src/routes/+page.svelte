<script lang="ts">
    let days: number[] = []
    for (let i = 1; i < 25; i++) {
        days.push(i)
    }

</script>

<div class="flex flex-col h-screen items-center justify-center">
    <h1 class="h1">Advent of Code</h1>
    <h3 class="h3 mt-2 mb-10">
        <a href="https://github.com/bedminer1/SVK_ADVENTOFSVELTE" target="_blank">by @bedminer1</a>
    </h3>
    <div class="grid grid-cols-4 items-center justify-center gap-2 w-1/3 overflow-auto mt-10">
        {#each days as day}
        <a href="/day{day}" class="text-center">Day {day}</a>
        {/each}
    </div>
</div>
