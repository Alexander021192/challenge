<template>
<div>
    <h1>About</h1>
    <p v-if="$nuxt.isOffline"> oops offline </p>
    <p v-if="$nuxt.isOnline"> is Online </p>

    <NuxtLink to="/">Home page</NuxtLink>
</div>
</template>