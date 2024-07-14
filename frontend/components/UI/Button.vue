<template>
    <button 
          :type="btn_type" 
          class="w-full flex justify-center py-3 px-4 font-montserrat border border-transparent rounded-md shadow-sm text-sm font-medium text-white hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
          style="background-color: #40BA21"
        >
          {{ text }}
        </button>
</template>

<script setup>
defineProps({
  text: {
    type: String,
    required: true,
  },
  btn_type: {
    type: String,
    required: true,
  }
});
</script>

<style lang="sass" scoped>

</style>