<template>
  <div class="flex flex-col sm:flex-row justify-between items-center">
    <div v-for="donator in donators" :key="donator.id" class="flex items-center bg-white shadow-md rounded-lg p-4 m-2">
      <div class="relative">
        <div class="icon absolute -top-4 -left-4 bg-white text-white font-ibm rounded-full w-8 h-8 flex items-center justify-center">
          <font-awesome-icon :icon="['fas', 'medal']" class="w-full h-full object-contain" style="color: #40BA21;" />
        </div>
        <div class="w-12 h-12 bg-gray-300 rounded-full flex items-center justify-center overflow-hidden">
          <font-awesome-icon :icon="['fas', 'user']" class="w-full h-full object-cover"/>
        </div>
      </div>
      <div class="ml-4">
        <div class=" font-ibm text-[#40BA21] font-semibold">{{ donator.name }}</div>
        <div class="text-gray-500 font-ibm" >{{ donator.amount }} roubles</div>
      </div>
    </div>
  </div>
</template>

<script setup>
const donators = [
  { id: 1, name: "Donator Full Name", amount: 200, avatar: "" },
  { id: 2, name: "Donator Full Name", amount: 200, avatar: "" },
  // Add more donators as needed
];
</script>

<style lang="sass" scoped>
.icon
    border: 1px solid #A9A9A9
</style>
