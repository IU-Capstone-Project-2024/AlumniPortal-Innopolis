<template>
    <div :class="[route.name == 'sign-in' || route.name =='sign-up' ? 'bg-opacity-75 bg-white !dark:bg-opacity-75 dark:bg-black' : 'bg-opacity-20 bg-white dark:!bg-black dark:!bg-opacity-20', route.name !== 'sign-in' && route.name !== 'sign-up' && route.name !== 'index' ? '!bg-[#40BA21] !bg-opacity-10' : '']"
        class="header flex items-center h-15 justify-between p-6 absolute top-0 left-0 w-screen z-10">
        <HeaderLogo />
        <ul v-show="!config.getMobile" class="flex gap-4 font-ibm">
            <li>
                <NuxtLink class="link hover:text-[#40BA21] font-ibm dark:text-white" to="/">Home</NuxtLink>
            </li>
            <li v-if="route.name != 'sign-in' && route.name != 'sign-up'">
                <NuxtLink class="link hover:text-[#40BA21] font-ibm dark:text-white" to="/projects">Projects</NuxtLink>
            </li>
            <li v-if="route.name != 'sign-in' && route.name != 'sign-up'">
                <NuxtLink class="link hover:text-[#40BA21] font-ibm dark:text-white" to="/volunteer">Volunteer</NuxtLink>
            </li>
            <li v-if="route.name != 'sign-in' && route.name != 'sign-up'">
                <NuxtLink class="link hover:text-[#40BA21] font-ibm dark:text-white" to="/user">Profile</NuxtLink>
            </li>
            <li v-if="route.name != 'sign-in' && route.name != 'sign-up'">
                <NuxtLink class="link hover:text-[#40BA21] font-ibm dark:text-white" to="/offers">Offers</NuxtLink>
            </li>
            <li>
                <NuxtLink class="link hover:text-[#40BA21] font-ibm dark:text-white" to="/sign-in">Sign In</NuxtLink>
            </li>
        </ul>
        <div v-show="config.getMobile" class="icon" :class="{ 'icon-active': mobileNav }">
          <font-awesome-icon @click="toggleMobileNav" v-show="config.getMobile" icon="fa-solid fa-bars" />
        </div>
        <transition name="mobile-nav">
          <ul v-show="mobileNav" class="dropdown-nav flex flex-col gap-4 font-ibm">
            <li>
              <NuxtLink class="link hover:text-[#40BA21] font-ibm dark:text-white" to="/">Home</NuxtLink>
            </li>
            <li v-if="route.name != 'sign-in' && route.name != 'sign-up'">
              <NuxtLink class="link hover:text-[#40BA21] font-ibm dark:text-white" to="/projects">Projects</NuxtLink>
            </li>
            <li v-if="route.name != 'sign-in' && route.name != 'sign-up'">
              <NuxtLink class="link hover:text-[#40BA21] font-ibm dark:text-white" to="/volunteer">Volunteer</NuxtLink>
            </li>
            <li v-if="route.name != 'sign-in' && route.name != 'sign-up'">
              <NuxtLink class="link hover:text-[#40BA21] font-ibm dark:text-white" to="/user">Profile</NuxtLink>
            </li>
            <li v-if="route.name != 'sign-in' && route.name != 'sign-up'">
              <NuxtLink class="link hover:text-[#40BA21] font-ibm dark:text-white" to="/offers">Offers</NuxtLink>
            </li>
            <li>
              <NuxtLink class="link hover:text-[#40BA21] font-ibm dark:text-white" to="/sign-in">Sign In</NuxtLink>
            </li>
          </ul>
        </transition>
    </div>
</template>
<script setup>
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import {useConfigStore} from "~/stores/configStore.js";

const route = useRoute()
const mobileNav = ref(false)
const windowWidth = ref(NaN)

const config = useConfigStore()

const toggleMobileNav = () => {
  mobileNav.value = !mobileNav.value
}

onMounted(() => {
  window.addEventListener('resize', checkScreen)
  checkScreen()
})

onBeforeRouteLeave((to, from, next) => {
  mobileNav.value = false
  next()
})

const checkScreen = () => {
  windowWidth.value = window.innerWidth
  if (windowWidth.value <=750) {
    config.enableMobile()
    return
  }
  config.disableMobile()
  mobileNav.value = false
}


</script>

<style lang="sass" scoped>
.router-link-active
    font-weight: bold
    color: #40BA21

.link
  transition: 0.2s ease all

.dropdown-nav
  padding: 2em
  position: fixed
  width: 100%
  max-width: 250px
  height: 100%
  background-color: white
  top: 0
  left: 0
  li
    margin-left: 0
    .link
      color: black

.icon
  transition: 0.5s ease all

.icon-active
  transform: rotate(90deg)
</style>
