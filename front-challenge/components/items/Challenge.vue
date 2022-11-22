<template>
	<div>
		<button
			class="text-sm bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-2 rounded-full mb-4" 
			@click="$fetch"
		>
			Refresh
		</button>
		
		<p v-if="$fetchState.pending">Fetching posts...</p>
		<div v-if="error">
			<ErrorPage :msg="error"/>
		</div>
		<div 
			v-for="post in posts.posts" 
			v-else
			:key="post.id"
			class="item md:flex"
		>
			<div class="md:flex-shrink-0">
				<img 
					class="rounded-lg md:w-56"
					:src="post.postImg" 
					:alt="post.title"
				/>
			</div>
			<div class="container mt-4 md:mt-0 md:ml-6">
				<div class="uppercase tracking-wide text-sm text-indigo-600 font-bold">
					{{post.title}}
				</div>
				<NuxtLink
					:to="post.Id.toString()"
					class="block mt-1 text-lg leading-tight font-semibold text-gray-900 hover:underline"
				> {{post.author}}
				</NuxtLink>
				<p class="mt-2 text-gray-600">
					{{post.postText}}
				</p>
			</div>
		</div>
	</div>
</template>

<script>
import ErrorPage from '../ErrorPage.vue'

export default {
	name: 'ChallengeItem',
	components: { ErrorPage },
	data() {
		return {
			posts: {},
			error: ""
		}
	},

	async fetch() {
		const bodyData = {
			"sessionId" : this.$store.state.sessionId,
		}
		this.posts = await fetch('http://localhost:8080/get_posts', {
			// headers: {
			//     "Content-Type": "application/json",
			// },
			method: 'POST',
			body: JSON.stringify(bodyData)
			}).then((response) => response.json())
			.catch((error) => {
				this.error = error;
			});
	},
	
}
</script>

<style scoped>
img {
	height: 200px;
	object-fit: cover;
}

.item {
	height: 200px;
	margin-bottom: 15px;
}

.container {
	width: 70%;
}

</style>