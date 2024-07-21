<template>
  <div class="mt-[6rem] container mx-auto p-4">
    <header class="text-center">
      <h1 class="text-4xl font-montserrat font-bold mb-4 text-[#40BA21]">Funding and Donation System</h1>
    </header>


    <div class="mb-8">
      <h2 class="text-xl font-semibold font-montserrat text-[#40BA21]">Progress</h2>
    </div>

<!--    <div class="mb-10">-->
<!--      <div class="text-left text-gray-800 mb-4">Goal: {{ currentProject.goal || '' + '$'}}</div>-->
<!--      <UIProgressBar :collected="currentProject.collected || 0" :goal="currentProject.goal || 0" />-->
<!--      <div class="flex justify-between mt-4">-->
<!--        <div class="text-left text-gray-800">Collected {{ currentProject.collected  || ''}}/{{ currentProject.goal  || ''}}</div>-->
<!--        <div class="text-right text-gray-800">Days left: {{ daysLeft }}</div>-->
<!--      </div>-->
<!--    </div>-->


    <div class="container flex flex-col sm:flex-row justify-between">
      <UIDonationForm />
      <div class="container mt-5 sm:mt-0 flex flex-col">
        <span class="wrapper sm:ml-5">ranking</span>
        <span class="wrapper sm:ml-5 mt-5">ranking</span>
      </div>
    </div>
  </div>



</template>

<script setup>
import { useProjectStore } from '/stores/projectStore.js'
import { storeToRefs } from 'pinia'
const route = useRoute()
const store = useProjectStore()
// const { projects } = storeToRefs(store)
// const currentProject = ref(null)
// currentProject.value = projects.value.filter(p => p.id == route.params.id)[0]
const daysLeft = 30;
</script>

<style lang="sass" scoped>
.wrapper
  background-color: white
  padding: 1.75rem
  border-radius: 0.5rem
  box-shadow: 0 1rem 1.5rem rgba(0, 0, 0, 0.1)
</style>