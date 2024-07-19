<template>
  <div class="w-full bg-gray-300 rounded-full h-8 relative">
    <div class="bg-gray-800 h-8 rounded-full absolute top-0 left-0 transition-all duration-500 ease-in-out" :style="{ width: `${progressPercentage}%` }"></div>
  </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  collected: {
    type: Number,
    required: true
  },
  goal: {
    type: Number,
    required: true
  }
});

const progressPercentage = computed(() => (props.collected / props.goal) * 100);
</script>
