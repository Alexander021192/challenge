<template>
  <main>
    <h1>Home page</h1>
    <NuxtLink to="/about">
      About (internal link that belongs to the Nuxt App)
    </NuxtLink>
    <NuxtLink to="/challenges">
      Challenges (internal link that belongs to the Nuxt App)
    </NuxtLink>
    <a href="https://nuxtjs.org">External Link to another page</a>
    <img
        class="w-auto rounded-lg mx-auto"
        src="https://res.cloudinary.com/du72cfh85/image/upload/v1667885828/cld-sample-2.jpg"
        alt="mountain.title"
        width="1310"
        height="873"
    />
  </main>
</template>