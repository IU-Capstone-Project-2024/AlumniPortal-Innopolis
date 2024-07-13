<template>
  <div class="max-w-sm bg-white rounded-lg shadow-md overflow-hidden">
    <img :src="image" alt="Card image" class="w-full h-48 object-cover">
    <div class="p-6">
      <h2 class="text-xl font-semibold mb-2">{{ header }}</h2>
      <p class="text-gray-700 mb-4">{{ description }}</p>
      <UIButton
        @click="learnMore"
        text="Learn more"
        class="mt-7"
      />
    </div>
  </div>
</template>

<script setup>

defineProps({
  image: {
    type: String,
    required: true,
  },
  header: {
    type: String,
    required: true,
  },
  description: {
    type: String,
    required: true,
  },
});

const learnMore = () => {
  // Add your logic here
  alert('Learn more clicked!');
};
</script>
