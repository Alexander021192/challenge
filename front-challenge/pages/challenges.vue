<template>
    <div class="container mx-auto mt-8">
        <NuxtLink to="/">Home</NuxtLink>
        <button
            v-if="!show"
            class="text-sm bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-2 rounded-full mb-4" 
            @click="showData"
        >
            Show Posts
        </button>
        <button
            v-if="show"
            class="text-sm bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-2 rounded-full mb-4" 
            @click="showData"
        >
            Hide Posts
        </button>
        <button  
            class="text-sm bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-2 rounded-full mb-4" 
            type="button" 
            @click="showModal = true"
        >
            Create Post
        </button>
        <PopupCreateCh v-show="showModal" @close-modal="showModal = false"/>
        <LazyItemsChallenge v-if="show"/>
        
    </div>
</template>
<script>

export default {
    data() {
        return {
            show: false,
            showModal: false,
        }
    },
    methods: {
        showData() {
            if (this.$store.state.sessionId === "") {
                alert("Session expired")
                this.$router.replace({ path: '/login' });
            } else {
                // console.log(this.$store.state.sessionId)
                this.show = !this.show
            }
        }
    }
}
</script>