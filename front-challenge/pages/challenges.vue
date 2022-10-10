<template>
    <div class="container mx-auto mt-8">
        <NuxtLink to="/">Home</NuxtLink>
        <button
            v-if="!show"
            class="text-sm bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-2 rounded-full mb-4" 
            @click="showData"
        >
            Show Data
        </button>
        <button
            v-if="show"
            class="text-sm bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-2 rounded-full mb-4" 
            @click="showData"
        >
            Hide Data
        </button>
        <LazyItemsChallenge v-if="show"/>
    </div>
</template>
<script>
export default {
    data() {
        return {
            show: false
        }
    },
    methods: {
        showData() {
            this.show = !this.show
        }
    }
}
</script>