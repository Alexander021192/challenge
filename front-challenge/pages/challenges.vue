<template>
    <div class="container mx-auto mt-8">
        <Challenge 
            v-for="mountain in mountains" 
            :product="mountain" 
            :key="mountain.title"
            class="mb-4"
        />
    </div>
</template>

<script>
export default {
    async fetch() {
        this.mountains = await fetch(
            'https://api.nuxtjs.dev/mountains'
        ).then((res) => res.json())
    },
    data() {
        return {
            mountains: []
        }
    }
}
</script>
<style lang="postcss">
img {
    height: 120px;
}
</style>