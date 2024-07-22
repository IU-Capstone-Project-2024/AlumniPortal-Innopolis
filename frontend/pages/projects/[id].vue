<template>
  <div class="mt-[6rem] container mx-auto p-4">
    <header class="text-center">
      <h1 class="text-4xl font-montserrat font-bold mb-4 text-[#40BA21]">Funding and Donation System</h1>
    </header>


    <div class="mb-8">
      <h2 class="text-xl font-semibold font-montserrat text-[#40BA21]">Progress</h2>
    </div>

    <div class="mb-10">
      <div class="text-left text-gray-800 mb-4">Goal: {{ currentProject.goal + '$'}}</div>
      <UIProgressBar :collected="currentProject.collected" :goal="currentProject.goal" />
      <div class="flex justify-between mt-4">
        <div class="text-left text-gray-800">Collected {{ currentProject.collected }}/{{ currentProject.goal }}</div>
        <div class="text-right text-gray-800">Days left: {{ daysLeft }}</div>
      </div>
    </div>


    <div class="container flex flex-col sm:flex-row justify-between">
      <UIDonationForm v-if="auth.user.role == 'Alumni'"/>
      
      <div class="wrapper sm:ml-5 mt-5 sm:mt-0">
        <h2 class="text-xl font-semibold font-montserrat text-[#40BA21]">Top Contributors</h2>
        <Ranking :donators="donators"/>
      </div>
    </div>
  </div>

</template>

<script setup>
import {useAuthStore} from "~/stores/authStore.js";

definePageMeta
({
  middleware
      : 'auth'
})

import { useProjectStore } from '/stores/projectStore.js'
import { storeToRefs } from 'pinia'
const route = useRoute()
const store = useProjectStore()
const auth = useAuthStore()
await store.fetchProjects()
const { projects } = storeToRefs(store)
const currentProject = ref(null)

const res = await fetch('https://api.alumni-portal.ru/donation/project/' + route.params.id, {
  credentials: "include",
})

const donators = await res.json()

donators.sort((a, b) => b.Amount - a.Amount)

currentProject.value = projects.value.filter(p => p.id == route.params.id)[0]
const daysLeft = 30;
</script>

<style lang="sass" scoped>
.wrapper
  background-color: white
  padding: 1.75rem
  border-radius: 0.5rem
  box-shadow: 0 1rem 1.5rem rgba(0, 0, 0, 0.1)
  max-height: min-content
</style>