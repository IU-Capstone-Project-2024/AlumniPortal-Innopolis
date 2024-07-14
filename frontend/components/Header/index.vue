<template>
    <div :class="route.name == 'sign-in' || route.name =='sign-up' ? 'bg-opacity-75 dark:bg-opacity-75' : 'bg-opacity-20 dark:bg-opacity-20', route.name !== 'sign-in' && route.name !== 'sign-up' && route.name !== 'index' ? 'bg-[#40BA21] bg-opacity-10' : ''"
        class="header flex items-center h-15 justify-between p-6 bg-white dark:bg-black absolute top-0 left-0 w-screen z-10">
        <HeaderLogo />
        <ul class="flex gap-4 font-ibm">
            <li>
                <NuxtLink class="hover:text-[#40BA21] font-ibm dark:text-white" to="/">Home</NuxtLink>
            </li>
            <li v-if="route.name != 'sign-in' && route.name != 'sign-up'">
                <NuxtLink class="hover:text-[#40BA21] font-ibm dark:text-white" to="/projects">Projects</NuxtLink>
            </li>
            <li v-if="route.name != 'sign-in' && route.name != 'sign-up'">
                <NuxtLink class="hover:text-[#40BA21] font-ibm dark:text-white" to="/volunteer">Volunteer</NuxtLink>
            </li>
            <li v-if="route.name != 'sign-in' && route.name != 'sign-up'">
                <NuxtLink class="hover:text-[#40BA21] font-ibm dark:text-white" to="/user">Profile</NuxtLink>
            </li>
            <li v-if="route.name != 'sign-in' && route.name != 'sign-up'">
                <NuxtLink class="hover:text-[#40BA21] font-ibm dark:text-white" to="/offers">Offers</NuxtLink>
            </li>
            <li>
                <NuxtLink class="hover:text-[#40BA21] font-ibm dark:text-white" to="/sign-in">Sign In</NuxtLink>
            </li>
        </ul>

    </div>
</template>
<script setup>
const route = useRoute()
</script>

<style lang="sass" scoped>
.router-link-active
    font-weight: bold
    color: #40BA21
</style>
