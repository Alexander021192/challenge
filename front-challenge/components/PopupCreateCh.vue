<template>
	<div class="modal-overlay">
		<div class="modal">
			<!-- Modal content -->
			<div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
				<button 
					type="button" 
					class="absolute top-3 right-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-800 dark:hover:text-white" 
					@click="$emit('close-modal')"
				>
					<svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
					<span class="sr-only">Close modal</span>
				</button>
				<div class="py-6 px-6 lg:px-8">
					<h3 class="mb-4 text-xl font-medium text-gray-900 dark:text-white">Создание нового компонента</h3>
					<form class="space-y-6" action="#">
						<div>
							<label for="title" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Заголовок</label>
							<input type="title" v-model="postTitle" id="title" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white" placeholder="Заголовок" required="">
						</div>
						<img
								v-if="imageUrl"
								class="w-auto rounded-lg mx-auto"
								:src="imageUrl"
								alt=""
								width="1310"
								height="873"
						/>
						<div>
							<label for="loc" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Location</label>
							<input type="loc" v-model="loc" id="loc" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white" placeholder="Location" required="">
						</div>
						<div class="flex-col w-full py-4 mx-auto mt-3 bg-white border-2 border-gray-200 sm:px-4 sm:py-4 md:px-4 sm:rounded-lg sm:shadow-sm">
							<label for="password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">Описание</label>
							<div class="py-2 px-4 bg-white rounded-t-lg dark:bg-gray-800">
								<label for="comment" class="sr-only">Your comment</label>
								<textarea id="comment" rows="5" v-model="postMsg" class="px-0 w-full text-sm text-gray-900 bg-white border-0 dark:bg-gray-800 focus:ring-0 dark:text-white dark:placeholder-gray-400" placeholder="Write a comment..." required=""></textarea>
							</div>
							<div class="flex justify-between items-center py-2 px-3 border-t dark:border-gray-600">
								<div class="flex pl-0 space-x-1 sm:pl-2">
									<button type="button" class="p-2 text-gray-500 rounded-lg cursor-pointer hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-600">
										<svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM7 9a1 1 0 100-2 1 1 0 000 2zm7-1a1 1 0 11-2 0 1 1 0 012 0zm-.464 5.535a1 1 0 10-1.415-1.414 3 3 0 01-4.242 0 1 1 0 00-1.415 1.414 5 5 0 007.072 0z" clip-rule="evenodd"></path></svg>
										<span class="sr-only">Add emoji</span>
									</button>
									<button 
										type="button" 
										class="inline-flex justify-center p-2 text-gray-500 rounded cursor-pointer hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-600"
										@click="$refs.uploadImage.click()"
									>
										<input id="uploadImage" type="file" @change="fileSelected" ref="uploadImage" hidden>
										<svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M4 3a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V5a2 2 0 00-2-2H4zm12 12H4l4-8 3 6 2-4 3 6z" clip-rule="evenodd"></path></svg>
										<span class="sr-only">Upload image</span>                        
									</button>
								</div>
							</div>
						</div>
						<button @click="uploadPost" type="button" class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Создать компонент</button>
					</form>
				</div>
			</div>
		</div>
	</div> 
</template>

<script>
export default {
	name: 'ThePopup',
	data() {
		return {
			postTitle:"",
			loc:"",
			postMsg: "",
			imageUrl: null,
		}
	},
		methods: {
		async fileSelected(e) {
			const file = e.target.files[0];

			/* Make sure file exists */
			if (!file) return;

			/* create a reader */
			const readData = (f) =>  new Promise((resolve) => {
				const reader = new FileReader();
				reader.onloadend = () => resolve(reader.result);
				reader.readAsDataURL(f);
			});

			/* Read data */
			const data = await readData(file);

			/* upload the converted data */
			this.$cloudinary.upload(data, {
				folder: 'upload-examples',
				uploadPreset: 'wf9lxaa4',
			}).then((res) => {
				this.imageUrl = res.url
				// console.log(res.url)
			})
		},
		async uploadPost() {
			const bodyData = {
				"sessionId" : this.$store.state.sessionId,
				"title" : this.postTitle,
				"location" : this.loc,
				"postText" : this.postMsg,
				"postImg" : this.imageUrl,
			}
			// console.log(this.$store.state.sessionId, this.postTitle, this.loc, this.postMsg, this.imageUrl)

			await fetch('http://localhost:8080/create_post', {
			// headers: {
			//     "Content-Type": "application/json",
			// },
			method: 'POST',
			body: JSON.stringify(bodyData)
			}).then((response) => response.json())
			.then((data) => {
				if (data.message) {
					alert(data.message)
					this.$router.replace({ path: '/login' });
				} else {
					alert("Post upload success")
					this.$emit('close-modal')
				}
			})
			.catch((error) => {
				alert(error);
			});
		},
	}
}
</script>

<style scoped>

.modal-overlay {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: center;
  background-color: #000000da;
}

.modal {
  text-align: center;
  background-color: white;
  height: 700px;
  width: 70%;
  margin-top: 10%;
  padding: 30px 0;
  border-radius: 20px;
  overflow-y: auto;
}

</style>