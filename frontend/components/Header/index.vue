<template>
    <div class="header flex items-center h-15 justify-between p-6 bg-white bg-opacity-20 absolute top-0 left-0 w-screen">
        <HeaderLogo />
        <ul class="flex gap-4">
            <li>
                <NuxtLink class="hover:text-slate-800" to="/">Home</NuxtLink>
            </li>
            <li>
                <NuxtLink class="hover:text-slate-800" to="/projects">Projects</NuxtLink>
            </li>
            <li>
                <NuxtLink class="hover:text-slate-800" to="/volunteer">Volunteer</NuxtLink>
            </li>
            <li>
                <NuxtLink class="hover:text-slate-800" to="/user">User Info</NuxtLink>
            </li>
            <li>
                <NuxtLink class="hover:text-slate-800" to="/offers">Offers</NuxtLink>
            </li>
            <li>
                <NuxtLink class="hover:text-slate-800" to="/donate">Donate</NuxtLink>
            </li>
        </ul>
    </div>
</template>

<script setup>

</script>

<style lang="sass" scoped>
.router-link-exact-active
    font-weight: bold
</style>