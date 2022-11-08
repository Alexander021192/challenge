<template>
  <div>
    <span>Upload to Cloudinary</span>
    <input
      type="file"
      accept=".jpeg,.jpg,.png,image/jpeg,image/png"
      aria-label="upload image button"
      @change="selectFile"
    />
  </div>
</template>

<script>
export default {
  methods: {
    async selectFile(e) {
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
            console.log(res.url)
        })
    }
  }  
}
</script>