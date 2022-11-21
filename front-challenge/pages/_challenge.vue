<template>
	<div class="container mx-auto">
		<div class="relative py-16 bg-white overflow-hidden">
			<NuxtLink to="/challenges" class="flex">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M10 18l-7-7m0 0l7-7m-7 7h18"
					/>
				</svg>
				<span>Back</span>
			</NuxtLink>
			<div class="relative px-4 sm:px-6 lg:px-8">
				<div class="text-lg max-w-prose mx-auto mb-6">
					<!-- <pre>{{post.post}}</pre -->
					<p class="text-base text-center leading-6 text-indigo-600 font-semibold tracking-wide uppercase">
						{{post.post.author}}
					</p>
					<h1 class="mt-2 mb-8 text-3xl text-center leading-8 font-extrabold tracking-tight text-gray-900 sm:text-4xl sm:leading-10">
						{{post.post.title}}
					</h1>
				</div>
				<figure>
					<img
						class="w-auto rounded-lg mx-auto"
						:src="post.post.postImg"
						:alt="post.post.title"
						width="1310"
						height="873"
					/>
					<figcaption class="text-center">{{post.title}}</figcaption>
				</figure>
				<div class="prose prose-lg text-gray-500 mx-auto">
					<ul>
						<li><span class="font-bold">Location</span>: {{post.post.location}}</li>
						<!-- <li><span class="font-bold">Country</span>: <span v-for="country in mountain.countries" :key="country">{{country}}</span></li> -->
					</ul>
					<p class="text-xl mb-10">
						{{post.post.postText}}
					</p>
					<Comments />
				</div>
			</div>
		</div>
	</div>
</template>

<script>
import Comments from '../components/Comments.vue'
	export default {
		name: 'ChallengePage',
		components: { Comments },
		// data() {
		// 	return {
		// 		// post: {},
		// 		error: "",
		// 	}
		// },
		async asyncData( {params, store}) {
			
			const bodyData = {
				"sessionId" : store.state.sessionId,
				"postId" : params.challenge,
			}
			// console.log(bodyData)
			const post = await fetch('http://localhost:8080/get_post_by_id', {
				// headers: {
				//     "Content-Type": "application/json",
				// },
				method: 'POST',
				body: JSON.stringify(bodyData)
				}).then((response) => response.json())
				// .catch((error) => {
				// 	this.error = error;
				// });
			return { post }
		},
	}
</script>

<style scoped>
svg {
	width: 20px;
	height: 20px;
	display: inline-block;
}

img {
	height: 600px;
	object-fit: cover;
}

li {
	margin-bottom: 10px;
	margin-top: 10px;
}

</style>