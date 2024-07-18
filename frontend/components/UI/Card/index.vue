<template>
  <div class="max-w-sm bg-white rounded-lg shadow-md overflow-hidden">
    <img :src='image' alt="Card image" class="w-full h-48 object-cover">
    <div class="p-6">
      <h2 class="text-xl font-semibold mb-2">{{ header }}</h2>
      <p class="text-gray-700 mb-4">{{ description }}</p>
      <UIButton
        @click="$emit('clicked', project.id)"
        text="Learn more"
        class="mt-7"
        btn_type="button"
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
  project: {
    type: {
      id: {
        type: Number,
      }
    },
    required: true,
  }
});
</script>
