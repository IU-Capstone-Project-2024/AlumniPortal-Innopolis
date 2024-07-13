<template>
  <div class="container mx-auto p-4">
    <header class="text-center">
      <h1 class="text-4xl font-bold mb-4">Funding and Donation System</h1>
    </header>
    
    <div class="text-center mb-8">
      <h2 class="text-xl font-semibold">Progress</h2>
    </div>
    
    <UIProgressBar :collected="collected" :goal="goal" />
    
    <div class="flex justify-between mt-4">
      <div class="text-left text-blue-600">Collected {{ collected }}/{{ goal }}</div>
      <div class="text-right text-blue-600">Days left: {{ daysLeft }}</div>
    </div>
  </div>
</template>

<script setup>


const collected = 75000;
const goal = 100000;
const daysLeft = 30;
</script>
