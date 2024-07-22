<template>
  <div class="flex flex-row flex-wrap">
    <div v-for="donator in donators" :key="donator.ID" class="flex items-center bg-white shadow-md rounded-lg p-4 m-1">
      <div class="relative">
        <div class="icon absolute -top-4 -left-4 bg-white text-white font-ibm rounded-full w-8 h-8 flex items-center justify-center">
          <font-awesome-icon :icon="['fas', 'medal']" class="w-full h-full object-contain" style="color: #40BA21;" />
        </div>
        <div class="w-12 h-12 bg-gray-300 rounded-full flex items-center justify-center overflow-hidden">
          <font-awesome-icon :icon="['fas', 'user']" class="w-full h-full object-cover"/>
        </div>
      </div>
      <div class="ml-4">
        <div class=" font-ibm text-[#40BA21] font-semibold">{{ donator.User.Name + ' ' + donator.User.LastName}}</div>
        <div class="text-gray-500 font-ibm" >{{ donator.Amount }} roubles</div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps(['donators'])
</script>

<style lang="sass" scoped>
.icon
    border: 1px solid #A9A9A9

</style>
