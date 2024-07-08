<template>
    <div class="header flex items-center h-15 justify-between p-6 bg-white bg-opacity-20 absolute top-0 left-0 w-screen">
        <HeaderLogo />
        <ul class="flex gap-4">
            <li>
                <NuxtLink to="/">Home</NuxtLink>
            </li>
            <li>
                <NuxtLink to="/projects">Projects</NuxtLink>
            </li>
            <li>
                <NuxtLink to="/volunteer">Volunteer</NuxtLink>
            </li>
            <li>
                <NuxtLink to="/user">User Info</NuxtLink>
            </li>
            <li>
                <NuxtLink to="/offers">Offers</NuxtLink>
            </li>
            <li>
                <NuxtLink to="/donate">Donate</NuxtLink>
            </li>
        </ul>
    </div>
</template>

<script setup>

</script>

<style lang="sass" scoped>

</style>