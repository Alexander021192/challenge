<template>

 <section class="relative flex items-center justify-center antialiased bg-white bg-gray-100 min-w-screen">
    
    <div class="container px-0 m-5 sm:px-5">

        <form>
        <div class="flex-col w-full py-4 mx-auto mt-3 bg-white border-b-2 border-r-2 border-gray-200 sm:px-4 sm:py-4 md:px-4 sm:rounded-lg sm:shadow-sm md:w-2/3">
            <div class="py-2 px-4 bg-white rounded-t-lg dark:bg-gray-800">
                <label for="comment" class="sr-only">Your comment</label>
                <textarea v-model="commentMsg" id="comment" rows="2" class="px-0 w-full text-sm text-gray-900 bg-white border-0 dark:bg-gray-800 focus:ring-0 dark:text-white dark:placeholder-gray-400" placeholder="Write a comment..." required=""></textarea>
                <img v-if="imageUrl" class="mx-auto w-50 h-40 rounded-lg" :src="imageUrl">
            </div>
            <div class="flex justify-between items-center py-2 px-3 border-t dark:border-gray-600">
                <!-- type="submit" -->
                <button
                    type="submit" 
                    class="inline-flex items-center py-2.5 px-4 text-xs font-medium text-center text-white bg-blue-700 rounded-lg focus:ring-4 focus:ring-blue-200 dark:focus:ring-blue-900 hover:bg-blue-800"
                    @click="postComment"
                >
                    Post comment
                </button>
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
        </form>
		<!-- <pre>{{comments}} </pre> -->
		<div v-if="comments.message">
			<ErrorPage :msg="comments.message"/>
        </div>
        <div
            v-for="comment in comments.comments"
			v-else
            :key="comment.comment"
            class="flex-col w-full py-4 mx-auto mt-3 bg-white border-b-2 border-r-2 border-gray-200 sm:px-4 sm:py-4 md:px-4 sm:rounded-lg sm:shadow-sm md:w-2/3"
		>
            <div class="flex flex-row">
                <img 
					class="object-cover w-12 h-12 border-2 border-gray-300 rounded-full" 
					:alt="comment.author" 
					:src="comment.authorImg"
				>
                <div class="flex-col mt-1">
                    <div class="flex items-center flex-1 px-4 font-bold leading-tight">{{comment.author}}
                        <span class="ml-2 text-xs font-normal text-gray-500">{{commentDate(comment.date)}}</span>
                    </div>
                    <div class="flex-1 px-2 ml-3 text-sm font-medium leading-loose text-gray-600">
						{{comment.comment}}
					</div>
                    <img 
						v-if="comment.commentImg" 
						class="mx-auto w-50 h-40 rounded-lg" 
						:src="comment.commentImg"
					>
                </div>
            </div>
        </div>
    </div>
</section>

</template>

<style scoped>
textarea:focus
{
  outline: none;
}
</style>

<script>
import moment from 'moment'
export default {
    name: 'TheComments',
    data() {
            return {
				commentMsg: "",
				imageUrl: null,
				comments: {}
            }
        },
	async fetch() {
		const bodyData = {
			// "sessionId" : this.$store.state.sessionId
			"sessionId" : "sessionTest"
		}
        this.comments = await fetch('http://localhost:8080/get_comments', {
			method: 'POST',
            body: JSON.stringify(bodyData)
		}).then((res) => res.json())
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
        async postComment() {
            const commentDate = moment(new Date()).format('Do MMMM YYYY, h:mm a');
            const bodyData = {
                "sessionId" : this.$store.state.sessionId,
                "date" : commentDate,
                "comment" : this.commentMsg,
                "commentImg" : this.imageUrl,
            }
            // console.log(this.$store.state.sessionId, commentDate, this.commentMsg, this.imageUrl)

            await fetch('http://localhost:8080/create_comment', {
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
                   alert("Comment create success")
                }
            })
            .catch((error) => {
                alert(error);
            });
        },
        commentDate(date) {
            return moment(date, "Do MMMM YYYY, h:mm a").fromNow()
        }
    }
}
</script>
