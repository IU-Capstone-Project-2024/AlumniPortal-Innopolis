<template>
  <div class="max-w-sm bg-white rounded-lg shadow-md overflow-hidden">
    <img src="~/assets/adminproject.png" alt="Card image" class="w-full h-48 object-contain">
    <div class="p-6">
      <h3 class="text-lg font-ibm font-semibold mb-2">{{ projectName }}</h3>
      <p class="text-gray-700 font-ibm mb-2"><strong>Description:</strong> {{ projectDescription }}</p>
      <p class="text-gray-700 font-ibm mb-2"><strong>Creator's Email:</strong> {{ creatorEmail }}</p>
      <p class="text-gray-700 font-ibm mb-4"><strong>Goal:</strong> ${{ projectGoal }}</p>
      <div class="flex space-x-2 mt-2">
        <UIButton
          btn_type="submit"
          text="Approve"
          class="text-white py-2 px-4 rounded"
          :style="{ backgroundColor: '#40BA21' }"
        />
        <UIButton
          btn_type="submit"
          text="Decline"
          class="text-white py-2 px-4 rounded"
          :style="{ backgroundColor: '#BA2121' }"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  projectName: {
    type: String,
    required: true,
  },
  projectDescription: {
    type: String,
    required: true,
  },
  creatorEmail: {
    type: String,
    required: true,
  },
  projectGoal: {
    type: Number,
    required: true,
  },
});
</script>
