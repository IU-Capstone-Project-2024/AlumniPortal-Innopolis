<script setup lang="ts">
import type { NuxtError } from '#app'
import Button from "~/components/UI/Button.vue";

const props = defineProps({
  error: Object as () => NuxtError
})

const handleError = () => clearError({ redirect: '/' })
</script>

<template>
  <div class="bg-sad h-screen bg-cover flex flex-col items-center justify-center">
    <div class="flex flex-col items-center justify-center w-1/2 mx-auto">
      <h2 class="block font-bold max-w-fit p-2 text-4xl text-white font-montserrat">
        Ошибочка {{ error.statusCode }}
      </h2>
      <template v-if="error.statusCode === 500">
        <h3 class="block font-bold max-w-fit p-2 text-3xl text-white font-montserrat">
          Извините у нас все сломалось :(
        </h3>
        <button
          class="mt-6 flex justify-center py-3 text-2xl px-4 font-montserrat border border-transparent bg-[#40BA21] rounded-md shadow-sm text-sm font-medium text-white hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
          @click="handleError"
        >
          Поплакать дома
        </button>
      </template>
      <template v-else>
        <h3 class="block font-bold max-w-fit p-2 text-3xl text-white font-montserrat">
          Извините, мы этого не сделали :(
        </h3>
        <button
          class="mt-6 flex justify-center py-3 text-2xl px-4 font-montserrat border border-transparent bg-[#40BA21] rounded-md shadow-sm text-sm font-medium text-white hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
          @click="handleError"
        >
          Поплакать дома
        </button>
      </template>
    </div>
  </div>
</template>

<style lang="sass">
.bg-sad
  background-image: url("/404.jpg")
</style>
