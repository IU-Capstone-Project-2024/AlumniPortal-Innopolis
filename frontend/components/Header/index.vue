<template>
    <div class="header flex items-center h-15 justify-between p-6 bg-white bg-opacity-20 absolute top-0 left-0 w-screen z-10">
        <HeaderLogo />
        <ul class="flex gap-4 font-ibm">
            <li>
                <NuxtLink class="hover:text-slate-800 font-ibm" to="/">Home</NuxtLink>
            </li>
            <li>
                <NuxtLink class="hover:text-slate-800 font-ibm" to="/projects">Projects</NuxtLink>
            </li>
            <li>
                <NuxtLink class="hover:text-slate-800 font-ibm" to="/volunteer">Volunteer</NuxtLink>
            </li>
            <li>
                <NuxtLink class="hover:text-slate-800 font-ibm" to="/user">User Info</NuxtLink>
            </li>
            <li>
                <NuxtLink class="hover:text-slate-800 font-ibm" to="/offers">Offers</NuxtLink>
            </li>
            <li>
                <NuxtLink class="hover:text-slate-800 font-ibm" to="/donate">Donate</NuxtLink>
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